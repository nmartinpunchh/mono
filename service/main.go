package main

import (
	"fmt"

	"github.com/nmartinpunchh/mono/imp/me"
	"github.com/nmartinpunchh/mono/imp/you"
)

func main() {
	fmt.Printf("Importing from %s %s\n", me.Me, me.Version)
	fmt.Printf("Importing from %s %s\n", you.You, you.Version)
}
