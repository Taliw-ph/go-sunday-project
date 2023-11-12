package main

import (
	"fmt"
)

func main1() {	
	// fmt.Println("Hello! สวัสดี!")

	// // Non assign value
	// // var txtInput string

	// // Assigned value
	// txtInput := ""

	// fmt.Print("Enter name: ")
	// fmt.Scan(&txtInput)
	// fmt.Println("Hello, " + txtInput + ".")

	// age := 2
	// fmt.Println(reflect.TypeOf(age))

	// // convert beetween int, float and string
	// height := "100"

	// h, err := strconv.Atoi(height)
	// if err != nil {
	// 	fmt.Println("Error: " + err.Error())
	// }
	// fmt.Println(h)

	// weather := "12.40"
	// fWeather, err := strconv.ParseFloat(weather, 64)
	// if err != nil {
	// 	fmt.Println("Error: " + err.Error())
	// }
	// fmt.Println(fWeather)

	// txtShow := "" // var txtShow string = ""

	// txtShow = fmt.Sprintf("Weather = %3.2f, Height = %d", fWeather, h)

	// fmt.Println(txtShow)

	for i := 0; i < 10; i++ {
		// fmt.Println(i)
	}

	score := 55

	switch score {
	case 80:
		fmt.Println("A")
	case 70:
		fmt.Println("B")
	}

	if score >= 80 {
		fmt.Println("A")
	} else if score < 80 && score >= 70 {
		fmt.Println("B")
	} else {
		fmt.Println("Unknown")
	}
}
