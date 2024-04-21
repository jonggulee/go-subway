package subway

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

const (
	realtimeStationArrival = "realtimeStationArrival"
	resultType             = "json"
	startIndex             = 0
	endIndex               = 1
	updnLine               = 0
	baseUrl                = "http://swopenapi.seoul.go.kr/api/subway"
	apiKey                 = "69515253776a6a6f3637664a744578"
	statnBokjeong          = "복정"
	statnNamwirye          = "남위례"
)

type JsonResponse struct {
	ErrorMessage        ErrorMessage          `json:"errorMessage"`
	RealtimeArrivalList []realtimeArrivalList `json:"realtimeArrivalList"`
}

type ErrorMessage struct {
	Status           int    `json:"status"`
	Code             string `json:"code"`
	Message          string `json:"message"`
	Link             string `json:"link"`
	DeveloperMessage string `json:"developerMessage"`
	Total            int    `json:"total"`
}

type realtimeArrivalList struct {
	UpdnLine string `json:"updnLine"`
	ArvlMsg2 string `json:"arvlMsg2"`
	SubwayId string `json:"subwayId"`
}

type Subway struct {
	SubwayNm string `json:"subwayNm"`
	Statn    string `json:"statn"`
	ArvlMsg  string `json:"arvlMsg"`
}

func GetRealtimeStationArrival(station string) Subway {
	var subway Subway

	url := fmt.Sprintf("%s/%s/%s/%s/%d/%d/%s", baseUrl, apiKey, resultType, realtimeStationArrival, startIndex, endIndex, station)
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error fetching data for station %s: %v", station, err)

	}
	defer resp.Body.Close()

	jsonData, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response body for station %s: %v", station, err)
	}

	var data JsonResponse
	if err := json.Unmarshal(jsonData, &data); err != nil {
		log.Printf("Error unmarshalling JSON for station %s: %v", station, err)
	}

	for _, arrival := range data.RealtimeArrivalList {
		if arrival.UpdnLine == "상행" {
			if arrival.SubwayId == "1008" {
				subway := Subway{SubwayNm: "8호선", Statn: station, ArvlMsg: arrival.ArvlMsg2}
				return subway
			} else if arrival.SubwayId == "1075" {
				subway := Subway{SubwayNm: "수인분당선", Statn: station, ArvlMsg: arrival.ArvlMsg2}
				return subway
			}
		}
	}
	return subway
}
