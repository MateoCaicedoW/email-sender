package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: tools database <command>")
		fmt.Println("Available commands:")
		fmt.Println(" - migrate")
		fmt.Println(" - create")
		fmt.Println(" - drop")
		fmt.Println(" - generate <name>")
		return
	}

	switch os.Args[1] {
	case "migrate":

	}

}
