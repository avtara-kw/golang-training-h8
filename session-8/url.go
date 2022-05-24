package main

import (
	"fmt"
	"net/url"
)

func main() {
	urlString := "https://www.kode.id/courses?name=scalable-services-golang&level=intermediate"

	u, err := url.Parse(urlString)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Host", u.Host)
	fmt.Println("Query", u.Query())
	fmt.Println("Schema", u.Scheme)
	fmt.Println("Path", u.Path)
	fmt.Println("Path", u.RawQuery)
}
