package src

import "github.com/gopherjs/gopherjs/js"

var (
	document = struct {
		getElementByID func(string) *js.Object
	}{func(s string) *js.Object { return js.Global.Get("document").Call("getElementById", s) }}
)

type event struct {
	*js.Object
}

func (e event) preventDefault() { e.Call("preventDefault") }
func (e event) dataTransfer() dataTransferItem {
	dt := e.Get("dataTransfer")
	var files []file
	for i := 0; i < dt.Get("files").Length(); i++ {
		files = append(files, file{dt.Get("files").Index(i)})
	}
	return dataTransferItem{dt, files}
}

type dataTransferItem struct {
	*js.Object
	files []file
}

type file struct {
	*js.Object
}
