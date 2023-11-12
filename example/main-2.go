package main

import (
	"fmt"
	"time"
)

func GetMember() {
	fmt.Println("Please wait ...")
	time.Sleep(time.Second)
	CheckName("Taliw", "P@ssw0rd")
}

func IsInSystem(username string) bool  {
	return true
}


func GetUserDetail(username string) (int, string)  {
	return 201, "manager"

}

func getDepartment(username string, department *string)  {
	if username != "" {
		*department = "home"
	}
}
func CheckName(username string, password string)  {
	if IsInSystem(username) {
		fmt.Println("Found user in system")
		status, role := GetUserDetail(username)
		fmt.Println(fmt.Sprintf("Status = %d, Role = %s", status, role))
		department := ""
		getDepartment(username, &department)
		fmt.Println(fmt.Sprintf("Username = %s, Department = %s", username, department))
	} else {
		fmt.Println("User not found")
	}
}

func LogEnd()  {
	fmt.Println("Completed program...")
	fmt.Println(time.Now())
}

func CheckServerResponse()  {
	fmt.Println("Check server time")
	time.Sleep(3*time.Second)
	panic("SERVER ERROR!")
}

func main2() {
	defer func () {
		fmt.Println("Anonymous Function")
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	
	defer LogEnd()

	GetMember()
	CheckServerResponse()
}
