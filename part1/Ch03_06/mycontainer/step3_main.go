package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

// docker           run image <CMD> <ARG>
// go run main.go   run       <CMD> <ARG>

// Step3: 새로운 UTS 설정 추가. Clone flag 추가 NEW UTS namespace. hostname 수동 변경 실습.
// 실습
// $ go run . run /bin/sh
// $ hostname box

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

	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS,
	}

	must(cmd.Run())
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
