package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Employee struct {
	ID       int
	Name     string
	Age      int
	Position string
}

func main() {
	var employe []Employee
	file, err := os.Open("employees.json")
	if err != nil {
		log.Fatal(err)
	}

	emp := json.NewDecoder(file)
	err2 := emp.Decode(&employe)
	if err2 != nil {
		log.Fatal(err2)
	}
	defer file.Close()
	for i,v := range employe{
		if v.ID == 3{
			v.Age =50
			employe[i]=v
		}
	}
	e1 := Employee{
		ID: 6,
		Name: "John Brown",
		Age: 35,
		Position: "Software Engineer",
	}
	employe = append(employe, e1)
	file1,err := os.OpenFile("employees_new.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC,0666)
	if err != nil{
		log.Fatal(err)
	}
	data1 := json.NewEncoder(file1)
	err = data1.Encode(employe)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Muvaffaqiyatli bajarildi!")

}
