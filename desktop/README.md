# Bfxr2 Desktop App

A standalone desktop wrapper for [Bfxr2](https://www.bfxr.net/), the sound effects generator for games. Runs as a local HTTP server serving the embedded web app and opens it in your default browser.

## Quick Start

```sh
make build
./bfxr2
```

This starts a local server on a random available port and opens your browser.

Use `--browser` to skip the webview (opens browser only):
```sh
./bfxr2 --browser
```

## Building

### Prerequisites

- [Go](https://go.dev/) 1.21+

### Commands

| Command | Description |
|---------|-------------|
| `make webfiles` | Copy web files into the embedded `web/` directory |
| `make build` | Build the Go binary |
| `make appimage` | Build AppImage (requires `appimagetool`) |
| `make run` | Build and run the desktop app |
| `make clean` | Remove build artifacts |

### AppImage

To build the AppImage:

```sh
make appimage
```

This produces `bfxr2-x86_64.AppImage` — a portable, self-contained executable.
Download and run it anywhere without installing:

```sh
chmod +x bfxr2-x86_64.AppImage
./bfxr2-x86_64.AppImage
```

### Webview Build (Optional)

For a native desktop window instead of a browser tab, build with the `webview` tag:

```sh
# Install system dependencies (Ubuntu/Debian)
sudo apt install libgtk-3-dev libwebkit2gtk-4.0-dev

# Build with webview
GOFLAGS=-tags=webview go build -o bfxr2 .
```

Note: The webview variant requires GTK3 and WebKit2GTK development libraries.

## Structure

```
desktop/
├── main.go          # Entry point (default: browser mode)
├── webview.go       # Webview variant (build tag: webview)
├── go.mod           # Go module
├── wv/              # Vendored webview_go (for webview builds)
├── Makefile         # Build automation
├── AppDir/          # AppImage skeleton
└── web/             # Embedded web files (generated)
```
