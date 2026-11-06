# The Battle on the Black Sea

[Русский](README.ru.md)

A native Go reconstruction of Vladimir Bashurov's (BADY) 1992 DOS naval strategy game. Arrange ten ships, fight the computer, advance through the Black Sea campaign and improve your captain's record.

The original 320×200, 16-colour artwork, bitmap lettering, screen transitions, combat animations and PC Speaker effects are retained. The recovered Pascal procedures execute as compiled Go functions, including the original fleet placement, computer strategy, scoring and 16-bit arithmetic. All game resources are embedded in the executable; no DOS installation or original executable is needed. The window uses the original display's 4:3 aspect ratio.

## Controls

| Key | Action |
| --- | --- |
| Arrow keys | Move the ship or aiming cursor |
| Enter | Place the ship, fire, or advance through information screens |
| Insert / Page Down | Rotate a ship during placement |
| F2 | Additional rotation shortcut for keyboards without Insert |
| F1 | Toggle sound |
| Esc | Exit during placement or combat |
| F11 / Alt+Enter | Toggle fullscreen |
| Backspace | Correct the captain's name |

The original mouse input is also supported. Enter the captain's name using Latin letters. Score cards are saved automatically when exiting through the game or closing the window, in the system's per-user configuration directory under `BattleOnTheBlackSea`.

## Build

Install **Go 1.24 or newer**. Run commands from the project root. The first build downloads the dependencies listed in `go.mod` and `go.sum`.

### Windows — amd64 and arm64

PowerShell, on Windows:

```powershell
$env:CGO_ENABLED="0"
$env:GOOS="windows"
$env:GOARCH="amd64"
go build -buildvcs=false -trimpath -ldflags="-s -w -H=windowsgui" -o BattleOnTheBlackSea-windows-amd64.exe .
$env:GOARCH="arm64"
go build -buildvcs=false -trimpath -ldflags="-s -w -H=windowsgui" -o BattleOnTheBlackSea-windows-arm64.exe .
```

No C compiler is needed. The PNG icon is embedded and used by the game window; `icon.ico` is supplied for executable-resource packaging.

### Linux — amd64 and arm64

On Debian/Ubuntu, install the native build dependencies:

```sh
sudo apt-get install build-essential pkg-config libasound2-dev libx11-dev libxrandr-dev libxcursor-dev libxi-dev libxinerama-dev libxxf86vm-dev libgl1-mesa-dev
```

Build on a Linux machine matching the target architecture:

```sh
# On an amd64 Linux machine:
CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -buildvcs=false -trimpath -ldflags="-s -w" -o BattleOnTheBlackSea-linux-amd64 .

# On an arm64 Linux machine:
CGO_ENABLED=1 GOOS=linux GOARCH=arm64 go build -buildvcs=false -trimpath -ldflags="-s -w" -o BattleOnTheBlackSea-linux-arm64 .
```

Linux requires an X11 display (or XWayland), OpenGL and ALSA. Cross-compiling Linux arm64 from amd64 additionally requires an AArch64 C compiler and the arm64 versions of these libraries; set `CC=aarch64-linux-gnu-gcc` and point `PKG_CONFIG_LIBDIR` at the target libraries.

### macOS — Intel and Apple Silicon

Build **on macOS**, with Apple's Command Line Tools installed:

```sh
xcode-select --install
CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 go build -buildvcs=false -trimpath -ldflags="-s -w" -o BattleOnTheBlackSea-macos-amd64 .
CGO_ENABLED=1 GOOS=darwin GOARCH=arm64 go build -buildvcs=false -trimpath -ldflags="-s -w" -o BattleOnTheBlackSea-macos-arm64 .
```

The macOS backend requires Cgo and Apple's SDK; `CGO_ENABLED=0` is not supported. `icon.icns` is supplied for an optional `.app` bundle.

For a native development run, use `go run .` after installing the dependencies for your platform. Each built executable contains the complete game; the icon and screenshot files do not have to accompany it.

## Credits and scope

Original game, artwork and characters: **Vladimir Bashurov (BADY), June 1992**. The original materials retain their authors' rights; the supplied DOS archive contains no redistribution licence. Ebitengine and its dependencies retain their respective licences.

Gameplay and rendering were reconstructed from the supplied executable and resources. The engine retains the original data layout and floating-point rounding; animation uses a fixed timing baseline because the DOS game's speed depended on the computer. Pixel comparisons covered the title and statistics screens; arithmetic and fleet placement were checked against the original procedures. Complete winning and losing games were exercised. Windows and Linux builds were checked for amd64/arm64; macOS builds require validation on a Mac. Identical timing and acoustic output on every original and modern computer are not certified.
