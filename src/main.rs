use std::collections::HashMap;

mod api;
mod utils;

use api::{get_fatest_time_for_race, get_race_titles_for_page_number};
use utils::convert_seconds_to_time;

//TODO: Modulize it a little bit
//TODO: Cache results.
//TODO: Results will need to be fetched either periodically or on request. Ideally, we could do this periodically, but on request will probably be easiest. I think the web client will just fetch all data at start and use it.
#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    let mut races = HashMap::<String, usize>::new();

    let mut previous_length = 0;
    let mut current_length = 1;
    let mut counter = 1;

    let client = reqwest::Client::new();

    while current_length > previous_length {
        previous_length = current_length;

        get_race_titles_for_page_number(&client, &mut races, counter).await?;

        println!("Current Iteration: {}", counter);

        counter += 1;
        current_length = races.len();
    }

    for (key, value) in races.iter_mut() {
        let seconds = get_fatest_time_for_race(&client, key.to_string(), 0).await?;

        println!("Getting fastest time for race {}", key);

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
