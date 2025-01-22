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
	Id2		int			`json:"Id"`
	Name2		string			`json:"Name"`
	Directions2 	[]string		`json:"Directions"`
	Address2	struct {
		Number2		int		`json:"Number"`
		Street2		string		`json:"Street"`
		City2		string		`json:"City"`
	}					`json:"Address"`
	Countries2	[]struct {		
		Name2		string		`json:"Name"`
		Language2	string		`json:"Language"`
		Greeting2	string		`json:"Greeting"`
	}					`json:"Countries"`
	Toggle2		bool			`json:"Toggle"`

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
	fmt.Println("Id:", AppConfig.Id2)
	fmt.Println("Name:", AppConfig.Name2)
	fmt.Println("Directions:")
	for _, direction := range AppConfig.Directions2 {
		fmt.Println("-", direction)
	}
	fmt.Println("Address:")
	fmt.Printf("- %d %s, %s.\n", AppConfig.Address2.Number2, AppConfig.Address2.Street2, AppConfig.Address2.City2)
	fmt.Println("Countries:")
	for _, country := range AppConfig.Countries2 {
		fmt.Printf("- In %s they speak %s and say '%s!'.\n", country.Name2, country.Language2, country.Greeting2)
	}
	fmt.Printf("Toggle: %t \n", AppConfig.Toggle2)
}


