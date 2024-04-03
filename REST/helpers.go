package main

import (
	"fmt"
	"net/http"
	"slices"
)

func DriverFromResponseDriver(responseDriver DriverResponse) Driver {
	wins := 0
	if responseDriver.Response[0].HighestRaceFinish.Position == 1 {
		wins = responseDriver.Response[0].HighestRaceFinish.Number

	}
	var teams []string
	for _, team := range responseDriver.Response[0].Teams {
		if slices.Contains(teams, team.Team.Name) {
			continue
		}
		teams = append(teams, team.Team.Name)

	}
	winsPercentage := 0.0
	podiumsPercentage := 0.0

	if responseDriver.Response[0].GrandsPrixEntered > 0 {
		winsPercentage = float64(wins) / float64(responseDriver.Response[0].GrandsPrixEntered) * 100
		podiumsPercentage = float64(responseDriver.Response[0].Podiums) / float64(responseDriver.Response[0].GrandsPrixEntered) * 100
	}
	winsPercentageStr := fmt.Sprintf("%.2f%%", winsPercentage)
	podiumsPercentageStr := fmt.Sprintf("%.2f%%", podiumsPercentage)

	return Driver{
		Id:                float64(responseDriver.Response[0].ID),
		Name:              responseDriver.Response[0].Name,
		Country:           responseDriver.Response[0].Nationality,
		Teams:             teams,
		Wins:              wins,
		Races:             responseDriver.Response[0].GrandsPrixEntered,
		Championships:     responseDriver.Response[0].WorldChampionships,
		Points:            responseDriver.Response[0].CareerPoints,
		Podiums:           responseDriver.Response[0].Podiums,
		Number:            responseDriver.Response[0].Number,
		WinsPercentage:    winsPercentageStr,
		PodiumsPercentage: podiumsPercentageStr,
	}
}

func CreateRequest(method string, url string, apiKey string) (*http.Request, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("x-rapidapi-key", apiKey)
	req.Header.Add("x-rapidapi-host", "v1.formula-1.api-sports.io")

	return req, nil
}
