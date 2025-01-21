package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// Argument adalah input yang diberikan oleh user ketika menjalankan program
	// Flag adalah input yang diberikan oleh user dengan format -flag=value

	// Contoh argument
	argsRaw := os.Args
	fmt.Printf("-> %#v\n", argsRaw)

	args := argsRaw[1:]
	fmt.Printf("-> %#v\n", args)

	// Penggunaan flag
	var name = flag.String("name", "anonymous", "type your name")
	var age = flag.Int("age", 20, "type your age")
	flag.Parse()

	fmt.Printf("Hello %s, you are %d years old\n", *name, *age)

}
