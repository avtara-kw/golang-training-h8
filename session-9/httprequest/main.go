package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	url := "https://jsonplaceholder.typicode.com"
	post(url + "/posts")
	get(url + "/posts/100")
}

func get(url string) {
	res, err := httpcall(url, "GET", []byte{})
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res["title"])
}

func post(url string) {
	data := map[string]interface{}{
		"title":  "Hello from hacktiv8",
		"body":   "Ini adalah testing request method post",
		"userId": 1,
	}

	requestByte, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	res, err := httpcall(url, "POST", requestByte)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res["title"])

}

func httpcall(url, method string, body []byte) (map[string]interface{}, error) {

	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return map[string]interface{}{}, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		return map[string]interface{}{}, err
	}
	defer res.Body.Close()

	var response map[string]interface{}

	err = json.NewDecoder(res.Body).Decode(&response)

	return response, err
}
