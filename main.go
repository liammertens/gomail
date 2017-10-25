package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"encoding/json"
	"gopkg.in/gomail.v2"
)

// main 
func main()  {
	// TODO: add a checker for config.json
	initialSetup()
}
// initialSetup check whether user want to go on
func initialSetup()  {
	var selection string

	fmt.Println("Do you want to send an e-mail? (y/n)")
	fmt.Scanln(&selection)
	
	if selection == "y" || selection == "Y" {
		_, err := os.Stat("config.json")
		if !os.IsNotExist(err) { //check whether config.json exists
			fmt.Printf("User config found...\n")
			constructMessage()
		} else if err == err.(*os.PathError) { //if config.json doesn't exist -> create
			fmt.Printf("No user config found: creating...\n")
//			constructConfig()
			constructUser()
		}else {
 			fmt.Printf("%v\n", err)
		} 
		
	} else if selection == "n" || selection == "N" {
		os.Exit(0)
	} else {
		fmt.Println("Command not recognized")
	}
}

func check(e error) { //easier to check for errors
	if e != nil {
		panic(e)
	}
}

// constructConfig
/*func constructConfig()  {
	var file, err = os.Create("config.json")
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	file.Close()
}*/

type User struct {
	Mail string
	ServerAdress string
	ServerPort int     
	UserName string    
	Password string
}
// constructUser 
func constructUser()  {
	var userMailAdress string
	var serverAdress string
	var serverPort int
	var userName string
	var password string

	fmt.Printf("What is your e-mail adress?\n")
	fmt.Scanln(&userMailAdress)

	fmt.Printf("What is your smtp server adress?\n")
	fmt.Scanln(&serverAdress)

	fmt.Printf("What is your smtp port?\n")
	fmt.Scanln(&serverPort)

	fmt.Printf("What is your username?\n")
	fmt.Scanln(&userName)

	fmt.Printf("What is your password?\n")
	fmt.Scanln(&password)

	user := User{Mail: userMailAdress, ServerAdress: serverAdress, ServerPort: serverPort, UserName: userName, Password: password}
	
	jsonInput, err := json.Marshal(user)
	check(err)

	fileError := ioutil.WriteFile("config.json", jsonInput, 0644)
	check(fileError)
	fmt.Printf("Wrote e-mail to config\n")
	constructMessage()
}


// constructMessage 
func constructMessage()  {
	var user User
	var receiver string
	var subject string
	var body string

	userData , err := ioutil.ReadFile("config.json")
	check(err)
	json.Unmarshal(userData, &user)
	
	m := gomail.NewMessage()

	m.SetHeader("From", user.Mail)

	fmt.Printf("To whom are you sending the message? (1 person)\n")
	fmt.Scanln(&receiver)
	m.SetHeader("To", receiver)

	fmt.Printf("What is the subject?\n")
	fmt.Scanln(&subject)
	m.SetHeader("Subject", subject)
	
	fmt.Printf("What is your message\n")
	fmt.Scanln(&body)
	m.SetBody("text/plain", body)

	d := gomail.NewDialer(user.ServerAdress, user.ServerPort, user.UserName, user.Password)
	e := d.DialAndSend(m)
	check(e)
}

