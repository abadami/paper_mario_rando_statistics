use std::collections::HashMap;

use axum::{extract::Query, http::StatusCode, extract::Json};
use axum_macros::debug_handler;
use reqwest::Client;

use crate::{api::{get_fastest_time_for_race, get_race_titles_and_entrants_by_page_number, FilterData}, data::{StatisticResponse, StatisticRequest, StatisticsRequest}, utils::calculate_statistics};

type AliasedResult<T> = Result<T, Box<dyn std::error::Error>>;

async fn get_races(client: &Client) -> AliasedResult<HashMap<String, usize>> {
  let mut races: HashMap<String, usize> = HashMap::<String, usize>::new();

  let mut previous_length = 0;
  let mut current_length = 1;
  let mut counter = 1;

  while current_length > previous_length {
      previous_length = current_length;

      get_race_titles_and_entrants_by_page_number(client, &mut races, counter).await?;

      println!("Current Iteration: {}", counter);

      counter += 1;
      current_length = races.len();
  }

  Ok(races)
}

#[debug_handler]
pub async fn get_statistics_with_filters (Query(params): Query<HashMap<String, String>>) -> Result<Json<StatisticResponse>, StatusCode> {
  let client = reqwest::Client::new();

  let racesResult = get_races(&client).await;

  let mut races = match racesResult {
    Ok(val) => val,
    Err(_) => return Err(StatusCode::INTERNAL_SERVER_ERROR)
  };

  //for (key, value) in races.iter_mut() {
  //    println!("Getting fastest time for race {}", key);
  //    let fastest_time_result = get_fastest_time_for_race(&client, key.to_string(), 0).await;
  //    let seconds = match fastest_time_result {
  //      Ok(val) => val,
  //      Err(_) => return Err(StatusCode::INTERNAL_SERVER_ERROR)
  //    };
  //    println!("Fastest time is {}", seconds);
  //    *value = seconds;
  //}

  let statistics = calculate_statistics(&races);

  Ok(Json(statistics))
}
