package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func GetDriverNewsFromApi(apiURL string, apiToken string, name string) ([]News, error) {
	client := &http.Client{}
	newsReq, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		log.Println(err)
		return []News{}, errors.New("connection to service error")
	}
	query := newsReq.URL.Query()
	query.Add("api_token", apiToken)
	query.Add("search", name)
	query.Add("language", "en")
	query.Add("limit", "5")
	newsReq.URL.RawQuery = query.Encode()

	newsRes, err := client.Do(newsReq)
	if err != nil {
		log.Println(err)
		return []News{}, errors.New("connection to service error")

	}
	defer newsRes.Body.Close()

	var newsData map[string]interface{}
	if err := json.NewDecoder(newsRes.Body).Decode(&newsData); err != nil {
		return []News{}, errors.New("json decode error")
	}
	topNews := newsData["data"].([]interface{})

	var news []News
	for _, val := range topNews {
		if newsMap, ok := val.(map[string]interface{}); ok {
			title := newsMap["title"].(string)
			description := newsMap["description"].(string)
			url := newsMap["url"].(string)
			news = append(news, News{
				Title:       title,
				Description: description,
				Url:         url,
			})
		}

	}

	return news, nil
}

func GetDriverFastestLapsFromApi(apiKey string, apiURL string, id int) (int, error) {

	client := &http.Client{}
	poleReq, err := CreateRequest(
		"GET", apiURL, apiKey)
	if err != nil {
		log.Println(err)
		return -1, errors.New("connection to service error")
	}

	//add query
	query := poleReq.URL.Query()
	query.Add("type", "race")
	query.Add("season", "2022")
	poleReq.URL.RawQuery = query.Encode()

	poleRes, err := client.Do(poleReq)

	if err != nil {
		log.Println(err)
		return -1, errors.New("connection to service error")
	}
	defer poleRes.Body.Close()

	var poleData map[string]interface{}
	if err := json.NewDecoder(poleRes.Body).Decode(&poleData); err != nil {
		return -1, errors.New("json decode error")
	}

	flData := poleData["response"].([]interface{})

	fl := 0

	for _, val := range flData {
		if flm, ok := val.(map[string]interface{}); ok {
			//log.Println(flm)

			if fastlap, ok := flm["fastest_lap"].(map[string]interface{}); ok {

				if driver, ok := fastlap["driver"].(map[string]interface{}); ok {

					if driverID, ok := driver["id"].(float64); ok {

						if int(driverID) == id {
							fl++
						}
					}
				}
			}
		}
	}

	return fl, nil

}

func GetDriverFromApiByName(apiKey string, apiURL string, driverName string) (Driver, error) {
	client := &http.Client{}

	req, err := CreateRequest(
		"GET", apiURL, apiKey)
	if err != nil {
		log.Println(err)
		return Driver{}, errors.New("connection to service error")
	}

	//add query
	query := req.URL.Query()
	query.Add("name", strings.ToLower(driverName))
	req.URL.RawQuery = query.Encode()

	res, err := client.Do(req)

	if err != nil {
		log.Println(err)
		return Driver{}, errors.New("connection to service error")
	}
	defer res.Body.Close()

	var response DriverResponse
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		fmt.Println("Błąd dekodowania JSONa:", err)
		return Driver{}, errors.New("json decode error")
	}

	if len(response.Response) == 0 {
		return Driver{}, errors.New("not found")

	}
	return DriverFromResponseDriver(response), nil

}
