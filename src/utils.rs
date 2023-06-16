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
