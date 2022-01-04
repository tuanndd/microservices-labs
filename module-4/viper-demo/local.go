package main

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/viper"
)

func main() {
	// https://medium.com/@jomzsg/the-easy-way-to-handle-configuration-file-in-golang-using-viper-6b3c88d2ee79

	// read config file
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config/")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("fatal error config file: local", err)
		os.Exit(1)
	}

	// get config vars

	viper.SetDefault("app.linetoken", "DefaultLineTokenValue")

	env := viper.GetString("app.env")
	producerbroker := viper.GetString("app.producerbroker")
	consumerbroker := viper.GetString("app.consumerbroker")
	linetoken := viper.GetString("app.linetoken")

	// print conf vars
	fmt.Println("app.env :", env)
	fmt.Println("app.producerbroker :", producerbroker)
	fmt.Println("app.consumerbroker :", consumerbroker)
	fmt.Println("app.linetoken :", linetoken)

	time.Sleep(300)
}
