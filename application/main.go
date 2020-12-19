package main

import (
	"dfs/client"
	"fmt"
)

// application program to interact with the client
func main() {
	res := client.Read("asdf")
	fmt.Println(string(res))
}
