package main

import (
	"fmt"
)

func main() {
	fmt.Println("Let's build backend APIs for bnaking services!")

	server := NewAPIServer(":3000")
	server.Run()
}
