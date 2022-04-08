package main

import (
	"fmt"
	"os"
	"os/exec"
)

// docker            run image <CMD> <ARG>
// go run main.go   run       <CMD> <ARG>

// Step2: 새로운 프로세스에서 명령어 실행

func main() {
	switch os.Args[1] {
	case "run":
		run()
	default:
		panic("bad command")
	}
}

func run() {
	fmt.Printf("Running: %v\n", os.Args[2:])
	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout

	cmd.Run()
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
