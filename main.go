package main

import "github.com/alexandregv/RP42/cmd"

// CLI entrypoint of the program.
// We use main.go and not directly cmd/ to allow a clean `go install github.com/alexandregv/RP42`
func main() {
	cmd.Execute()
}
