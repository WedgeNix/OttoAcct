package src

import (
	"github.com/go-ole/go-ole"
	load "github.com/mrmiguu/Loading"
)

func qbXML() {
	done := load.New("Connecting to QuickBooks via COM")

	conn := ole.Connection{}
	must(conn.Initialize())
	defer conn.Uninitialize()
	done <- false

	must(conn.Create("QBXMLRP2.RequestProcessor.2"))
	defer conn.Release()
	done <- false

	disp, err := conn.Dispatch()
	must(err)
	defer disp.Release()
	done <- false

	_, err = disp.Call("OpenConnection2", "", "OttoAcct", 1)
	must(err)
	done <- false

	res, err := disp.Call("BeginSession", "", 2)
	must(err)
	tkt := res.ToString()
	done <- false

	println("qbXML", "tkt", tkt)

	var XML struct {
	}
	disp.Call("ProcessRequest", tkt, XML)
	done <- false

	_, err = disp.Call("EndSession", tkt)
	must(err)
	done <- false

	_, err = disp.Call("CloseConnection")
	must(err)
	done <- true
}
