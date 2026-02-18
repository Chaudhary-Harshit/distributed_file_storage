package main

import (
	"log"

	"github.com/harshitchaudhary/distributed_file_system/p2p"
)

func main() {
	tr := p2p.NewTCPTransport(":3000")

	// shorthand declaration and assignment of variable, and error handling in one line, before ; - is assignment, after ; is the condition to check for error, if error is not nil, then execute the code in the if block
	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}

	select {} //blocks forever (keeps main alive forever, so that goroutines can run in background)

}
