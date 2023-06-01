package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func main() {

	viper.SetConfigFile("config.json")
	err := viper.ReadInConfig()
	if err !=nil{
		fmt.Printf("Error: %v", err)
		return
	}

	api_key := viper.GetString("api_key")
	fmt.Printf("api_key: %v",api_key)

}