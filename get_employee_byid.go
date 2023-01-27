package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func get_employee_byid(){
	c:=http.Client{}
	inp_Json := bytes.NewBuffer([]byte(`{"id":"3"}`))
	req,_ := http.NewRequest("GET", "http://localhost:6000/get_employee_byid", inp_Json)
	req.Header.Add("Accept",`application/json`)
	resp,err := c.Do(req)
	if err !=nil{
		fmt.Println("Error in sending request",err.Error())
		return
	}
	defer resp.Body.Close()
	body,err :=io.ReadAll(resp.Body)
	if err!=nil{
		fmt.Println("Error in reading response",err.Error())
	}
	fmt.Printf("Body:%s\n",string(body))
	fmt.Printf("Status code:%d\n",resp.StatusCode)
}