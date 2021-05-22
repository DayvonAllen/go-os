package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"os/exec"
)

func executeUnixCommand(command string,  args []string) error {
	// returns a struct that we can set properties on and execute the run method later
	commandObj := exec.Command(command, args...)

	// set the standard out to our OS's standard out so we can output things onto our terminal
	commandObj.Stdout = os.Stdout

	// set the standard error to our OS's standard error so we can output errors onto our terminal
	commandObj.Stderr = os.Stderr

	// set the standard input to our OS's standard input so we can input values into the program from out terminal
	commandObj.Stdin = os.Stdin

	// execute command stored in the commandObject struct
	err := commandObj.Run()

	if err != nil {
		return err
	}

	return err
}

func executeWinCommand(args []string) error {
	// returns a struct that we can set properties on and execute the run method later
	// for win must use command, and /C to run commands
	commandObj := exec.Command("cmd", args...)

	// set the standard out to our OS's standard out so we can output things onto our terminal
	commandObj.Stdout = os.Stdout

	// set the standard error to our OS's standard error so we can output errors onto our terminal
	commandObj.Stderr = os.Stderr

	// set the standard input to our OS's standard input so we can input values into the program from out terminal
	commandObj.Stdin = os.Stdin

	// execute command stored in the commandObject struct
	err := commandObj.Run()

	if err != nil {
		return err
	}

	return err
}

func connectToServer() {
	conn, err := net.Dial("tcp", "golang.org:80")

	if err != nil {
		// handle error
		fmt.Println(err)
	}
	_, err = fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")

	if err != nil {
		fmt.Println(err)
	}

	status, err := bufio.NewReader(conn).ReadString('\n')

	fmt.Println("[+] Connection established.. ", conn.RemoteAddr().String())
	fmt.Println(status)

	err = conn.Close()

	if err != nil {
		fmt.Println(err)
	}
}

func createServer() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		// handle error
		fmt.Println(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
		}
		go func() {
			fmt.Println(conn)
		}()
	}
}

func main() {
	// use flags in the terminal like this (-arg=test)
	// 1st arg = name, 2nd arg = default value, 3rd arg = usage message
	// returns a string pointer
	//commandInterfaceFlag := flag.String("arg", "", "Usage to give an argument to be used with a command")
	//nameInterfaceFlag := flag.String("name", "", "Usage to give an argument name")

	// parse all declared flags
	//flag.Parse()
	// because the 'commandInterfaceFlag contains a pointer to a string, we must dereference it to get the actual
	// string object on the heap
	// example command: go run main.go -arg=mkdir -name=test
	//err := executeWinCommand([]string{"/C", *commandInterfaceFlag, *nameInterfaceFlag})

	//if err != nil {
	//	fmt.Println(err)
	//}

	connectToServer()
}
