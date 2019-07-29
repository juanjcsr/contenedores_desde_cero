package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
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
	fmt.Printf("Ejecutando en el padre%v\n", os.Args[2:])

	// Preparamos al programa para ejecutar nuestro comando
	args := os.Args

	args[1] = "child"

	cmd := exec.Command("/proc/self/exe", args[1:]...)

	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS,
	}

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Aquí ejecutamos el comando:
	err := cmd.Run()
	if err != nil {
		panic(fmt.Sprintf("error al ejecutar: %v\n", err))
	}
}

func child() {
	// El resto de los argumentos (del 2 en adelante) es el comando que queremos ejecutar
	fmt.Printf("hijo %v\n", os.Args[2:])

	// Preparamos al programa para ejecutar nuestro comando
	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := syscall.Sethostname([]byte("container"))
	if err != nil {
		panic(fmt.Sprintf("sethostname: %v\n", err))
	}

	// Aquí ejecutamos el comando:
	err = cmd.Run()
	if err != nil {
		panic(fmt.Sprintf("error al ejecutar: %v\n", err))
	}
}
