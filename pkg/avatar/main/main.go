package main

import (
	"fmt"
	"pkg/pkg/avatar"
)

func main() {
	var username string
	println("Enter your username: ")
	fmt.Scanln(&username)
	avatar.AvatarGenerator().GenerateAndSaveAvatar(avatar.Information{UserName: username})
}
