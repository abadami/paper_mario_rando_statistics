use reqwest::Client;
use std::collections::HashMap;

use crate::data::{RaceByPageResponse, RaceDetail};
use crate::duration::Duration;

pub async fn get_race_titles_and_entrants_by_page_number(
    client: &Client,
    races: &mut HashMap<String, usize>,
    page_number: u32,
) -> Result<(), Box<dyn std::error::Error>> {
    let mut base_url = "https://racetime.gg/pm64r/races/data?page=".to_string();

    base_url.push_str(&page_number.to_string());

    let response: String = client.get(base_url).send().await?.text().await?;

    let races_response_data: RaceByPageResponse = serde_json::from_str(&response)?;

    for race in races_response_data.races {
        races.insert(race.name, 0);
    }

    Ok(())
}

pub async fn get_fastest_time_for_race(
    client: &Client,
    title: String,
    participate_limit: u8,
) -> Result<usize, Box<dyn std::error::Error>> {
    let mut base_url = "https://racetime.gg/".to_string();

    base_url.push_str(&title);
    base_url.push_str("/data");

    let response: String = client.get(base_url).send().await?.text().await?;

    let response_data: RaceDetail = serde_json::from_str(&response)?;

    if response_data.entrants_count < participate_limit.into() {
        println!("Less than limit");
        return Ok(0);
    }

    let first_entrant = response_data.entrants.first().unwrap();

    let duration_string = match &first_entrant.finish_time {
        Some(val) => val.to_string(),
        None => "PT0H0M0S".to_string(),
    };

    let duration = Duration::from_iso8601_str(&duration_string);

    Ok(duration.num_of_seconds() as usize)
}
