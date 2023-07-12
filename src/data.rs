use serde::{Deserialize, Serialize};

//Data definitions

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct RaceByPageResponse {
    pub count: u32,
    pub num_pages: u32,
    pub races: Vec<Race>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct RaceStatus {
    pub value: String,
    pub verbose_value: String,
    pub help_text: String,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct EntrantStatus {
    pub value: String,
    pub verbose_value: String,
    pub help_text: String,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct Category {
    pub name: String,
    pub short_name: String,
    pub slug: String,
    pub url: String,
    pub data_url: String,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct Goal {
    pub name: String,
    pub custom: bool,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct User {
    pub full_name: String,
    pub name: String,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct Entrant {
    pub user: User,
    pub status: EntrantStatus,
    pub finish_time: Option<String>,
    pub place: Option<u32>,
    pub place_ordinal: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct Race {
    pub name: String,
    pub status: RaceStatus,
    pub url: String,
    pub data_url: String,
    pub goal: Goal,
    pub info: String,
    pub entrants_count: u32,
    pub entrants_count_finished: u32,
    pub entrants_count_inactive: u32,
    pub opened_at: String,
    pub started_at: String,
    pub time_limit: String,
    pub entrants: Vec<Entrant>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct RaceDetail {
    pub name: String,
    pub category: Category,
    pub status: RaceStatus,
    pub url: String,
    pub data_url: String,
    pub goal: Goal,
    pub info: String,
    pub entrants_count: u32,
    pub entrants_count_finished: u32,
    pub entrants_count_inactive: u32,
    pub opened_at: String,
    pub started_at: String,
    pub time_limit: String,
    pub entrants: Vec<Entrant>,
}
