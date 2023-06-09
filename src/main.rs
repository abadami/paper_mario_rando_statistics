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

//TODO: Figure out how to parse page number
//TODO: Figure out how to traverse each page
//TODO: Figure out how to fetch each detail page
//TODO: Figure out STATISTICS
fn main() -> Result<(), Box<dyn std::error::Error>> {
    let mut races = HashMap::<String, usize>::new();

    get_race_titles_for_page_number(&mut races, 1)?;

    let race_title = races.iter().next().unwrap().0;

    println!("Race Title: {}", race_title);

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

    /*
    for element in document.select(&selector) {
        let race_title = element.text().collect::<String>();



        println!("{}", element.text().collect::<String>());
    }
    */

    Ok(())
}
