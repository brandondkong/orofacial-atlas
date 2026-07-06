package main

import "github.com/brandondkong/orofacial-atlas/pkg/config"

func main() {
	print("Initializing Orofacial Atlas\n")

	if config.DoesConfigExist() {
		// load configs
		print("Loading existing configs\n")
	} else {
		// onboard new configs
		print("First time using Orofacial Atlas? Welcome!\n")
	}

	// load data to RAM based on configs
}
