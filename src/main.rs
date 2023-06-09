use scraper::{Html, Selector};

//TODO: Figure out how to parse the top times from a detail page (use: https://racetime.gg/pm64r/kind-tastytonic-9752)
//TODO: Figure out how to parse page number
//TODO: Figure out how to traverse each page
//TODO: Figure out how to fetch each detail page
//TODO: Figure out STATISTICS
fn main() -> Result<(), Box<dyn std::error::Error>> {
    let response = reqwest::blocking::get("https://racetime.gg/pm64r")?.text()?;

    let document = Html::parse_document(&response);
    let selector = Selector::parse("span.slug").unwrap();

    let race_title_element = document.select(&selector).next().unwrap();

    let race_title = race_title_element.text().collect::<String>();

    println!("Race Title: {}", race_title);

    let mut race_details_url: String = "https://racetime.gg/pm64r/".to_string();

    race_details_url.push_str(&race_title);

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
