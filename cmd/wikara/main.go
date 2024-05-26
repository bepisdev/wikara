package main

import (
	"log"
	"fmt"
	"net/http"
	"github.com/spf13/viper"
	"github.com/joshburnsxyz/wikara/pkg/api"
	"github.com/joshburnsxyz/wikara/pkg/templates"
)

func main() {
	// Set up config defaults
	viper.SetDefault("Port", "8080")
	viper.SetDefault("Host", "0.0.0.0")

	// Load config file
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(fmt.Errorf("Fatal error config file: %w", err))
	}

	// Set up routes
	http.HandleFunc("/", api.FrontPageHandler)
	http.HandleFunc("/view/", api.MakeHandler(api.ViewHandler))
	http.HandleFunc("/edit/", api.MakeHandler(api.EditHandler))
	http.HandleFunc("/save/", api.MakeHandler(api.SaveHandler))

	// Init templates cache
	templates.Init()
	
	bindAddr := fmt.Sprintf("%s:%s", viper.GetString("Host"), viper.GetString("Port"))
	log.Println(fmt.Sprintf("Server started on %s", bindAddr))
	log.Fatal(http.ListenAndServe(bindAddr, nil))
}
