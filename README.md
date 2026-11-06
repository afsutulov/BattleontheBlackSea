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

Install **Go 1.25 or newer**. **Ebitengine 2.10.5** and its dependencies are pinned in `go.mod` and `go.sum`. Run commands from the project root. The first build downloads these dependencies.

All six desktop targets support **`CGO_ENABLED=0`**. Cross-compilation needs only Go; no C compiler, target development headers or Apple SDK is required.

### Windows — amd64 and arm64

PowerShell:

```powershell
$env:CGO_ENABLED="0"
$env:GOOS="windows"
$env:GOARCH="amd64"
go build -buildvcs=false -trimpath -ldflags="-s -w -H=windowsgui" -o BattleOnTheBlackSea-windows-amd64.exe .
$env:GOARCH="arm64"
go build -buildvcs=false -trimpath -ldflags="-s -w -H=windowsgui" -o BattleOnTheBlackSea-windows-arm64.exe .
```

The PNG icon is embedded as the window icon; `icon.ico` is supplied for executable-resource packaging.

### Linux — amd64 and arm64

From Linux or macOS:

```sh
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -buildvcs=false -trimpath -ldflags="-s -w" -o BattleOnTheBlackSea-linux-amd64 .
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -buildvcs=false -trimpath -ldflags="-s -w" -o BattleOnTheBlackSea-linux-arm64 .
```

From **Windows CMD**, for Linux ARM64:

```cmd
set "CGO_ENABLED=0"
set "GOOS=linux"
set "GOARCH=arm64"
go build -buildvcs=false -trimpath -ldflags="-s -w" -o BattleOnTheBlackSea-linux-arm64 .
set "CGO_ENABLED="
set "GOOS="
set "GOARCH="
```

The resulting file runs on Linux ARM64. After transferring it, run `chmod +x BattleOnTheBlackSea-linux-arm64` and then `./BattleOnTheBlackSea-linux-arm64`.

At runtime, Linux needs an X11 display (or XWayland) and an OpenGL driver. Audio uses PulseAudio/PipeWire or an ALSA fallback. On Debian/Ubuntu, the basic graphics libraries can be installed with:

```sh
sudo apt install libx11-6 libgl1 libglx-mesa0 libxcursor1 libxi6 libxinerama1 libxrandr2 libxrender1 libxext6
```

For the ALSA fallback, install `libasound2` on Debian 12/Ubuntu 22.04, or `libasound2t64` on Debian 13/Ubuntu 24.04 and later. These are runtime libraries, not build dependencies.

### macOS — Intel and Apple Silicon

From a Unix shell on Linux or macOS:

```sh
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -buildvcs=false -trimpath -ldflags="-s -w" -o BattleOnTheBlackSea-macos-amd64 .
CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -buildvcs=false -trimpath -ldflags="-s -w" -o BattleOnTheBlackSea-macos-arm64 .
```

For Windows CMD, use the same `set` syntax as the Linux example, with `GOOS=darwin` and the desired `GOARCH`. `icon.icns` is supplied for an optional `.app` bundle.

For a native development run, use `go run .` from a graphical desktop session. Each built executable contains the complete game; the icon and screenshot files do not have to accompany it.

Verification: `go test ./...` and `go vet ./...`. The audio regression tests retain coverage of the Windows startup deadlock fix.

## Credits and scope

Original game, artwork and characters: **Vladimir Bashurov (BADY), June 1992**. The original materials retain their authors' rights; the supplied DOS archive contains no redistribution licence. Ebitengine and its dependencies retain their respective licences.

Gameplay and rendering were reconstructed from the supplied executable and resources. The engine retains the original data layout and floating-point rounding; animation uses a fixed timing baseline because the DOS game's speed depended on the computer. Pixel comparisons covered the title and statistics screens; arithmetic and fleet placement were checked against the original procedures. Complete winning and losing games were exercised. Windows and Linux builds were checked for amd64/arm64; macOS builds require validation on a Mac. Identical timing and acoustic output on every original and modern computer are not certified.
