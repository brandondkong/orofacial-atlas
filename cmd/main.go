package main

import (
	"fmt"

	"github.com/brandondkong/orofacial-atlas/pkg/config"
	"github.com/brandondkong/orofacial-atlas/pkg/input"
)

func main() {
	println("Initializing Orofacial Atlas")

	if config.DoesConfigExist() {
		// load configs
		println("Loading existing configs")
	} else {
		// onboard new configs
		println("First time using Orofacial Atlas? Welcome!")

		char, err := input.PromptSingleCharacter("Are you excited? (y/n): ")
		if err != nil {
			panic(err)
		}

		fmt.Printf("%c\n", char)
	}

	// load data to RAM based on configs
}
