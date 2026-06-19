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
	CGO_ENABLED=0 go build -o $(BINARY) .
	@echo "Binary built: $(BINARY)"

appdir: build
	@echo "Creating AppDir..."
	mkdir -p $(APPDIR)/usr/bin
	cp $(BINARY) $(APPDIR)/usr/bin/
	cp bfxr2.desktop $(APPDIR)/
	cp AppRun $(APPDIR)/
	chmod +x $(APPDIR)/AppRun
	cp bfxr2.svg $(APPDIR)/bfxr2.png 2>/dev/null; \
	cp ../icons/bfxr2.png $(APPDIR)/ 2>/dev/null; \
	cp ../../bfxr2.png $(APPDIR)/ 2>/dev/null; \
	cp bfxr2.png $(APPDIR)/ 2>/dev/null; \
	[ -f "$(APPDIR)/bfxr2.png" ] || touch $(APPDIR)/bfxr2.png
	@echo "AppDir created."

APPIMAGETOOL := $(shell command -v appimagetool 2>/dev/null || echo /tmp/appimagetool)

appimage: appdir
	@echo "Building AppImage..."
	if [ -x "$(APPIMAGETOOL)" ]; then \
		"$(APPIMAGETOOL)" $(APPDIR) $(APPIMAGE) && echo "AppImage built: $(APPIMAGE)"; \
	else \
		echo "appimagetool not found. Install it from https://github.com/AppImage/appimagetool/releases"; \
	fi

clean:
	rm -rf $(BINARY) web $(APPDIR) $(APPIMAGE)

webview-setup:
	@echo "Setting up vendored webview_go (webkit2gtk-4.1 patch)..."
	rm -rf wv
	go mod download github.com/webview/webview_go@v0.0.0-20240831120633-6173450d4dd6
	cp -r "$(shell go env GOMODCACHE)/github.com/webview/webview_go@v0.0.0-20240831120633-6173450d4dd6" wv
	chmod -R u+w wv
	sed -i 's/webkit2gtk-4\.0/webkit2gtk-4.1/g' wv/webview.go
	$(RM) wv/go.mod wv/go.sum
	@echo "Webview vendored. Run: go build -tags webview -o $(BINARY) ."

webview-build: webfiles webview-setup
	CGO_ENABLED=1 go build -tags webview -o $(BINARY)-webview .

run: build
	./$(BINARY)

.PHONY: all clean webfiles build appdir appimage run webview-setup webview-build
