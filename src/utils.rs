use std::collections::HashMap;

use crate::data::StatisticResponse;

//TODO: Handle ISO Duration
pub fn convert_time_to_seconds(time: String) -> usize {
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

pub fn convert_seconds_to_time(seconds: usize) -> String {
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

fn get_average(races: &HashMap<String, usize>) -> usize {
    let total_seconds: usize = races.values().filter(|value| value > &&0).sum();

    let average = total_seconds / races.values().filter(|value| value > &&0).count();

    average
}

pub fn calculate_statistics(races: &HashMap<String, usize>) -> StatisticResponse {
    let mut deviations: Vec<usize> = Vec::new();

    if races.values().filter(|value| value > &&0).count() == 0 {
        return StatisticResponse { average: "00:00:00".to_string(), deviation: "00:00:00".to_string(), race_number: 0 }
    }

    let average = get_average(races);

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

    if deviation_length == 0 {
        return StatisticResponse { average: "00:00:00".to_string(), deviation: "00:00:00".to_string(), race_number: 0 }
    }

    let standard_deviation_in_seconds = deviation_sum / deviation_length;

    StatisticResponse { average: convert_seconds_to_time(average), deviation: convert_seconds_to_time(standard_deviation_in_seconds), race_number: deviation_length }
}