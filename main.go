package main

import (
	"os"

	"github.com/WLaoDuo/olive/command"
)

func main() {
	command.Execute(os.Args[1:])
}
