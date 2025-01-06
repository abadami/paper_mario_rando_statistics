package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func getRaceTitlesAndEntrantsByPage(pageNum int) RaceByPageResponse {
	url := fmt.Sprintf("https://racetime.gg/pm64r/races/data?show_entrants=1&page=%d", pageNum)

	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	response := RaceByPageResponse{}
	unmarshalErr := json.Unmarshal(body, &response)
	if unmarshalErr != nil {
		return RaceByPageResponse{}
	}

	return response
}

func getRaceDetails(raceName string) RaceDetail {
	url := fmt.Sprintf("https://racetime.gg/%s/data", raceName)

	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	response := RaceDetail{}
	unmarshalErr := json.Unmarshal(body, &response)
	if unmarshalErr != nil {
		return RaceDetail{}
	}

	return response
}
