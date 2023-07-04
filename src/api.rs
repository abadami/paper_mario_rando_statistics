use reqwest::Client;
use scraper::{Html, Selector};
use std::collections::HashMap;
use std::convert;

use crate::data::{Race, RaceByPageResponse, RaceDetail};

use crate::utils::convert_time_to_seconds;

pub async fn get_race_titles_for_page_number_by_scrapping(
    client: &Client,
    races: &mut HashMap<String, usize>,
    page_number: u32,
) -> Result<(), Box<dyn std::error::Error>> {
    let mut base_url = "https://racetime.gg/pm64r?page=".to_string();

    base_url.push_str(&page_number.to_string());

    let response = client.get(base_url).send().await?.text().await?;

    let document = Html::parse_document(&response);
    let selector = Selector::parse("span.slug")?;

    for element in document.select(&selector) {
        let race_title = element.text().collect::<String>();

        races.insert(race_title, 0);
    }

    Ok(())
}

pub async fn get_fatest_time_for_race_by_scraping(
    client: &Client,
    title: String,
    participate_limit: u8,
) -> Result<usize, Box<dyn std::error::Error>> {
    let mut base_url = "https://racetime.gg/pm64r/".to_string();

    base_url.push_str(&title);

    let response = client.get(base_url).send().await?.text().await?;

    let document = Html::parse_document(&response);
    let selector = Selector::parse("time.finish-time")?;
    let participant_selector = Selector::parse("li.entrant-row")?;

    let participants: u8 = document
        .select(&participant_selector)
        .count()
        .try_into()
        .unwrap();

    if participants < participate_limit {
        return Ok::<usize, Box<dyn std::error::Error>>(0);
    }

    let finish_time_node = document.select(&selector).next().unwrap();

    let finish_time = finish_time_node.text().collect::<String>();

    let finish_time_in_seconds = convert_time_to_seconds(finish_time);

    Ok(finish_time_in_seconds)
}

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

    match &first_entrant.finish_time {
        Some(val) => Ok(convert_time_to_seconds(val.to_string())),
        None => Ok(0),
    }
}
