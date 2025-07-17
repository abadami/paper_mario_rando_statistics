package domain

import "time"

type RaceByPageResponse struct {
	Count    int    `json:"count"`
	NumPages int    `json:"num_pages"`
	Races    []Race `json:"races"`
}

type RaceStatus struct {
	Value        string `json:"value"`
	VerboseValue string `json:"verbose_value"`
	Helpstring   string `json:"help_string"`
}

type EntrantStatus struct {
	Value        string `json:"value"`
	VerboseValue string `json:"verbose_value"`
	Helpstring   string `json:"help_string"`
}

type Category struct {
	Name      string `json:"name"`
	ShortName string `json:"short_name"`
	Slug      string `json:"slug"`
	Url       string `json:"url"`
	DataUrl   string `json:"data_url"`
}

type Goal struct {
	Name   string `json:"name"`
	Custom bool   `json:"custom"`
}

type User struct {
	FullName string `json:"full_name"`
	Name     string `json:"name"`
}

type Entrant struct {
	User         User          `json:"user"`
	Status       EntrantStatus `json:"status"`
	FinishTime   string        `json:"finish_time"`
	Place        int           `json:"place"`
	PlaceOrdinal string        `json:"place_ordinal"`
}

type Race struct {
	Name                  string     `json:"name"`
	Url                   string     `json:"url"`
	Status                RaceStatus `json:"status"`
	DataUrl               string     `json:"data_url"`
	Goal                  Goal       `json:"goal"`
	Info                  string     `json:"info"`
	EntrantsCount         int        `json:"entrants_count"`
	EntrantsCountFinished int        `json:"entrants_count_finished"`
	EntrantsCountInactive int        `json:"entrants_count_inactive"`
	OpenedAt              string     `json:"opened_at"`
	StartedAt             string     `json:"started_at"`
	TimeLimit             string     `json:"time_limit"`
	Entrants              []Entrant  `json:"entrants"`
}

type RaceDetail struct {
	Name                  string     `json:"name"`
	Category              Category   `json:"category"`
	Status                RaceStatus `json:"status"`
	Url                   string     `json:"url"`
	DataUrl               string     `json:"data_url"`
	Goal                  Goal       `json:"goal"`
	Info                  string     `json:"info"`
	EntrantsCount         int        `json:"entrants_count"`
	EntrantsCountFinished int        `json:"entrants_count_finished"`
	EntrantsCountInactive int        `json:"entrants_count_inactive"`
	OpenedAt              string     `json:"opened_at"`
	StartedAt             string     `json:"started_at"`
	TimeLimit             string     `json:"time_limit"`
	Entrants              []Entrant  `json:"entrants"`
}

type StatisticRequest struct {
	ParticipantLimit int
	BeforeTime       string
	AfterTime        string
	ContainsEntrant  string
}

type StatisticsRequest struct {
	ParticipantLimit int
	BeforeTime       string
	AfterTime        string
	ContainsEntrant  int
	PageNumber       int
	Goal 						 string
}

type StatisticsResponse struct {
	Average    string `json:"average"`
	Deviation  string `json:"deviation"`
	RaceNumber int `json:"raceNumber"`
	DnfCount int `json:"dnfCount"`
	RawData []RaceEntrantAndRaceRecord `json:"rawData"`
}

type RaceEntrantAndRaceRecord struct {
	Id                  int       `json:"id"`
	Race_id             int       `json:"race_id"`
	Entrant_id          int       `json:"entrant_id"`
	Finish_time         string    `json:"finish_time"`
	Place               int       `json:"place"`
	Place_ordinal       string    `json:"place_ordinal"`
	Status              string    `json:"status"`
	Name                string    `json:"name"`
	Category_name       string    `json:"category_name"`
	Category_short_name string    `json:"category_short_name"`
	Url                 string    `json:"url"`
	Goal_name           string    `json:"goal_name"`
	Started_at          time.Time `json:"started_at"`
}

type RaceRecord struct {
	Id                  int `json:"id"`
	Name                string `json:"name"`
	Category_name       string `json:"category_name"`
	Category_short_name string `json:"category_short_name"`
	Url                 string `json:"url"`
	Goal_name           string `json:"goal_name"`
	Started_at          string `json:"started_at"`
}

type EntrantRecord struct {
	Id        int `json:"id"`
	Name      string `json:"name"`
	Full_name string `json:"full_name"`
}

type RaceEntrantRecord struct {
	Id            int `json:"id"`
	Race_id       int `json:"race_id"`
	Entrant_id    int `json:"entrant_id"`
	Finish_time   string `json:"finish_time"`
	Place         int `json:"place"`
	Place_ordinal string `json:"place_ordinal`
	Status        string `json:"status"`
}

type TaskLogRecord struct {
	Id            int `json:"id"`
	Date_ran      string `json:"date_ran"`
	Races_fetched int `json:"races_fetch"`
	Successful    bool `json:"successful"`
}
