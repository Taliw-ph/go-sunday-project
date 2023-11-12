package main

import (
	"fmt"
)

type SaveLogInterface interface {
	SaveLog()
}

func SaveLog(st SaveLogInterface)  {
	st.SaveLog()
}

type TransferBBL struct {
	name string
}

func (tBBL *TransferBBL) SaveLog()  {
	tBBL.name = "test"
	fmt.Println("Save to database", tBBL.name)
}

type TransferKTB struct {
	name string
}

func (tKTB TransferKTB) SaveLog()  {
	fmt.Println("Save to database", tKTB.name)
}

func main5() {
	fmt.Println("Interface")

	// Interface
	transA := TransferBBL{ name: "BBL" }
	transB := TransferKTB{ name: "TTB" }
	SaveLog(&transA)
	fmt.Println(transA)
	SaveLog(transB)
}
