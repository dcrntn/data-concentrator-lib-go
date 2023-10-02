package dconc

import (
	"testing"
)

func TestGetAllDnode(t *testing.T) {
	myDc := DataConcentrator{Ip: "http://127.0.0.1:8000"}
	lenArr := len(myDc.GetAllDnode())
	neededLen := 2
	if lenArr != neededLen {
		t.Fatalf(`[TEST ERR] GetAllDnode() returned: %d objects | object in DB: %d`, lenArr, neededLen)
	}
}

func TestWriteDnodeValue(t *testing.T) {
	myDc := DataConcentrator{Ip: "http://127.0.0.1:8000"}
	dnodeResp := myDc.WriteDnodeValue("ayH7nFocwL2urRvOQOfQ", "41")
	neededResp := "\"{'changed_count': '1'}\""

	if dnodeResp != neededResp {
		t.Fatalf(`[TEST ERR] WriteDnodeValue() returned resp: %s | needed resp: %s`, dnodeResp, neededResp)
	}

}

func TestGetDnodeValue(t *testing.T) {
	myDc := DataConcentrator{Ip: "http://127.0.0.1:8000"}
	dnodeVal := myDc.GetDnodeValue("ayH7nFocwL2urRvOQOfQ")
	neededVal := "41"

	if dnodeVal != neededVal {
		t.Fatalf(`[TEST ERR] GetDnodeValue() returned value: %s | needed value: %s`, dnodeVal, neededVal)
	}

}
