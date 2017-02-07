package trafficbot

import (
	"net/http"
	// "io/ioutil"
	"encoding/json"
	"googlemaps.github.io/maps"
	// "github.com/kr/pretty"
    "golang.org/x/net/context"
    "log"    
    "time"
)

const trafficEventURL = "http://rtr.pbs.gov.tw/pbsmgt/RoadAllServlet?ajaxAction=roadAllCache"

type TrafficResponse struct {
	FormData []Event `json:"formData"`
}

type Event struct {
	Srcdetail    string `json: "srcdetail"`
	Region       string `json:"region"`
	Highway      string `json:"highway"`
	Updatetime   string `json:"updatetime"`
	Direction    string `json:"direction"`
	Lastmodified string `json:"lastmodified"`
	Tokm         string `json:"tokm"`
	Continuedate string `json:"continuedate"`
	Speedlow     string `json:"speedlow"`
	Level        string `json:"level"`
	Area_sn      string `json:"area_sn"`
	Road_bak2    string `json:"road_bak2"`
	Road_bak1    string `json:"road_bak1"`
	Name         string `json:"name"`
	To1          string `json:"to1"`
	To2          string `json:"to2"`
	Roadtype     string `json:"roadtype"`
	Fromkm       string `json:"fromkm"`
	Updater      string `json:"updater"`
	Happentime   string `json:"happentime"`
	Number       string `json:"number"`
	Canceldate   string `json:"canceldate"`
	Speedtop     string `json:"speedtop"`
	From2        string `json:"from2"`
	Active       string `json:"active"`
	Continuetime string `json:"continuetime"`
	Comment      string `json:"comment"`
	Updatedate   string `json:"updatedate"`
	From1        string `json:"from1"`
	Happendate   string `json:"happendate"`
}

func getJson(url string, target interface{}) error {

	client := &http.Client{} // or new(http.Client)
	reqest, _ := http.NewRequest("GET", url, nil)
	reqest.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	reqest.Header.Set("Accept-Charset", "zh-tw,utf-8;q=0.7,*;q=0.3")
	reqest.Header.Set("Accept-Encoding", "gzip,deflate,sdch")
	reqest.Header.Set("Accept-Language", "zh-TW,zh;q=0.8")
	reqest.Header.Set("Cache-Control", "max-age=0")
	reqest.Header.Set("Connection", "keep-alive")

	response, err := client.Do(reqest)
	if response.StatusCode == 200 {
		return json.NewDecoder(response.Body).Decode(&target)
	}
	return err

}

func GetTrafficEvents() (*TrafficResponse, error) {
	trafficResponse := &TrafficResponse{}
	err := getJson(trafficEventURL, trafficResponse)
	if err != nil {
		return nil, err
	}
	return trafficResponse, nil	
}

func getTrafficDuration(origin string, destination string) time.Duration {

	c, err := maps.NewClient(maps.WithAPIKey("AIzaSyBGHsozp7Ua2p_KMidEB710DXMyfBQP0R8"))
    if err != nil {
        log.Fatalf("fatal error: %s", err)
    }

    r := &maps.DirectionsRequest{
        Origin:      origin,
        Destination: destination,        
        TrafficModel: maps.TrafficModelBestGuess,
        Language: "zh-TW",
        Units: "metric",
        DepartureTime: "now",
    }

    resp, _, err := c.Directions(context.Background(), r)
    if err != nil {
        log.Fatalf("fatal error: %s", err)
    }    
	
	return resp[0].Legs[0].DurationInTraffic; 
}

func GetDurationFromHomeToAirport() string {
	origin := "24.934725,121.3990663";
	destination := "25.077713,121.2305153";
	duration := getTrafficDuration(origin, destination).String();
	return duration;
}

func GetDurationFromAirportToHome() string {
	origin := "25.077713,121.2305153";
	destination := "24.934725,121.3990663";
	duration := getTrafficDuration(origin, destination).String();
	return duration;
	   
}