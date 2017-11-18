package main

import (
	"strconv"

	"github.com/gopherjs/gopherjs/js"
	"github.com/mrmiguu/sock"
)

var (
	document = struct {
		getElementByID func(string) *js.Object
	}{func(s string) *js.Object { return js.Global.Get("document").Call("getElementById", s) }}

	console = struct {
		log func(string)
	}{func(s string) { js.Global.Get("console").Call("log", s) }}
)

type event struct{ *js.Object }

func (e event) preventDefault()          { e.Call("preventDefault") }
func (e event) dataTransfer() *js.Object { return e.Get("dataTransfer") }

func main() {
	Files := sock.Wbytes()

	console.log("starting!")
	dnd := document.getElementByID("drag_n_drop")

	dnd.Set("ondrop", func(ev event) {
		console.log("Drop")
		ev.preventDefault()
		dt := ev.dataTransfer()
		for i := 0; i < dt.Get("files").Length(); i++ {
			console.log("... file[" + strconv.Itoa(i) + "].name = " + dt.Get("files").Index(i).Get("name").String())
			Files <- []byte(dt.Get("files").Index(i).Get("name").String())
		}
	})

	dnd.Set("ondragover", func(ev event) {
		console.log("dragOver")
		ev.preventDefault()
	})

	dnd.Set("ondragend", func(ev event) {
		console.log("dragEnd")
		ev.preventDefault()
	})
}
