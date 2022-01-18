package main

import (
	"fmt"
	"go-clean-architecture/config"
	load "go-clean-architecture/init"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

func init() {
	load.RunInit()
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	headers := handlers.AllowedHeaders([]string{"Origin", "Accept", "Keep-Alive", "User-Agent", "If-Modified-Since", "Cache-Control", "Referer", "Authorization", "Content-Type", "X-Requested-With"})
	origins := handlers.AllowedOrigins([]string{"*"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS", "DELETE", "PUT", "HEAD"})

	config.Setup(router)

	fmt.Println("Server running at :" + viper.GetString("server.port"))
	log.Fatal(http.ListenAndServe(":"+viper.GetString("server.port"), handlers.CORS(origins, headers, methods)(router)))
}
