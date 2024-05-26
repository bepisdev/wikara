package main

import (
	"log"
	"fmt"
	"os"
	"path/filepath"
	"net/http"
	"github.com/spf13/viper"
	"github.com/joshburnsxyz/wikara/pkg/api"
	"github.com/joshburnsxyz/wikara/pkg/templates"
)

func getExecPath() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	return exPath
}

func main() {
	// Set up config defaults
	viper.SetDefault("Port", "8080")
	viper.SetDefault("Host", "0.0.0.0")

	// Load config file
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(getExecPath())
	if err := viper.ReadInConfig(); err != nil {
	    if _, ok := err.(viper.ConfigFileNotFoundError); ok {
		    log.Println("No config file found, Running with default configuration")
	    } else {
		    log.Fatal("Error in config file %w", err)
	    }
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
