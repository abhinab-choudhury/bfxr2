package main

import (
	"embed"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"runtime"
)

//go:embed web
var webFiles embed.FS

var runWebview func(port int)

func main() {
	port := getAvailablePort()
	go startServer(port)

	if runWebview != nil && (len(os.Args) < 2 || os.Args[1] != "--browser") {
		runWebview(port)
	} else {
		url := fmt.Sprintf("http://localhost:%d", port)
		fmt.Printf("Bfxr2 running at %s\n", url)
		openBrowser(url)
		fmt.Println("Press Ctrl+C to stop")
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, os.Interrupt)
		<-sig
	}
}

func startServer(port int) {
	addr := fmt.Sprintf(":%d", port)
	http.Handle("/", http.FileServer(http.FS(webFiles)))
	log.Printf("Serving bfxr2 on http://localhost:%d\n", port)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}

func openBrowser(url string) {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "rundll32"
		args = []string{"url.dll,FileProtocolHandler", url}
	case "darwin":
		cmd = "open"
		args = []string{url}
	default:
		cmd = "xdg-open"
		args = []string{url}
	}

	if err := exec.Command(cmd, args...).Start(); err != nil {
		log.Printf("Failed to open browser: %v", err)
	}
}

func getAvailablePort() int {
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		return 8080
	}
	port := listener.Addr().(*net.TCPAddr).Port
	listener.Close()
	return port
}
