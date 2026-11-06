package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"image"
	_ "image/png"
	"os"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

//go:embed icon.png
var iconPNG []byte

// The window and sound device are the only platform layer. Gameplay, drawing,
// transitions, AI and the campaign execute in the recovered Go procedures.
type application struct {
	state   *gameState
	surface *ebiten.Image
	sound   *soundStream
	player  *audio.Player
	started bool
}

func repeated(k ebiten.Key) bool {
	n := inpututil.KeyPressDuration(k)
	return n == 1 || (n >= 24 && n%3 == 0)
}
func (a *application) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyF11) || (inpututil.IsKeyJustPressed(ebiten.KeyEnter) && ebiten.IsKeyPressed(ebiten.KeyAlt)) {
		ebiten.SetFullscreen(!ebiten.IsFullscreen())
	}
	keys := []int{}
	for _, r := range ebiten.AppendInputChars(nil) {
		if r >= 32 && r <= 126 {
			keys = append(keys, int(r))
		}
	}
	special := []struct {
		key  ebiten.Key
		code int
	}{{ebiten.KeyEnter, 13}, {ebiten.KeyNumpadEnter, 13}, {ebiten.KeyEscape, 27}, {ebiten.KeyBackspace, 8}, {ebiten.KeyTab, 9}, {ebiten.KeyArrowUp, 256 + 72}, {ebiten.KeyArrowDown, 256 + 80}, {ebiten.KeyArrowLeft, 256 + 75}, {ebiten.KeyArrowRight, 256 + 77}, {ebiten.KeyInsert, 256 + 82}, {ebiten.KeyPageDown, 256 + 81}, {ebiten.KeyF1, 256 + 59}}
	for _, k := range special {
		if repeated(k.key) {
			if k.key == ebiten.KeyEnter && ebiten.IsKeyPressed(ebiten.KeyAlt) {
				continue
			}
			keys = append(keys, k.code)
		}
	}
	// Mac keyboards have no Insert: F2 is an additional rotation shortcut.
	if repeated(ebiten.KeyF2) {
		keys = append(keys, 256+82)
	}
	x, y := ebiten.CursorPosition()
	y = y * 200 / 240
	buttons := 0
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		buttons |= 1
	}
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
		buttons |= 2
	}
	if x < 0 {
		x = 0
	}
	if x > 319 {
		x = 319
	}
	if y < 0 {
		y = 0
	}
	if y > 199 {
		y = 199
	}
	if a.started {
		a.state.resume <- inputState{keys: keys, mouseX: x, mouseY: y, mouseButtons: buttons, quit: ebiten.IsWindowBeingClosed()}
	} else {
		a.started = true
	}
	f, ok := <-a.state.frames
	if !ok {
		if a.state.failure != nil {
			return a.state.failure
		}
		return ebiten.Termination
	}
	a.surface.WritePixels(f.pixels)
	a.sound.append(f.pcm)
	if a.player != nil && !a.player.IsPlaying() {
		a.player.Play()
	}
	return nil
}
func (a *application) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(1, 1.2)
	op.Filter = ebiten.FilterNearest
	screen.DrawImage(a.surface, op)
}
func (a *application) Layout(w, h int) (int, int) { return 320, 240 }
func main() {
	ebiten.SetWindowTitle("The Battle on the Black Sea")
	ebiten.SetWindowSize(960, 720)
	ebiten.SetWindowClosingHandled(true)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetTPS(60)
	ebiten.SetScreenClearedEveryFrame(false)
	if icon, _, err := image.Decode(bytes.NewReader(iconPNG)); err == nil {
		ebiten.SetWindowIcon([]image.Image{icon})
	}
	state := newState()
	stream := newSoundStream()
	defer stream.close()
	app := &application{state: state, surface: ebiten.NewImage(320, 200), sound: stream}
	// Audio initialization errors are reported, while keyboard and video remain usable.
	context := audio.NewContext(48000)
	if p, err := audio.NewPlayer(context, stream); err == nil {
		app.player = p
		p.SetBufferSize(time.Second / 20)
		defer func() { stream.close(); p.Close() }()
	} else {
		fmt.Fprintln(os.Stderr, "Audio:", err)
	}
	go state.run()
	if err := ebiten.RunGame(app); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
