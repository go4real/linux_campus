package main

import (
	"fmt"
	"os"
	"os/exec"
)

// docker           run image <CMD> <ARG>
// go run main.go   run       <CMD> <ARG>

// Step2: 새로운 프로세스에서 명령어 실행  예) ls -l

func main() {
	switch os.Args[1] {
	case "run":
		run()
	default:
		os.Exit(1)
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
