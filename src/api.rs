use axum::http::StatusCode;
use chrono::{DateTime, ParseError, Utc};
use reqwest::Client;
use std::collections::HashMap;

use crate::data::{Race, RaceByPageResponse, RaceDetail};
use crate::duration::Duration;

pub struct FilterData {
    pub participant_limit: Option<u8>,
    pub before_time: Option<DateTime<Utc>>,
    pub after_time: Option<DateTime<Utc>>,
    pub contains_entrant: Option<String>,
    pub page_number: u32,
}

pub async fn get_race_titles_and_entrants_by_page_number(
    client: &Client,
    races: &mut HashMap<String, usize>,
    page_number: u32,
) -> Result<(), Box<dyn std::error::Error>> {
    let mut base_url = "https://racetime.gg/pm64r/races/data?show_entrants=1&page=".to_string();

    base_url.push_str(&page_number.to_string());

    let response: String = client.get(base_url).send().await?.text().await?;

    let races_response_data: RaceByPageResponse = serde_json::from_str(&response)?;

    for race in races_response_data.races {
        races.insert(race.name, 0);
    }

    Ok(())
}

pub async fn get_race_titles_by_filter(
    client: &Client,
    filter_data: FilterData,
    races: &mut HashMap<String, usize>,
) -> Result<(), Box<dyn std::error::Error>> {
    let mut base_url = "https://racetime.gg/pm64r/races/data?show_entrants=1&page=".to_string();

    base_url.push_str(&filter_data.page_number.to_string());

    let response: String = client.get(base_url).send().await?.text().await?;

    let races_response_data: RaceByPageResponse = serde_json::from_str(&response)?;

    let filtered_races = filter_races_by_filter_data(races_response_data, filter_data)?;

    for race in filtered_races {
        races.insert(race.name, 0);
    }

    Ok(())
}

pub async fn get_fastest_time_for_race(
    client: &Client,
    title: String,
    participate_limit: u8,
) -> Result<usize, StatusCode> {
    let mut base_url = "https://racetime.gg/".to_string();

    base_url.push_str(&title);
    base_url.push_str("/data");

    let request = client.get(base_url).send().await;

    let request_data = match request {
        Ok(val) => val,
        Err(_) => return Err(StatusCode::INTERNAL_SERVER_ERROR)
    };

    let response = request_data.text().await;

    let response_data = match response {
        Ok(val) => val,
        Err(_) => return Err(StatusCode::INTERNAL_SERVER_ERROR)
    };

    let json_request = serde_json::from_str(&response_data);

    let json_data: RaceDetail = match json_request {
        Ok(val) => val,
        Err(_) => return Err(StatusCode::INTERNAL_SERVER_ERROR)
    };

    if json_data.entrants_count < participate_limit.into() {
        println!("Less than limit");
        return Ok(0);
    }

    let first_entrant = json_data.entrants.first().unwrap();

    let duration_string = match &first_entrant.finish_time {
        Some(val) => val.to_string(),
        None => "PT0H0M0S".to_string(),
    };

    let duration = Duration::from_iso8601_str(&duration_string);

    Ok(duration.num_of_seconds() as usize)
}

fn filter_races_by_filter_data(
    races: RaceByPageResponse,
    filter_data: FilterData,
) -> Result<Vec<Race>, ParseError> {
    let filtered_races = races
        .races
        .into_iter()
        .filter(|race| filter_race_by_filter_data(race, &filter_data).unwrap_or(false))
        .collect();

    Ok(filtered_races)
}

fn filter_race_by_filter_data(race: &Race, filter_data: &FilterData) -> Result<bool, ParseError> {
    let participant_clause = match filter_data.participant_limit {
        Some(participate_limit) => is_race_within_participant_limit(race, participate_limit),
        None => true,
    };

    let before_time_clause = match filter_data.before_time {
        Some(before_time) => is_race_before_time(race, before_time)?,
        None => true,
    };

    let after_time_clause = match filter_data.after_time {
        Some(after_time) => is_race_after_time(race, after_time)?,
        None => true,
    };

    let entrant_clause = match &filter_data.contains_entrant {
        Some(entrant) => does_race_contain_entrant(race, entrant),
        None => true,
    };

    Ok(participant_clause && before_time_clause && after_time_clause && entrant_clause)
}

fn is_race_within_participant_limit(race: &Race, participate_limit: u8) -> bool {
    race.entrants_count >= participate_limit.into()
}

fn is_race_before_time(race: &Race, before_time: DateTime<Utc>) -> Result<bool, ParseError> {
    let parsed_time = race.started_at.parse::<DateTime<Utc>>()?;

    Ok(parsed_time < before_time)
}

fn is_race_after_time(race: &Race, after_time: DateTime<Utc>) -> Result<bool, ParseError> {
    let parsed_time = race.started_at.parse::<DateTime<Utc>>()?;

    Ok(parsed_time > after_time)
}

fn does_race_contain_entrant(race: &Race, entrant: &String) -> bool {
    race.entrants.iter().any(|val| val.user.name == *entrant)
}
