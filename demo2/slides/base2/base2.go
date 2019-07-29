package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	// Esperamos "run" como primer argumento
	switch os.Args[1] {
	case "run":
		run()
	case "child":
		child()
	default:
		panic("mal argumento")
	}

	fmt.Println("== Fin ==")
}

func run() {
	// El resto de los argumentos (del 2 en adelante) es el comando que queremos ejecutar
	fmt.Printf("padre levantando hijo, pid: %d \n", os.Getpid())

	// Preparamos al programa para ejecutar nuestro comando
	args := os.Args

	args[1] = "child"

	cmd := exec.Command("/proc/self/exe", args[1:]...)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Aquí ejecutamos el comando:
	must(cmd.Run())
}

func child() {
	// El resto de los argumentos (del 2 en adelante) es el comando que queremos ejecutar
	fmt.Printf("hijo ejecutando %v, pid: %d\n", os.Args[2:], os.Getpid())

	// Preparamos al programa para ejecutar nuestro comando
	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Aquí ejecutamos el comando:
	must(cmd.Run())
}

func must(err error) {
	if err != nil {
		log.Panicln(err)
	}
}
