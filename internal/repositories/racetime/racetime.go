package racetime

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/abadami/randomizer-statistics/domain"
)

type RacetimeRepository struct {

}

func NewRacetimeRepository() *RacetimeRepository {
	return &RacetimeRepository{}
}

func (*RacetimeRepository) GetRaceTitlesAndEntrantsByPage(pageNum int) domain.RaceByPageResponse {
	url := fmt.Sprintf("https://racetime.gg/pm64r/races/data?show_entrants=1&page=%d&per_page=100", pageNum)

	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	response := domain.RaceByPageResponse{}
	unmarshalErr := json.Unmarshal(body, &response)
	if unmarshalErr != nil {
		return domain.RaceByPageResponse{}
	}

	return response
}

func (*RacetimeRepository) GetRaceDetails(raceName string) domain.RaceDetail {
	url := fmt.Sprintf("https://racetime.gg/%s/data", raceName)

	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	response := domain.RaceDetail{}
	unmarshalErr := json.Unmarshal(body, &response)
	if unmarshalErr != nil {
		return domain.RaceDetail{}
	}

	return response
}