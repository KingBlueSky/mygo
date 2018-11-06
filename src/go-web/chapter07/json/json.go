package main

import (
	"encoding/json"
	"fmt"
)

type Server struct {
	ServerName string
	ServerIP   string
}

type ServerSlice struct {
	Servers []Server
}

func main() {
	var s ServerSlice
	str := `{"servers":[{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1"},{"serverName":"Beijing_VPN","serverIP":"127.0.0.2"}]}`
	json.Unmarshal([]byte(str), &s)
	fmt.Println(s)

	b := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)
	var f interface{}
	err := json.Unmarshal(b, &f)
	fmt.Println(err, f)

	for k, v := range f.(map[string]interface{}) {
		switch vt := v.(type) {
		case string:
			fmt.Println(k, "is string", vt)
		case int:
			fmt.Println(k, "is int", vt)
		case float64:
			fmt.Println(k, "is float64", vt)
		case []interface{}:
			fmt.Println(k, "is an array:")
			for i, u := range vt {
				fmt.Println(i, u)
			}
		default:
			fmt.Println(k, "is of a type I don't know how to handle")
		}
	}

}
