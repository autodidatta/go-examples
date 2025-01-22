/*
Read JSON config file into a struct.
*/

package main

import (
	"fmt"
	"io/ioutil"
	"encoding/json"
)

type Cfg struct {
	Id		int
	Name		string
	Directions 	[]string
	Address		struct {
		Number		int
		Street		string
		City		string
	}
	Countries	[]struct {
		Name		string
		Language	string
		Greeting	string
	}
	Toggle		bool
}

var AppConfig Cfg

func init() {
    raw, err := ioutil.ReadFile("jsoncfg.json")
    if err != nil {
       fmt.Println("Error occured while reading config")
       return
    }
    json.Unmarshal(raw, &AppConfig)
}

func main() {
	fmt.Println("Id:", AppConfig.Id)
	fmt.Println("Name:", AppConfig.Name)
	fmt.Println("Directions:")
	for _, direction := range AppConfig.Directions {
		fmt.Println("-", direction)
	}
	fmt.Println("Address:")
	fmt.Printf("- %d %s, %s.\n", AppConfig.Address.Number, AppConfig.Address.Street, AppConfig.Address.City)
	fmt.Println("Countries:")
	for _, country := range AppConfig.Countries {
		fmt.Printf("- In %s they speak %s and say '%s!'.\n", country.Name, country.Language, country.Greeting)
	}
	fmt.Printf("Toggle: %t \n", AppConfig.Toggle)
}


