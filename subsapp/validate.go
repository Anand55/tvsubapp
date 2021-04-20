package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func ValidateUser(email string) bool {
	var arr User
	flag := false
	file, err := ioutil.ReadFile("user.json")
	if err != nil {
		fmt.Println(err)
	}
	json.Unmarshal(file, &arr.Users)
	for i := 0; i < len(arr.Users); i++ {
		if arr.Users[i].EmaidId == email {
			flag = true
			break
		}
	}
	return flag

}
