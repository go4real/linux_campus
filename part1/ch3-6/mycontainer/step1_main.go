package main

import (
	"fmt"
	"os"
)

// docker            run image <CMD> <ARG>
// go run main.go   run       <CMD> <ARG>

// Step1: 명령어 종류에 따른 함수 실행. "run" 명령어 전달 시 run 함수 실행.

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
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
