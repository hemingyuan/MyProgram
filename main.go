package main

import (
	"MyProgram/getLocalAddr"
	"fmt"
)

func main() {
	r := getLocalAddr.GetLocalAddr()
	fmt.Println(r)
}
