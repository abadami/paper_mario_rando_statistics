use scraper::{Html, Selector};
use std::collections::HashMap;

fn get_race_titles_for_page_number(
    races: &mut HashMap<String, usize>,
    page_number: u32,
) -> Result<(), Box<dyn std::error::Error>> {
    let mut base_url = "https://racetime.gg/pm64r?page=".to_string();

    base_url.push_str(&page_number.to_string());

    let response = reqwest::blocking::get(base_url)?.text()?;

    let document = Html::parse_document(&response);
    let selector = Selector::parse("span.slug")?;

    for element in document.select(&selector) {
        let race_title = element.text().collect::<String>();

        races.insert(race_title, 0);
    }

    Ok(())
}

fn get_fatest_time_for_race(title: String) -> Result<usize, Box<dyn std::error::Error>> {
    let mut base_url = "https://racetime.gg/pm64r/".to_string();

    base_url.push_str(&title);

    let response = reqwest::blocking::get(base_url)?.text()?;

    let document = Html::parse_document(&response);
    let selector = Selector::parse("time.finish-time")?;

    let finish_time_node = document.select(&selector).next().unwrap();

    let finish_time = finish_time_node.text().collect::<String>();

    Ok(0)
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

//TODO: Figure out how to traverse each page
//TODO: Figure out how to fetch each detail page
//TODO: Figure out STATISTICS
fn main() -> Result<(), Box<dyn std::error::Error>> {
    let mut races = HashMap::<String, usize>::new();

    let mut previous_length = 0;
    let mut current_length = 1;
    let mut counter = 1;

    while current_length > previous_length {
        previous_length = current_length;

        get_race_titles_for_page_number(&mut races, counter)?;

        counter += 1;
        current_length = races.len();
    }

    let race_title = races.iter().next().unwrap().0;

    println!("Race Title: {}", race_title);
    println!("Races Length: {}", races.len());

    let mut race_details_url: String = "https://racetime.gg/pm64r/".to_string();

    race_details_url.push_str(race_title);

    let details_response = reqwest::blocking::get(race_details_url)?.text()?;

    let details_document = Html::parse_document(&details_response);
    let finish_time_selector = Selector::parse("time.finish-time").unwrap();

    let finish_time_node = details_document
        .select(&finish_time_selector)
        .next()
        .unwrap();

    let finish_time = finish_time_node.text().collect::<String>();

    println!("Finish Time: {}", finish_time);

    Ok(())
}
