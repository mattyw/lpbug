package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
)

// run go install to but this into your path.
// Open this file (or any other in vim)
// Find bug links like this lp:1
// get to the start of the word then
// :!lpbug C-rC-a
func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Take a single argument in the form of the bug number")
		fmt.Println("Uses regular expressions to find the number so lp: and bug: are accepted")
		return
	}

	bug := args[1]
	re := regexp.MustCompile("[0-9]+")
	number := re.FindString(bug)
	url := fmt.Sprintf("http://launchpad.net/bugs/%s", number)

	cmd := exec.Command("xdg-open", url)
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}

}
