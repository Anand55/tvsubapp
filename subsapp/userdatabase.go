package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func UpdateDatabase(userdetails UserDetails, userarr User) {
	file, err := ioutil.ReadFile("user.json")
	if err != nil {
		fmt.Println(err)
	}
	json.Unmarshal(file, &userarr.Users)
	userarr.Users = append(userarr.Users, userdetails)

	dataBytes, err := json.Marshal(userarr.Users)

	if err != nil {
		fmt.Println(err)
	}
	err = ioutil.WriteFile("user.json", dataBytes, 0644)
	if err != nil {
		fmt.Println(err)
	}
}

func UpdatePhoneAndEmail(email string, newemail string, newphone string) {
	arr := ReadData()
	for i := 0; i < len(arr); i++ {
		if arr[i].EmaidId == email {
			arr[i].EmaidId = newemail
			arr[i].MobileNo = newphone
		}
	}
	dataBytes, err := json.Marshal(arr)

	if err != nil {
		fmt.Println(err)
	}
	err = ioutil.WriteFile("user.json", dataBytes, 0644)
	if err != nil {
		fmt.Println(err)
	}
}
