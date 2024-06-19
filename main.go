package go_di

import (
	"fmt"
	"go_di/pkg"
)

func main() {
	userService := pkg.InitializeUserService()
	err := userService.GetUser("1")
	if err != nil {
		fmt.Println(err)
	}
}
