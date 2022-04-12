package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

// docker            run image <CMD> <ARG>
// go run main.go   run       <CMD> <ARG>

// Step5: 컨테이너 환경에서 ps명령 실행 시 제한된 프로세스 정보만 조회. 루트 파일 시스템 변경.
//        실습으로 ps, cat /os-release 명령 실행.

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
	fmt.Printf("Running: %v as %d\n", os.Args[2:], os.Getpid())
	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout

	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
	}

	must(cmd.Run())
}

func child() {
	fmt.Printf("Running child: %v as %d\n", os.Args[2:], os.Getpid())

	// hostname 설정
	syscall.Sethostname([]byte("container"))

	/* 루트 파일시스템 다운로드
	# https://github.com/tianon/docker-brew-ubuntu-core/raw/88ba31584652db8b96a29849ea2809d99ce3cc31/focal/ubuntu-focal-oci-amd64-root.tar.gz
	# mkdir /tmp/ubuntu
	# tar zxf ubuntu-focal-oci-amd64-root.tar.gz -C /tmp/ubuntu
	*/
	// 루트 파일시스템 변경
	// syscall.Chroot("/tmp/ubuntu")
	// syscall.Chdir("/")
	// syscall.Mount("proc", "proc", "proc", 0, "")
	// defer syscall.Unmount("proc", 0)

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
