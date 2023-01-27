package main

import (
	"fmt"
	"net/http"
	"bytes"
	"io"
)

func update_details(){
	c:=http.Client{}
	update_Json := bytes.NewBuffer([]byte(`{"id":"1","balance":"updated_balance"}`))
	req,_:= http.NewRequest("PUT","http://localhost:6000/update_employee_by_id",update_Json)
	req.Header.Add("Accept",`application/json`)
	resp,err := c.Do(req)
	if err!= nil{
		fmt.Println("Error in sending request",err.Error())
		return
	}
	defer resp.Body.Close()
	body,err := io.ReadAll(resp.Body)
	if  err!= nil{
		fmt.Println("Error in reading response",err.Error())
	}
	fmt.Printf("Body: %s\n",string(body))
	fmt.Printf("Status code: %d\n",resp.StatusCode)
}