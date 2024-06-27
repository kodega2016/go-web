package main

import (
	"github.com/fatih/color"
)

const portNumber = ":8080"

func main() {
	c := color.New(color.FgGreen).Add(color.Underline)
	c.Printf("application server is running on %s.", portNumber)
}
