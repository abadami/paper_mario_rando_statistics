use std::collections::HashMap;

mod api;
mod data;
mod duration;
mod utils;
mod route;

use api::{get_fastest_time_for_race, get_race_titles_and_entrants_by_page_number, FilterData};
use data::StatisticResponse;
use reqwest::Client;
use utils::convert_seconds_to_time;

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

async fn get_race_times(client: &Client, race: String, times: u8) -> Vec<usize> {
    todo!()
}

async fn get_average() -> usize {
    0
}

async fn get_deviation() -> usize {
    0
}

async fn get_statistics(filter_data: FilterData) -> AliasedResult<StatisticResponse> {
    let client = reqwest::Client::new();

    let mut races = get_races(&client).await?;

    for (key, value) in races.iter_mut() {
        println!("Getting fastest time for race {}", key);

        let seconds = get_fastest_time_for_race(&client, key.to_string(), 0).await?;

        println!("Fastest time is {}", seconds);

        *value = seconds;
    }

    let total_seconds: usize = races.values().filter(|value| value > &&0).sum();

    let average = total_seconds / races.values().filter(|value| value > &&0).count();

    let mut deviations: Vec<usize> = Vec::new();

    for value in races.values() {
        if value == &0 {
            continue;
        }

        let deviation = if average > *value {
            (average - value) ^ 2
        } else {
            (value - average) ^ 2
        };

        deviations.push(deviation);
    }

    let deviation_sum: usize = deviations.iter().sum();
    let deviation_length = deviations.len();

    let standard_deviation_in_seconds = deviation_sum / deviation_length;

    let average_time = convert_seconds_to_time(average);
    let standard_deviation = convert_seconds_to_time(standard_deviation_in_seconds);

    Ok(StatisticResponse {
        average: average_time,
        deviation: standard_deviation,
        race_number: deviation_length
    })


}

//TODO: Cache results.
//TODO: API Conversion
//TODO: Maybe use a DB? Take a look at SurrealDB?
#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    let race_statistics = get_statistics(FilterData { participant_limit: None, before_time: None, after_time: None, contains_entrant: None, page_number: 1 }).await?;

    println!("Races Considered: {}", race_statistics.race_number);
    println!("Average Time: {}", race_statistics.average);
    println!("Standard Deviation: {}", race_statistics.deviation);

    Ok(())
}
