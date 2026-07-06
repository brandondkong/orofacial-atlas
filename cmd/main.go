package main

import (
	"fmt"

	"github.com/brandondkong/orofacial-atlas/pkg/config"
	"github.com/fatih/color"
)

func main() {
	color.New(color.Bold, color.FgCyan).Println("\nInitializing Orofacial Atlas...")
	fmt.Println("Looking for configs...")
	if config.DoesConfigExist() {
		// load configs
		println("Loading existing configs")
	} else {
		// onboard new configs
		color.Yellow("No configs found")
		fmt.Print("Booting config wizard...\n\n")
		err := config.PromptNewConfigs()
		if err != nil {
			println(err.Error())
		}
	}

	// load data to RAM based on configs
}
