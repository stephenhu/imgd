package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
)

const (
	APP_CONF_FILE            	= ".imgd.json"
	APP_NAME									= "imgd"
	APP_VERSION								= "1.0"
)

type RedisConf struct {
	Host					string					`json:"host"`
	Port					string					`json:"port"`
}

type ImgdConf struct {
	Redis					RedisConf				`json:"redis"`
	Address       string          `json:"address"`
}

var client 		*redis.Client 	= nil

var conf = ImgdConf{}


func appLog(msg string) {
	log.Printf("%s", msg)
} // appLog


func appLogd(fn string, msg string) {
	log.Printf("%s(): %s", fn, msg)
} // appLogd


func parseConfig() {

	_, err := os.Stat(APP_CONF_FILE)
	
	if err != nil || os.IsNotExist(err) {
		log.Fatalf("%s not found.", APP_CONF_FILE)
	} else {

		buf, err := ioutil.ReadFile(APP_CONF_FILE)

		if err != nil {
			log.Fatalf("Unable to read contents of %s: %s", APP_CONF_FILE, err.Error())
		} else {

			err := json.Unmarshal(buf, &conf)

			if err != nil {
				log.Fatalf("Unable to parse %s: %s", APP_CONF_FILE, err.Error())
			}

		}

	}

} // parseConfig


func initRoutes() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/crawlers", crawlerHandler)
	router.HandleFunc("/api/pics", picsHandler)
	
	return router

} // initRoutes


func main() {

	parseConfig()

	r := Red{}

	r.Connect()

	defer client.Close()

	go productJob()

	log.Println(fmt.Sprintf("Listening on port %s", conf.Address))
	log.Fatal(http.ListenAndServe(conf.Address, initRoutes()))

} // main
