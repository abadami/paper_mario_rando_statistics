use std::collections::HashMap;

mod api;
mod data;
mod duration;
mod utils;
mod route;

use api::get_race_titles_and_entrants_by_page_number;
use axum::{Router, routing::get};
use data::StatisticRequest;
use reqwest::Client;
use route::get_statistics_with_filters;

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

// async fn get_statistics(filter_data: StatisticRequest) -> AliasedResult<StatisticResponse> {
//     let client = reqwest::Client::new();

//     let mut races = get_races(&client, filter_data).await?;

//     for (key, value) in races.iter_mut() {
//         println!("Getting fastest time for race {}", key);

//         let seconds = get_fastest_time_for_race(&client, key.to_string(), 0).await;

//         println!("Fastest time is {}", seconds);

//         *value = seconds;
//     }

//     let response = calculate_statistics(&races);

//     Ok(response)


// }

//TODO: Cache results.
//TODO: API Conversion
//TODO: Maybe use a DB? Take a look at SurrealDB?
#[tokio::main]
async fn main() {
    //build application
    let app = Router::new().route("/", get(get_statistics_with_filters));

    let listener = tokio::net::TcpListener::bind("0.0.0.0:3000").await.unwrap();
    axum::serve(listener, app).await.unwrap()

}
