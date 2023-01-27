package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func post_details(){
	c:=http.Client{}
	test_Json := bytes.NewBuffer([]byte(`{"id":"4","name":"test_name","balance":"188425"}`))
	req,_:= http.NewRequest("POST","http://localhost:6000/post_employee",test_Json)
	req.Header.Add("Accept", `application/json`)
	resp,err :=c.Do(req)
	if err!=nil{
		fmt.Println("Error in sending request",err.Error())
		return
	}
	defer resp.Body.Close()
	body,err := io.ReadAll(resp.Body)
	if err !=nil{
		fmt.Println("Error in reading response",err.Error())
	}
	fmt.Printf("Body: %s\n",string(body))
	fmt.Printf("Status code: %d\n",resp.StatusCode)
}	