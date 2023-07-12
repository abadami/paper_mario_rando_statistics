use std::collections::HashMap;

mod api;
mod data;
mod duration;
mod utils;

use api::{get_fastest_time_for_race, get_race_titles_and_entrants_by_page_number};
use utils::convert_seconds_to_time;

//TODO: Cache results.
//TODO: API Conversion
//TODO: Maybe use a DB? Take a look at SurrealDB?
#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    let mut races = HashMap::<String, usize>::new();

    let mut previous_length = 0;
    let mut current_length = 1;
    let mut counter = 1;

    let client = reqwest::Client::new();

    while current_length > previous_length {
        previous_length = current_length;

        get_race_titles_and_entrants_by_page_number(&client, &mut races, counter).await?;

        println!("Current Iteration: {}", counter);

        counter += 1;
        current_length = races.len();
    }

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

    println!("Races Considered: {}", deviation_length);
    println!("Average Time: {}", average_time);
    println!("Standard Deviation: {}", standard_deviation);

    Ok(())
}
