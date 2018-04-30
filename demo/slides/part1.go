package main

import (
	"fmt"
	"os"
	"os/exec"
)

// go run main.go run <cmd> <args>
func main() {
	switch os.Args[1] {
	case "run":
		fmt.Printf("ejecutando %v con pid %d\n", os.Args[2:], os.Getpid())
		run()
	case "child":
		fmt.Printf("ejecutando hijo con pid %d\n", os.Getpid())
		run()
	default:
		panic("OMG!")
	}
}

func parent() {
	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	must(cmd.Run())
}

func child() {
	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	must(cmd.Run())
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
