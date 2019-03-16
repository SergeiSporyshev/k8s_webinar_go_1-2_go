package main

import (
	"os"
	"time"
)


func main(){
	file, err := os.OpenFile("logs/log.log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	t := time.Now()

	bytes := []byte("This is cron " + t.Format(time.RFC822) + "\n")
	_, errWrite := file.Write(bytes)
	if errWrite != nil{
		panic(errWrite)
	}
}