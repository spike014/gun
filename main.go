package main

import (
	"log"
	"path/filepath"
)

const sign = `
██████╗ ██╗   ██╗███╗   ██╗
██╔════╝ ██║   ██║████╗  ██║
██║  ███╗██║   ██║██╔██╗ ██║
██║   ██║██║   ██║██║╚██╗██║
╚██████╔╝╚██████╔╝██║ ╚████║
 ╚═════╝  ╚═════╝ ╚═╝  ╚═══╝  v0.1.1-beta
`

var (
	base     string
	basePath string
)

func main() {
	log.Print(sign, "\n")
	basePath, _ = getExcutePath()

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
