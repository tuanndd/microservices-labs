package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"

	_ "github.com/spf13/viper/remote"
)

func main() {
	// ref: https://learnku.com/articles/33908
	v := viper.New()

	v.AddRemoteProvider("consul", "localhost:8500", "config")
	v.SetConfigType("json")
	if err := v.ReadRemoteConfig(); err != nil {
		log.Println(err)
		return
	}

	fmt.Println("port: ", v.GetInt("port"))
	fmt.Println("mysql.url: ", v.GetString(`mysql.url`))
	fmt.Println("mysql.username: ", v.GetString(`mysql.username`))
	fmt.Println("mysql.password: ", v.GetString(`mysql.password`))
	fmt.Println("redis: ", v.GetStringSlice("redis"))
	fmt.Println("smtp: ", v.GetStringMap("smtp"))
}
