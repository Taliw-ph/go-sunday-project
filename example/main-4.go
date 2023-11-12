package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Address struct {
	Address  string
	PostCode string
}

type UserProfile struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Age       int
	Height    float32
	Address   Address
	Bill      struct {
		BillAddress string
	}
}

// Implement function to strunct
func (u UserProfile) ToFullname() string {
	return fmt.Sprintf("%s %s", u.Firstname, u.Lastname)
}

func main4() {
	fmt.Println("Map and Struct")

	// Map
	user := map[string]string{}
	user["username"] = "phanuwat"
	user["password"] = "P@ssw0rd"
	fmt.Println(user)
	fmt.Println(user["username"])

	// Struct
	userProfile := UserProfile{
		Firstname: "Phanuwat",
		Lastname:  "Phoowichai",
		Age:       23,
	}
	fmt.Println(userProfile)
	userProfile.Address.PostCode = "10900"

	fmt.Println(userProfile.ToFullname())

	// struct to json
	byteTxtJson, err := json.MarshalIndent(userProfile, "", "  ")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(byteTxtJson))

	// json to struct
	dataJson := `{
		"firstname": "Taliw",
		"lastname": "Degile",
		"Age": 23,
		"Address": {
		  "Address": "Geeky Base",
		  "PostCode": "10900"
		},
		"Bill": {
		  "BillAddress": "Bangkok"
		}
	  }`
	var taliwProfile UserProfile
	err = json.Unmarshal([]byte(dataJson), &taliwProfile)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(taliwProfile)
	fmt.Println(taliwProfile.ToFullname())
}
