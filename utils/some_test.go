package utils

import (
	"fmt"
	"gin-vue/conf"
	"os"
	"testing"
)

func TestXxx(t *testing.T) {

	_ = os.Chdir("..")

	conf.InitConfig()
	token, _ := GenerateToken(1, "test")
	fmt.Println(token)

	fmt.Println("=======")
	claims, _ := ParseToken(token)
	fmt.Println(claims)

	fmt.Println("=======")
	fmt.Println(IsTokenValid(token))
	fmt.Println(IsTokenValid(token + "xxx"))
}
