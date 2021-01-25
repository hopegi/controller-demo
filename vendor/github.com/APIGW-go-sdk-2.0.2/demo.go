package main

import (
	"APIGW-go-sdk-2.0.2/core"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	demoApp()
}

func demoApp() {
	s := core.Signer{
		Key:    "TKKQPOYGAGG0TYP2AHTY",
		Secret: "PxU3kXN1V2jYe7Am6kweep9n8vtUq9v1Qgqil17n",
	}
	r, _ := http.NewRequest("GET", "https://vpc.cn-south-1.myhuaweicloud.com/v1/afeee5482fd64101bc4d4511da06e02c/routetables", nil)
	r.Header.Add("content-type", "application/json")
	s.Sign(r)
	fmt.Println(r.Header)
	client := http.DefaultClient
	resp, err := client.Do(r)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(body))
}
