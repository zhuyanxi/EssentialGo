package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func region(ip ...string) string {
	method := "POST"

	ipStr, _ := json.Marshal(ip)

	//payload := strings.NewReader("[\"10.132.123.220\",\"10.134.41.91\"]")
	payload := strings.NewReader(string(ipStr))

	client := &http.Client{}
	req, err := http.NewRequest(method, "http://localhost:8081/iplocations", payload)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Query ip location failed:", err)
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Query ip read failed:", err)
	}

	fmt.Println(string(body))
	return ""
}
