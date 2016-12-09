package main

import (
	"github.com/kynwu/trafficbot"
	"fmt"
)

func main() {

	// trafficRes := new(trafficbot.TrafficResponse)
	trafficRes, err := trafficbot.GetTrafficEvents()	
	if err != nil {
		fmt.Println(err)
	}
	
	for _, element := range trafficRes.FormData {
		if element.Highway == "1001" {
			fmt.Println(element.Name)
		}
	}
}