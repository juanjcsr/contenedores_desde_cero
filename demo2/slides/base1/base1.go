package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	// We expect "run" as the first argument
	switch os.Args[1] {
	case "run":
		run()
	default:
		panic("mal argumento")
	}

	fmt.Println("== Fin ==")
}

func run() {
	// Arguments 2 onwards are the arbitrary command we're going to run
	fmt.Printf("Ejecutando %v\n", os.Args[2:])

	// Set up a struct that describes the command we want to run
	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// This is where we run the command
	must(cmd.Run())

}

func must(err error) {
	if err != nil {
		log.Panicln(err)
	}
}
