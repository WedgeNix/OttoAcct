package src

import (
	"os/exec"

	"github.com/mrmiguu/jsutil"

	load "github.com/mrmiguu/Loading"
	"github.com/mrmiguu/sock"
)

func compile() {
	done := load.New("Compiling OttoAcct.go")

	must(jsutil.CompileWithGzip(sock.Root + "/OttoAcct.go"))
	done <- false

	out, err := exec.Command("cmd", "/c", "start", "http://"+sock.Addr+"/").Output()
	must(err)
	if len(out) > 0 {
		panic(string(out))
	}
	done <- true
}
