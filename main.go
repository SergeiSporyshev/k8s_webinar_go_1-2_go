package main

import (
	"fmt"
	"net/http"
	"os"
)

var dbHost, dbPort, dbUser, dbPassword, logLevel, podIp string

func init()  {
	dbHost = os.Getenv("DB_HOST")
	dbPort = os.Getenv("DB_PORT")
	dbUser = os.Getenv("DB_USER")
	dbPassword = os.Getenv("DB_PASSWORD")
	logLevel = os.Getenv("LOG_LEVEL")
	podIp = os.Getenv("POD_IP")
}

func handle(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello world!\n")
	fmt.Fprintf(w, "%s:%s \nuser: %s, password: %s\nlog level: %s", dbHost, dbPort, dbUser, dbPassword, logLevel)
	file, err := os.OpenFile("logs/log.log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil{
		panic(err)
	}

	defer file.Close()

	bytes := []byte("new request to " + podIp + "!\n")
	_, errWrite := file.Write(bytes)
	if errWrite != nil {
		panic(errWrite)
	}
}


func main()  {

	file, err := os.OpenFile("logs/log.log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil{
		panic(err)
	}

	defer file.Close()

	bytes := []byte("Start " + podIp + "!\n")
	_, errWrite := file.Write(bytes)
	if errWrite != nil {
		panic(errWrite)
	}

	http.HandleFunc("/check", handle)
	http.ListenAndServe(":4000", nil)
}