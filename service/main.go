package main

import (
	"fmt"

	"github.com/hunterpunchh/mono/imp/me"
	"github.com/hunterpunchh/mono/imp/you"
)

func main() {
	fmt.Printf("Importing from %s %s", me.Me, me.Version)
	fmt.Printf("Importing from %s %s", you.You, you.Version)
}
