package main

import (
	"fmt"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("config")                   // name of config file (without extension)
	viper.AddConfigPath("./viper-practice/config/") // path to look for the config file in
	err := viper.ReadInConfig()                     // Find and read the config file
	if err != nil {                                 // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	fmt.Println(viper.GetInt("age"))
}
