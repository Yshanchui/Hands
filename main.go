package main

import "Hands/cmd"

func main() {
	defer cmd.Clean()
	cmd.Start()
}
