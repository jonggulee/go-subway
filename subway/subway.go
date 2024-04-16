package subway

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

const (
	realtimePosition = "realtimePosition"
	resultType       = "json"
	startIndex       = 1
	endIndex         = 9
	updnLine         = 0
	encodedLine      = "8%ED%98%B8%EC%84%A0"
	baseUrl          = "http://swopenapi.seoul.go.kr/api/subway"
)

type JsonResponse struct {
	ErrorMessage         ErrorMessage       `json:"errorMessage"`
	RealtimePositionList []RealtimePosition `json:"realtimePositionList"`
}

type ErrorMessage struct {
	Status           int    `json:"status"`
	Code             string `json:"code"`
	Message          string `json:"message"`
	Link             string `json:"link"`
	DeveloperMessage string `json:"developerMessage"`
	Total            int    `json:"total"`
}

type RealtimePosition struct {
	SubwayId     string `json:"subwayId"`
	SubwayNm     string `json:"subwayNm"`
	StatnId      string `json:"statnId"`
	StatnNm      string `json:"statnNm"`
	TrainNo      string `json:"trainNo"`
	LastRecptnDt string `json:"lastRecptnDt"`
	RecptnDt     string `json:"recptnDt"`
	UpdnLine     string `json:"updnLine"`
	StatnTid     string `json:"statnTid"`
	StatnTnm     string `json:"statnTnm"`
	TrainSttus   string `json:"trainSttus"`
	DirectAt     string `json:"directAt"`
	LstcarAt     string `json:"lstcarAt"`
}

func GetRealtimePosition() []RealtimePosition {
	url := fmt.Sprintf("%s/%s/%s/%s/%d/%d/%s", baseUrl, apiKey, resultType, realtimePosition, startIndex, endIndex, encodedLine)

	resp, err := http.Get(url)
	if err != nil {
		log.Println("Error:", err)
	}

	defer resp.Body.Close()

	jsonData, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
	}

	var data JsonResponse
	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		fmt.Println("Error:", err)
	}

	var resultData []RealtimePosition
	if data.ErrorMessage.Status == 200 {
		for _, subway := range data.RealtimePositionList {
			intStantId, _ := strconv.Atoi(subway.StatnId)
			if subway.UpdnLine == "0" && intStantId >= 1008000821 {
				resultData = append(resultData, subway)
			}
		}
	}
	return resultData
}
