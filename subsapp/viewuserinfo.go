package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func ReadData() []UserDetails {
	var arr User
	file, err := ioutil.ReadFile("user.json")
	if err != nil {
		fmt.Println(err)
	}
	json.Unmarshal(file, &arr.Users)

	return arr.Users
}

func ViewBalance(email string) int {
	arr := ReadData()
	var bal int
	for i := 0; i < len(arr); i++ {
		if arr[i].EmaidId == email {
			bal = arr[i].AvailableBalance
			break
		}
	}
	return bal
}

func RechargeAccount(email string, amount int) int {
	arr := ReadData()
	var bal int
	for i := 0; i < len(arr); i++ {
		if arr[i].EmaidId == email {
			arr[i].AvailableBalance = arr[i].AvailableBalance + amount
			bal = arr[i].AvailableBalance
			break
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
	return bal
}

func ViewPacks() {
	fmt.Println("***Available packs for subscription***")
	packs := ViewAvailablePacks()
	for k, v := range packs {
		for k1, v1 := range v {
			fmt.Println(k, ":", k1, ":", v1, " Rs.")
		}
	}
	fmt.Println("***Available Channels for subscription***")
	chans := ViewAvailableChannels()

	for k, v := range chans {
		fmt.Println(k, ":", v, " Rs.")
	}

	fmt.Println("***Available Services for subscription***")
	serv := ViewAvailableServices()

	for k, v := range serv {
		fmt.Println(k, ":", v, " Rs.")
	}

}
