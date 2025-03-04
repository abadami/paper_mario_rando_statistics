package main

type RaceByPageResponse struct {
	Count    int    `json:"count"`
	NumPages int    `json:"num_pages"`
	Races    []Race `json:"races"`
}

type RaceStatus struct {
	Value        string `json:"value"`
	VerboseValue string `json:"verbose_value"`
	HelpText     string `json:"help_text"`
}

type EntrantStatus struct {
	Value        string `json:"value"`
	VerboseValue string `json:"verbose_value"`
	HelpText     string `json:"help_text"`
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
	ContainsEntrant  string
	PageNumber       int
}

type StatisticsResponse struct {
	Average    string
	Deviation  string
	RaceNumber int
}
