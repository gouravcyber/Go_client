package main

import (
	"fmt"
	"io"
	"net/http"
//	"time"
)

func get_all(){
	c:= http.Client{}
	req,_ := http.NewRequest("GET","http://localhost:6000/get_all_employees",nil)
	req.Header.Add("Accept", `application/json`)
	resp,err:=c.Do(req)
	if err != nil{
		fmt.Println("Error sending request",err.Error())
        return
    }
    defer resp.Body.Close()
	body,err := io.ReadAll(resp.Body)
	if err!= nil{
		fmt.Println("Error reading response",err.Error())
	}
	fmt.Printf("Body: %s\n",string(body))
	fmt.Printf("Status code: %d\n",resp.StatusCode)
}