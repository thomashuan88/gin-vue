package main

import (
	"gin-vue/cmd"
)

// @title Swagger Example API
// @version 1.1
// @description This is a sample server celler server.
// @termsOfService http://swagger.io/terms/
func main() {
	defer cmd.Clean()
	cmd.Start()

}
