package main

import (
	"fmt"
	"io"
	"net/http"
	"bytes"
)

func delete_details(){
	c:= http.Client{}
	delete_Json := bytes.NewBuffer([]byte(`{"id":"3"}`))
	req, _ := http.NewRequest("DELETE","http://localhost:6000/delete_employee_byid",delete_Json)
	req.Header.Add("Accept",`application/json`)
	resp,err := c.Do(req)
	if err!=nil{
		fmt.Println("Error in sending request",err.Error())
		return
	}
	defer resp.Body.Close()
	body,err := io.ReadAll(resp.Body)
	if err != nil{
		fmt.Println("Error in reading response",err.Error())
	}
	fmt.Printf("Body: %s\n",string(body))
	fmt.Printf("Status code:%d\n",resp.StatusCode)
}