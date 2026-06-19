//go:build webview

package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/webview/webview_go"
)

func init() {
	runWebview = func(port int) {
		w := webview.New(true)
		defer w.Destroy()

		w.SetTitle("Bfxr2")
		w.SetSize(1200, 800, webview.HintNone)
		w.Navigate(fmt.Sprintf("http://localhost:%d", port))

		sig := make(chan os.Signal, 1)
		signal.Notify(sig, os.Interrupt)
		go func() {
			<-sig
			w.Terminate()
		}()

		w.Run()
	}
}
