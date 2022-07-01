package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"os/exec"
	"runtime"
	"time"
)

var cmd *exec.Cmd

func run(program string) {
	cmd = exec.Command("./" + program)
	err := outputStd(cmd)
	if err != nil {
		log.Fatalln("GUN ---- get stdout(stderr) failed with", err)
	}
	err = cmd.Run()
	if err != nil {
		log.Println("GUN ---- run cmd.Run() failed with", err)
	}
}

func build(program string) error {
	cmd = exec.Command("go", "build", "-o", program, "main.go")
	err := outputStd(cmd)
	if err != nil {
		log.Fatalln("GUN ---- get stdout(stderr) failed with", err)
	}
	log.Println("Begin building...")
	err = cmd.Run()
	if err != nil {
		log.Println("GUN ----build cmd.Run() failed with", err)
		return err
	}
	log.Println("Built Successfully!")
	return nil
}

func outputStd(cmd *exec.Cmd) error {
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return err
	}
	go print(stdout)
	go print(stderr)
	return nil
}

func print(read io.ReadCloser) {
	scanner := bufio.NewScanner(read)
	for scanner.Scan() {
		msg := scanner.Bytes()
		log.Println(string(msg))
	}
}

func kill() {
	if cmd != nil && cmd.Process != nil {
		// Windows does not support Interrupt
		if runtime.GOOS == "windows" {
			cmd.Process.Signal(os.Kill)
		} else {
			cmd.Process.Signal(os.Interrupt)
		}
		ch := make(chan struct{}, 1)
		go func() {
			cmd.Wait()
			ch <- struct{}{}
		}()

		select {
		case <-ch:
			return
		case <-time.After(10 * time.Second):
			log.Println("Timeout. Force kill cmd process")
			err := cmd.Process.Kill()
			if err != nil {
				log.Printf("Error while killing cmd process: %s", err)
			}
			return
		}
	}
}
