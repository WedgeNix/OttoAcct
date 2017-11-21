package src

import (
	"github.com/mrmiguu/Loading"
	"github.com/mrmiguu/sock"
)

func init() {
	qbXML()
	compile()
	run()
}

func run() {
	done := load.New("Running OttoAcct")

	for File := range sock.Rbytes() {
		println(string(File))
	}

	done <- true
}
