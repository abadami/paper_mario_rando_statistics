use reqwest::Client;
use scraper::{Html, Selector};
use std::collections::HashMap;

async fn get_race_titles_for_page_number(
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

async fn get_fatest_time_for_race(
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

fn convert_time_to_seconds(time: String) -> usize {
    let mut total_time = 0;

    let time_split = time.split(':');

    let mut multiplier = 3600;

    for time in time_split {
        let parsed_time = time.parse::<usize>();

        let converted_seconds = match parsed_time {
            Ok(val) => val * multiplier,
            _ => 0,
        };

        total_time += converted_seconds;
        multiplier /= 60;
    }

    total_time
}

fn convert_seconds_to_time(seconds: usize) -> String {
    let mut time: String = String::new();

    let mut seconds_tracker = seconds;

    let hours = seconds_tracker / 3600;

    seconds_tracker -= hours * 3600;

    let minutes = seconds_tracker / 60;

    seconds_tracker -= minutes * 60;

    time.push_str(&hours.to_string());
    time.push(':');
    time.push_str(&minutes.to_string());
    time.push(':');
    time.push_str(&seconds_tracker.to_string());

    time
}

//TODO: Filter races based on number of participants
//TODO: Convert to API?
//TODO: Multithreads?
//TODO: API fetches "latest" result, with results calculated every {time period}
//TODO: Modulize it a little bit
//TODO: performance
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
