package execing_processes

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func RunSample() {
	fmt.Println("Exec processes package output:")
	binary, err := exec.LookPath("ls")
	if err != nil {
		panic("ls missing")
	}
	args := []string{"ls", "-l", "-a", "-h"}
	env := os.Environ()
	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
	fmt.Println("")
}
