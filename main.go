package main

import (
	"fmt"
	"path/filepath"
)

const sign = `
██████╗ ██╗   ██╗███╗   ██╗
██╔════╝ ██║   ██║████╗  ██║
██║  ███╗██║   ██║██╔██╗ ██║
██║   ██║██║   ██║██║╚██╗██║
╚██████╔╝╚██████╔╝██║ ╚████║
 ╚═════╝  ╚═════╝ ╚═╝  ╚═══╝  v0.1-beta
`

var (
	base     string
	basePath string
)

func main() {
	fmt.Print(sign, "\n")
	basePath, _ = getExcutePath()

	if checkIfIsNotGoProject(basePath) {
		panic("It is NOT a go project. Can NOT find main.go")
	}

	updateFilespath()
	go watch()
	base = filepath.Base(basePath)
	err := build(base)
	if err != nil {
		return
	}
	run(base)
	run := make(chan bool)
	<-run
}
