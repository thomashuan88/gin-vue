package main

import (
	"gin-vue/cmd"
)

func main() {
	defer cmd.Clean()
	cmd.Start()

}
