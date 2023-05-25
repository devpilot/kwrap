package main

import (
	"bytes"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/chzyer/readline"
)

func main() {
	const colorYellow string = "\033[33m"
	const colorReset string = "\033[0m"
	user := "john"
	host := "supercomputer"
	rl, err := readline.New(colorYellow + user + "@" + host + colorReset + "~$ kubectl ")
	if err != nil {
		panic(err)
	}
	defer rl.Close()
	for {
		line, err := rl.Readline()
		if err == readline.ErrInterrupt {
			log.Println("ctrl c pressed")
		}
		if err != nil { // io.EOF
			break
		}
		args := strings.Split(line, " ")
		cmd := exec.Command("kubectl", args...)
		// err = cmd.Start()

		var stdBuffer bytes.Buffer
		mw := io.MultiWriter(os.Stdout, &stdBuffer)
		cmd.Stdout = mw
		cmd.Stderr = mw
		// Execute the command
		if err := cmd.Run(); err != nil {
			log.Panic(err)
		}

		// log.Println(stdBuffer.String())

		// if err != nil {
		// 	log.Println("failed executing command", err)
		// }
		// output := string(out[:])
		// log.Println(output)
	}
}
