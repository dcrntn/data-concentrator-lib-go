package dconc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type DataConcentrator struct {
	Ip string
}

type DataNode struct {
	Node_uid          string    `json:"node_uid"`
	Node_val          string    `json:"node_val"`
	Node_last_update  MongoDate `json:"node_last_update"`
	Node_name         string    `json:"node_name"`
	Node_rw_direction string    `json:"node_rw_direction"`
}

type DataNodeValue struct {
	Node_uid string `json:"node_uid"`
	Node_val string `json:"node_val"`
}

type MongoDate struct {
	Nate MongoDateVal `json:"$date"`
}

type MongoDateVal struct {
	NumberLong string `json:"$numberLong"`
}

// Gets all the data nodes from the server.
func (dc *DataConcentrator) GetAllDnode() []DataNode {

	dnodeArr := []DataNode{}
	resp, err := http.Get(fmt.Sprintf("%s/getall/bucket", dc.Ip))

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	jsonErr := json.Unmarshal(body, &dnodeArr)
	if jsonErr != nil {
		log.Fatal(err)
	}

	return dnodeArr
}

// Gets the value for a specific data node
func (dc *DataConcentrator) GetDnodeValue(uid string) string {

	dnodeVal := DataNodeValue{}

	resp, err := http.Get(fmt.Sprintf("%s/r/%s", dc.Ip, uid))

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	jsonErr := json.Unmarshal(body, &dnodeVal)
	if jsonErr != nil {
		log.Fatal(err)
	}

	return dnodeVal.Node_val
}

// Write a value for a specific data node
func (dc *DataConcentrator) WriteDnodeValue(uid string, value string) string {
	url := fmt.Sprintf("%s/w", dc.Ip)
	dnodeValSend := DataNodeValue{
		Node_uid: uid,
		Node_val: value,
	}
	jsonStr, err := json.Marshal(dnodeValSend)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("X-Custom-Header", "*")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	return string(body)
}
