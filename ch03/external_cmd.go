package main

import (
	"io"
	"log"
	"os"
	"os/exec"
	"sync"
	"syscall"
)

func main() {
	tr(os.Stdin, os.Stdout, os.Stderr)
}

func tr(src io.Reader, dst io.Writer, errDist io.Writer) error {
	cmd := exec.Command("tr", "a-z", "A-Z")
	// execute command is tr a-z A-Z
	stdin, _ := cmd.StdinPipe()
	stdout, _ := cmd.StdoutPipe()
	stderr, _ := cmd.StderrPipe()
	err := cmd.Start() // start command
	if err != nil {
		return err
	}
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		// copy from src to command's stdin
		_, err := io.Copy(stdin, src)
		if e, ok := err.(*os.PathError); ok && e.Err == syscall.EPIPE {
			// ignore EPIPE
		} else if err != nil {
			log.Println("failed to write to STDIN", err)
		}
		stdin.Close()
		wg.Done()
	}()
	go func() {
		// copy from stdout to command's stdout
		io.Copy(dst, stdout)
		stdout.Close()
		wg.Done()
	}()
	go func() {
		// copy from command's stderr to errDist
		io.Copy(errDist, stderr)
		stderr.Close()
		wg.Done()
	}()
	wg.Wait()
	// wait for all goroutine
	return cmd.Wait()
	// wait for command execution
}
