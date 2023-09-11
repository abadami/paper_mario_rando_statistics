use std::collections::HashMap;

mod api;
mod data;
mod duration;
mod utils;
mod route;

use api::{get_fastest_time_for_race, get_race_titles_and_entrants_by_page_number, FilterData};
use data::{StatisticResponse, StatisticRequest};
use reqwest::Client;
use utils::{convert_seconds_to_time, calculate_statistics};

type AliasedResult<T> = Result<T, Box<dyn std::error::Error>>;

async fn get_races(client: &Client, request: StatisticRequest) -> AliasedResult<HashMap<String, usize>> {
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

async fn get_race_times(client: &Client, race: String, times: u8) -> Vec<usize> {
    todo!()
}



async fn get_statistics(filter_data: StatisticRequest) -> AliasedResult<StatisticResponse> {
    let client = reqwest::Client::new();

    let mut races = get_races(&client, filter_data).await?;

    for (key, value) in races.iter_mut() {
        println!("Getting fastest time for race {}", key);

        let seconds = get_fastest_time_for_race(&client, key.to_string(), 0).await?;

        println!("Fastest time is {}", seconds);

        *value = seconds;
    }

    let response = calculate_statistics(&races);

    Ok(response)


}

//TODO: Cache results.
//TODO: API Conversion
//TODO: Maybe use a DB? Take a look at SurrealDB?
#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    let race_statistics = get_statistics(StatisticRequest { participant_limit: None, before_time: None, after_time: None, contains_entrant: None }).await?;

    println!("Races Considered: {}", race_statistics.race_number);
    println!("Average Time: {}", race_statistics.average);
    println!("Standard Deviation: {}", race_statistics.deviation);

    Ok(())
}
