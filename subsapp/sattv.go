package main

import (
	"fmt"
)

func WelcomeConsole(emailid string) {
	fmt.Println("*****WELCOME TO SATTV*****")
	fmt.Println("1. View current balance in the account")
	fmt.Println("2. Recharge Account")
	fmt.Println("3. View available packs, channels and services")
	fmt.Println("4. Subscribe to base packs")
	fmt.Println("5. Add channels to an existing subscription")
	fmt.Println("6. Subscribe to special services")
	fmt.Println("7. View current subscription details")
	fmt.Println("8. Update email and phone number for notifications")
	fmt.Println("9. Exit")
	fmt.Println("PLEASE SELECT THE OPTION : ")
	var option int
	fmt.Scanln(&option)
	switch option {
	case 1:
		bal := ViewBalance(emailid)
		fmt.Println("Current Balance is ", bal, " Rs.")
		fmt.Println()
		WelcomeConsole(emailid)
	case 2:
		var amount int
		fmt.Println("Enter the amount to recharge : ")
		fmt.Scanln(&amount)
		bal := RechargeAccount(emailid, amount)
		fmt.Println("Recharge completed successfully. Current balance is ", bal)
		fmt.Println()
		WelcomeConsole(emailid)
	case 3:
		ViewPacks()
		fmt.Println()
		WelcomeConsole(emailid)
	case 4:
		fmt.Println("Subscribe to Channels packs")
		fmt.Println("Enter the Pack you wish to subscribe: (Silver: ‘S’, Gold: ‘G’)")
		var packsub string
		fmt.Scanln(&packsub)
		fmt.Println("Enter the months : ")
		var months int
		fmt.Scanln(&months)
		mnprice, flag, accbal := AddSubscription(emailid, packsub, months)
		if flag == true {
			fmt.Println("Monthly Price : ", mnprice, " Rs.")
			fmt.Println("No. of Months : ", months)
			fmt.Println("Subscription Amount : ", months*mnprice, " Rs.")
			if months >= 3 {
				fmt.Println("Discount Applied : ", int((months*mnprice)*10/100))
				fmt.Println("Final Price after Discount : ", ((months * mnprice) - int((months*mnprice)*10/100)))
			}
			fmt.Println("Account Balance : ", accbal, " Rs.")
		} else {
			fmt.Println("Please recharge your account")
		}
		fmt.Println()
		WelcomeConsole(emailid)
	case 5:
		fmt.Println("Add channels to an existing subscription")
		fmt.Println("Enter Channels names to add (seprated by commas) : ")
		var channels string
		fmt.Scanln(&channels)
		bal := AddChannels(emailid, channels)
		fmt.Println("Account Balance : ", bal, " Rs.")
		fmt.Println()
		WelcomeConsole(emailid)
	case 6:
		fmt.Println("Subscribe to special services")
		fmt.Println("Enter the Service Name : ")
		var service string
		fmt.Scanln(&service)
		bal := AddServices(emailid, service)
		fmt.Println("Service Subscribed Successfully")
		fmt.Println("Account Balance : ", bal, " Rs.")
		fmt.Println("Email Notification Sent Successfully")
		fmt.Println("SMS Notification Sent Successfully")
		fmt.Println()
		WelcomeConsole(emailid)
	case 7:
		fmt.Println("View current subscription details")
		ViewSubscriptionDetails(emailid)
		fmt.Println()
		WelcomeConsole(emailid)
	case 8:
		fmt.Println("Update email and phone number for notifications")
		var newemail string
		var newphone string
		fmt.Println("Enter the email : ")
		fmt.Scanln(&newemail)
		fmt.Println("Enter Phone : ")
		fmt.Scanln(&newphone)
		UpdatePhoneAndEmail(emailid, newemail, newphone)
		fmt.Println("Email and Phone updated successfully")

	case 9:
		fmt.Println("Exit")
		break
	}
}

func main() {

	fmt.Println("1. Existing User")
	fmt.Println("2. New User")
	var userd int
	fmt.Println("Enter option here : ")
	fmt.Scanln(&userd)
	if userd == 1 {
		fmt.Println("Please Enter EmailID : ")
		var emailid string
		fmt.Scanln(&emailid)
		val := ValidateUser(emailid)
		if val == true {
			WelcomeConsole(emailid)
		} else {
			fmt.Println("User Not Found")
		}
	} else if userd == 2 {
		fmt.Println("Enter User Details")
		var name string
		var email string
		var mobileno string
		fmt.Println("Enter Name :")
		fmt.Scanln(&name)
		fmt.Println("Enter Mobile Number :")
		fmt.Scanln(&mobileno)
		fmt.Println("Enter EmailID :")
		fmt.Scanln(&email)
		var users User
		userdetails := UpdateNewUserDetails(name, mobileno, email)
		if userdetails.EmaidId != "" {
			UpdateDatabase(userdetails, users)
		}
		WelcomeConsole(email)
	}

}
