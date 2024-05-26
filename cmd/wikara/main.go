package main

import (
	"log"
	"fmt"
	"net/http"
	"github.com/spf13/viper"
	"github.com/joshburnsxyz/wikara/pkg/api"
	"github.com/joshburnsxyz/wikara/pkg/templates"
	"github.com/joshburnsxyz/wikara/pkg/utils"
)



func main() {
	// Set up config defaults
	viper.SetDefault("Port", "8080")
	viper.SetDefault("Host", "0.0.0.0")
	viper.SetDefault("SSL", false)
	viper.SetDefault("SSLCertFile", "./cert.pem")
	viper.SetDefault("SSLKeyFile", "./key.pem")
	viper.SetDefault("ContentDir", "data")
	viper.SetDefault("FrontPageTitle", "FrontPage")
	viper.SetDefault("SiteTitle", "Wikara")

	// Load config file
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(utils.GetExecPath())
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

	if viper.GetBool("SSL") {
		certFile := viper.GetString("SSLCertFile")
		keyFile := viper.GetString("SSLKeyFile")
		log.Println(fmt.Sprintf("Server started on https://%s", bindAddr))
		log.Fatal(http.ListenAndServeTLS(bindAddr, certFile, keyFile, nil))
	} else if !viper.GetBool("SSL"){
		log.Println(fmt.Sprintf("Server started on http://%s", bindAddr))
		log.Fatal(http.ListenAndServe(bindAddr, nil))
	} else {
		log.Fatal("Could not start server, please check config file.")
	}
}
