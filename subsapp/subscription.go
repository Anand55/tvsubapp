package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

var AvailablePacks map[string]map[string]int
var AvailableChannels map[string]int
var AvailableServices map[string]int

func ViewAvailablePacks() map[string]map[string]int {
	Totalpacks := make(map[string]map[string]int)
	pack := make(map[string]int)
	pack["Zee, Sony, Star Plus"] = 50
	Totalpacks["Silver"] = pack
	pack1 := make(map[string]int)
	pack1["Zee, Sony, Star Plus, Discovery, NatGeo"] = 100
	Totalpacks["Gold"] = pack1
	AvailablePacks = Totalpacks
	return AvailablePacks
}

func ViewAvailableChannels() map[string]int {
	chans := make(map[string]int)
	chans["Zee"] = 10
	chans["Sony"] = 15
	chans["Star Plus"] = 20
	chans["Discovery"] = 10
	chans["NatGeo"] = 20
	AvailableChannels = chans
	return AvailableChannels
}

func ViewAvailableServices() map[string]int {
	serv := make(map[string]int)
	serv["LearnEnglish"] = 200
	serv["LearnCooking"] = 100
	AvailableServices = serv
	return AvailableServices
}

func AddSubscription(email string, pack string, months int) (int, bool, int) {
	arr := ReadData()
	var mnprice int
	var flag bool
	var accbal int
	flag = true
	for i := 0; i < len(arr); i++ {
		avpacks := ViewAvailablePacks()
		if arr[i].EmaidId == email {
			if pack == "S" {
				arr[i].Subscription = arr[i].Subscription + "Silver "
				for k, v := range avpacks {
					if k == "Silver" {
						for _, v1 := range v {
							mnprice = v1
							if (arr[i].AvailableBalance) >= (months * v1) {
								if months >= 3 {
									arr[i].AvailableBalance = arr[i].AvailableBalance - (months * v1) + int((months*v1)*10/100)
									accbal = arr[i].AvailableBalance
								} else {
									arr[i].AvailableBalance = arr[i].AvailableBalance - (months * v1)
									accbal = arr[i].AvailableBalance
								}
							} else {
								flag = false
							}
						}
					}
				}

			} else if pack == "G" {
				arr[i].Subscription = arr[i].Subscription + "Gold "
				for k, v := range avpacks {
					if k == "Gold" {
						for _, v1 := range v {
							mnprice = v1
							if (arr[i].AvailableBalance) >= (months * v1) {
								if months >= 3 {
									arr[i].AvailableBalance = arr[i].AvailableBalance - (months * v1) + int((months*v1)*10/100)
									accbal = arr[i].AvailableBalance
								} else {
									arr[i].AvailableBalance = arr[i].AvailableBalance - (months * v1)
									accbal = arr[i].AvailableBalance
								}
							} else {
								flag = false
							}
						}
					}
				}
			}
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
	return mnprice, flag, accbal
}

func AddChannels(email string, channels string) int {
	arr := ReadData()
	var totalAmount int
	var bal int
	avchannels := ViewAvailableChannels()
	chanarr := strings.Split(channels, ",")
	for i := 0; i < len(arr); i++ {
		if arr[i].EmaidId == email {
			for j := 0; j < len(chanarr); j++ {
				totalAmount = totalAmount + avchannels[chanarr[j]]
			}
			if arr[i].AvailableBalance >= totalAmount {
				for k := 0; k < len(chanarr); k++ {
					arr[i].Subscription = arr[i].Subscription + chanarr[k] + " "
				}
				arr[i].AvailableBalance = arr[i].AvailableBalance - totalAmount
				bal = arr[i].AvailableBalance
			}
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

func AddServices(email string, service string) int {
	arr := ReadData()
	var bal int
	avserv := ViewAvailableServices()
	for i := 0; i < len(arr); i++ {
		if arr[i].EmaidId == email {
			if arr[i].AvailableBalance >= avserv[service] {
				arr[i].SpecialServices = service
				arr[i].AvailableBalance = arr[i].AvailableBalance - avserv[service]
				bal = arr[i].AvailableBalance
			}
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

func ViewSubscriptionDetails(email string) {
	arr := ReadData()
	// var bal int
	// var chpackstr string
	for i := 0; i < len(arr); i++ {
		if arr[i].EmaidId == email {
			if arr[i].Subscription == "" {
				fmt.Println("Currently subscribed packs and channels : No Channels and Packs Subscribed yet")
			} else {
				chpack := strings.Split(arr[i].Subscription, " ")
				chpackstr := strings.Join(chpack, "+")
				fmt.Println("Currently subscribed packs and channels : ", chpackstr[0:len(chpackstr)-1])
			}
			if arr[i].SpecialServices == "" {
				fmt.Println("Currently subscribed services : No services Subscribed yet")
			} else {
				fmt.Println("Currently subscribed services : ", arr[i].SpecialServices)
			}
		}
	}
}
