.PHONY: build-desktop dev-desktop build-appimage clean-desktop

build-desktop: install-deps
	cd src-tauri && cargo build --release

dev-desktop:
	cd src-tauri && cargo tauri dev

build-appimage: build-desktop
	cd src-tauri && cargo tauri build --bundles appimage

install-deps:
	@echo "Ensuring system dependencies..."
	@if command -v apt-get >/dev/null 2>&1; then \
		sudo apt-get install -y libgtk-3-dev libwebkit2gtk-4.1-dev \
		libsoup-3.0-dev libjavascriptcoregtk-4.1-dev libappindicator3-dev; \
	elif command -v dnf >/dev/null 2>&1; then \
		sudo dnf install -y gtk3-devel webkit2gtk4.1-devel \
		libsoup3-devel javascriptcoregtk4.1-devel libappindicator-gtk3-devel; \
	elif command -v pacman >/dev/null 2>&1; then \
		sudo pacman -S --noconfirm gtk3 webkit2gtk-4.1 \
		libsoup3 javascriptcoregtk-4.1 libappindicator-gtk3; \
	elif command -v brew >/dev/null 2>&1; then \
		echo "On macOS, Xcode and Xcode Command Line Tools are required."; \
		echo "No additional brew packages needed for Tauri."; \
	fi

clean-desktop:
	cd src-tauri && cargo clean
	rm -rf src-tauri/frontend/
