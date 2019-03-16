package main

import "os"


func main(){
	file, err := os.OpenFile("logs/log.log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	bytes := []byte("This is job")
	_, errWrite := file.Write(bytes)
	if errWrite != nil{
		panic(errWrite)
	}
}