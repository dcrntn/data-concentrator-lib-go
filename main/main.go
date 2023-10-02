package main

import (
	"dataconcentrator/dconc"
	"fmt"
)

func main() {
	myDc := dconc.DataConcentrator{Ip: "http://127.0.0.1:8000"}

	fmt.Println("getAllDnode(): ", myDc.GetAllDnode())
	fmt.Println("writeDnodeValue(): ", myDc.WriteDnodeValue("ayH7nFocwL2urRvOQOfQ", "532"))
	fmt.Println("getDnodeValue(): ", myDc.GetDnodeValue("ayH7nFocwL2urRvOQOfQ"))
}
