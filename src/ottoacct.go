package src

import (
	"github.com/mrmiguu/Loading"
	"github.com/mrmiguu/sock"
)

func init() {
	done := load.New("OttoAcct running")

	// some parts can be properly automated
	// while others require drap-n-drop functionality

	sock.Root = "src/www"

	for File := range sock.Rbytes() {
		println(string(File))
	}

	done <- true
}
