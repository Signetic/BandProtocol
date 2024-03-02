package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type jsonCode struct {
	Symbol    string `json:"symbol"`
	Price     uint64 `json:"price"`
	Timestamp uint64 `json:"timestamp"`
}

type respPOST struct {
	Tx_hash string `json:"tx_hash"`
}
type respGET struct {
	Tx_status string `json:"tx_status"`
}

func callApi(url string, method string, payload []byte) string {
	request, _ := http.NewRequest(method, url, bytes.NewBuffer(payload))
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return "Call API Error"
	}
	defer response.Body.Close()
	responseBody, _ := io.ReadAll(response.Body)
	return string(responseBody)
}

func RepInfo(resp string) string {
	switch resp {
	case "CONFIRMED":
		return `CONFIRMED : Transaction has been processed and confirmed`
	case "FAILED":
		return `FAILED : Transaction failed to process`
	case "PENDING":
		return `PENDING : Transaction is awaiting processing`

	case "DNE":
		return `DNE : Transaction does not exist`
	default:
		return `Error Response`
	}
}
func main() {
	solve := "GET"
	switch solve {
	case "POST":
		urlPost := "https://mock-node-wgqbnxruha-as.a.run.app/broadcast"
		payload := jsonCode{
			Symbol:    "ETH",
			Price:     4500,
			Timestamp: 1678912345,
		}
		jsonBytes, _ := json.Marshal(payload)
		respP := callApi(urlPost, "POST", jsonBytes)
		var dataP respPOST
		errP := json.Unmarshal([]byte(respP), &dataP)
		if errP != nil {
			fmt.Println(errP)
		}
		fmt.Println(dataP.Tx_hash)
	case "GET":
		Tx_hash := "5c34b30b5b3578f1f3524c852a3b9aed129f85c0cc689068165ba7d81b1cd642"
		urlGet := "https://mock-node-wgqbnxruha-as.a.run.app/check/" + Tx_hash
		respG := callApi(urlGet, "GET", nil)
		var dataG respGET
		errG := json.Unmarshal([]byte(respG), &dataG)
		if errG != nil {
			fmt.Println(errG)
		}
		fmt.Println(RepInfo(dataG.Tx_status))
	case "INFO":
		respData := "DNE"
		fmt.Println(RepInfo(respData))
	default:
		fmt.Println("Please Select")
	}
}
