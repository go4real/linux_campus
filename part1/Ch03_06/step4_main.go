package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

// docker           run image <CMD> <ARG>
// go run main.go   run       <CMD> <ARG>

// Step4: 컨테이너 환경 시작시 호스트명을 container로 변경.
// $ go run . run /bin/sh
// $ hostname

func main() {
	switch os.Args[1] {
	case "run":
		run()
	case "child":
		child()
	default:
		os.Exit(1)
	}
}

func run() {
	fmt.Printf("Running: %v\n", os.Args[2:])
	// cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout

	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS,
	}

	must(cmd.Run())
}

func child() {
	fmt.Printf("Running child: %v\n", os.Args[2:])

	// hostname 설정
	syscall.Sethostname([]byte("container"))

	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout

	must(cmd.Run())
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
