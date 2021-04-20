package main

import "fmt"

type User struct {
	Users []UserDetails
}

type UserDetails struct {
	Name             string `json:"Name"`
	MobileNo         string `json:"MobileNo"`
	EmaidId          string `json:"EmailID"`
	Subscription     string `json:"Subscription"`
	AvailableBalance int    `json:"AvailableBalance"`
	SpecialServices  string `json:"SpecialServices"`
}

func UpdateNewUserDetails(name, mobileNo, emaidId string) UserDetails {
	arr := ReadData()
	flag := false
	for i := 0; i < len(arr); i++ {
		if arr[i].EmaidId == emaidId {
			fmt.Println("Email Address already exists")
			flag = true
			break
		}
	}
	u := new(UserDetails)
	if flag == false {

		u.Name = name
		u.MobileNo = mobileNo
		u.EmaidId = emaidId
		u.AvailableBalance = 100
	}
	return *u
}

func (user *User) UpdateUserDatabase(UserDetails UserDetails) User {
	user.Users = append(user.Users, UserDetails)
	return *user
}
