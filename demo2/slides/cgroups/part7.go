package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"syscall"
)

func main() {
	// Esperamos "run" como primer argumento
	switch os.Args[1] {
	case "run":
		fmt.Printf("host ejecutando hijo, pid: %d \n",
			os.Getpid())
		parent()
	case "child":
		// El resto de los argumentos (del 2 en adelante) es el comando que queremos ejecutar
		fmt.Printf("hijo ejecutando %v, pid: %d\n",
			os.Args[2:],
			os.Getpid())
		child()
	default:
		panic("mal argumento")
	}
}

func parent() {
	// Preparamos al padre para ejecutar el hijo
	args := os.Args
	args[1] = "child"
	cmd := exec.Command("/proc/self/exe", args[1:]...)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS |
			syscall.CLONE_NEWPID |
			syscall.CLONE_NEWNS,
		// Unshareflags: syscall.CLONE_NEWNS,
	}

	// Aquí ejecutamos el comando:
	must(cmd.Run())
	fmt.Println("== Fin host ==")
}

func child() {

	cg()
	// Preparamos al hijo para ejecutar nuestro comando
	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// syscall para el hostname
	must(syscall.Sethostname([]byte("container")))
	must(syscall.Chroot("/home/vagrant/containers/fs/rootfs-ubuntu"))
	must(syscall.Chdir("/"))
	must(syscall.Mount("proc", "proc", "proc", 0, ""))
	// Aquí ejecutamos el comando:
	must(cmd.Run())

	must(syscall.Unmount("proc", 0))

	fmt.Println("== Fin child ==")
}

func must(err error) {
	if err != nil {
		log.Panicln(err)
	}
}

func cg() {
	cgroups := "/sys/fs/cgroup/"
	pids := filepath.Join(cgroups, "pids")

	must(os.MkdirAll(filepath.Join(pids, "demo3"), 755))
	must(ioutil.WriteFile(filepath.Join(pids, "demo3/pids.max"), []byte("20"), 0700))

	//asigna el pid y elimina el cgroup cuando el contenedor termina
	must(ioutil.WriteFile(filepath.Join(pids, "demo3/notify_on_release"), []byte("1"), 0700))
	must(ioutil.WriteFile(filepath.Join(pids, "demo3/cgroup.procs"), []byte(strconv.Itoa(os.Getpid())), 0700))
}
