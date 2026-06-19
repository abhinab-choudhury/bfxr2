APPNAME := bfxr2
BINARY := $(APPNAME)
APPDIR := AppDir
APPIMAGE := $(APPNAME)-$(shell uname -m).AppImage

.PHONY: all clean webfiles build appdir appimage

all: webfiles build appimage

webfiles:
	@echo "Copying web files..."
	mkdir -p web/js/audio web/js/synths web/js/third_party web/css/third_party
	cp ../index.html web/
	cp ../js/index.js web/js/
	cp ../js/Tab.js web/js/
	cp ../js/SaveLoad.js web/js/
	cp ../js/globals.js web/js/
	cp ../js/audio/*.js web/js/audio/
	cp ../js/synths/*.js web/js/synths/
	cp ../js/third_party/*.js web/js/third_party/
	cp ../css/index.css web/css/
	cp ../css/slider.css web/css/
	cp ../css/third_party/*.css web/css/third_party/
	cp -r ../sounds web/ 2>/dev/null || true
	@echo "Web files copied."

build: webfiles
	@echo "Building Go binary..."
	CGO_ENABLED=1 go build -o $(BINARY) .
	@echo "Binary built: $(BINARY)"

appdir: build
	@echo "Creating AppDir..."
	mkdir -p $(APPDIR)/usr/bin
	cp $(BINARY) $(APPDIR)/usr/bin/
	cp ../icons/bfxr2.png $(APPDIR)/ 2>/dev/null || \
		convert -size 256x256 xc:transparent -font Helvetica -pointsize 128 \
			-fill '#4a90d9' -gravity center -annotate +0+0 'B' $(APPDIR)/bfxr2.png 2>/dev/null || \
		touch $(APPDIR)/bfxr2.png
	cat > $(APPDIR)/$(APPNAME).desktop <<-EOF
		[Desktop Entry]
		Type=Application
		Name=Bfxr2
		Comment=Sound effects generator for games
		Exec=usr/bin/$(BINARY)
		Icon=$(APPNAME)
		Terminal=false
		Categories=AudioVideo;Audio;Game;
	EOF
	cat > $(APPDIR)/AppRun <<-EOF
		#!/bin/bash
		exec "$$APPDIR/usr/bin/$(BINARY)"
	EOF
	chmod +x $(APPDIR)/AppRun
	@echo "AppDir created."

appimage: appdir
	@echo "Building AppImage..."
	which appimagetool && appimagetool $(APPDIR) $(APPIMAGE) || \
		echo "appimagetool not found. Install it from https://github.com/AppImage/appimagetool/releases"
	@echo "AppImage built: $(APPIMAGE)"

clean:
	rm -rf $(BINARY) web $(APPDIR) $(APPIMAGE)

run: build
	./$(BINARY)

.PHONY: all clean webfiles build appdir appimage run
