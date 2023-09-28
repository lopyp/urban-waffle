package main

/*
=== Взаимодействие с ОС ===

Необходимо реализовать собственный шелл

встроенные команды: cd/pwd/echo/kill/ps
поддержать fork/exec команды
конвеер на пайпах

Реализовать утилиту netcat (nc) клиент
принимать данные из stdin и отправлять в соединение (tcp/udp)
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"
)

func main() {
	for {
		fmt.Print("> ")
		var input string
		fmt.Scanln(&input)

		args := strings.Fields(input)
		if len(args) == 0 {
			continue
		}

		switch args[0] {
		case "cd":
			changeDirectory(args)
		case "pwd":
			printWorkingDirectory()
		case "echo":
			echoCommand(args)
		case "kill":
			killProcess(args)
		case "ps":
			psCommand()
		default:
			executeExternalCommand(args)
		}
	}
}

func changeDirectory(args []string) {
	var dir string

	if len(args) < 2 {
		dir = getUserHomeDirectory()
	} else {
		dir = args[1]
	}

	absPath, err := filepath.Abs(dir)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	err = os.Chdir(absPath)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func getUserHomeDirectory() string {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error:", err)
	}
	return home
}

func printWorkingDirectory() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(dir)
	}
}

func echoCommand(args []string) {
	fmt.Println(strings.Join(args[1:], " "))
}

func killProcess(args []string) {
	if len(args) < 2 {
		fmt.Println("Usage: kill <process>")
		return
	}
	processName := args[1]
	cmd := exec.Command("pkill", processName)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func psCommand() {
	cmd := exec.Command("ps")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func executeExternalCommand(args []string) {
	child := exec.Command(args[0], args[1:]...)
	child.Stdin = os.Stdin
	child.Stdout = os.Stdout
	child.Stderr = os.Stderr

	err := child.Start()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		err = child.Wait()
		if err != nil {
			if exitError, ok := err.(*exec.ExitError); ok {
				status := exitError.Sys().(syscall.WaitStatus)
				fmt.Printf("Command exited with status %d\n", status.ExitStatus())
			} else {
				fmt.Println("Error:", err)
			}
		}
	}
}
