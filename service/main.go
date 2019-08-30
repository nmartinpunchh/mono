package main

import (
	"fmt"

	"github.com/hunterpunchh/mono/imp/me"
	"github.com/hunterpunchh/mono/imp/you"
)

func main() {
	fmt.Printf("Importing from %s %s\n", me.Me, me.Version)
	fmt.Printf("Importing from %s %s\n", you.You, you.Version)
}
