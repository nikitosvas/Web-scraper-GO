

package main

import (
	"log" 

	"Web-parsing/internal/scraper" 
	"github.com/spf13/viper"       

func main() { 
	viper.SetConfigFile("config.yaml") 
	viper.SetConfigType("yaml")       
	viper.AddConfigPath(".")           

	if err := viper.ReadInConfig(); err != nil { 
		log.Fatal("Error reading config file, ", err) 	}

	scraper.Run()
}
