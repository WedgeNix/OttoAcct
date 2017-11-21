package src

import (
	"github.com/mrmiguu/jsutil"

	"github.com/mrmiguu/sock"
)

func init() {
	run()
}

func run() {
	Files := sock.Wbytes()

	println("starting!")
	dnd := document.getElementByID("drag_n_drop")

	dnd.Set("ondrop", func(ev event) {
		println("drop")
		ev.preventDefault()
		dt := ev.dataTransfer()
		for _, f := range dt.files {
			go func(f file) {
				Files <- <-jsutil.FetchBlob(f.Object)
			}(f)
		}
	})

	dnd.Set("ondragover", func(ev event) { ev.preventDefault() })
}
