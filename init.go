package main

import "os"
var inPodIp string

func init(){
	inPodIp = os.Getenv("POD_IP")
}

func main(){
	file, err := os.OpenFile("logs/log.log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	bytes := []byte("Init container run " + inPodIp + "\n")
	_, errWrite := file.Write(bytes)
	if errWrite != nil{
		panic(errWrite)
	}
}