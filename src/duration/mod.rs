pub struct Duration {
    pub hours: u64,
    pub minutes: u64,
    pub seconds: u64,
}

impl Duration {
    pub fn from_iso8601_str(iso_str: &str) -> Self {
        let mut iso_string = iso_str.to_string();

        iso_string = iso_string.replace("P0DT", "");
        iso_string = iso_string.replace('H', ".");
        iso_string = iso_string.replace('M', ".");

        let iso_split: Vec<&str> = iso_string.split('.').collect();

        let hours = iso_split[0].parse().unwrap_or(0);

        let minutes = iso_split[1].parse().unwrap_or(0);

        let seconds = iso_split[2].parse().unwrap_or(0);

        Duration {
            hours,
            minutes,
            seconds,
        }
    }

    pub fn from_seconds(seconds: u64) -> Self {
        let mut seconds_tracker = seconds;

        let hours = seconds_tracker / 3600;

        seconds_tracker -= hours * 3600;

        let minutes = seconds_tracker / 60;

        seconds_tracker -= minutes * 60;

        Duration {
            hours,
            minutes,
            seconds: seconds_tracker,
        }
    }

    pub fn num_of_seconds(self) -> u64 {
        (self.hours * 3600) + (self.minutes * 60) + self.seconds
    }
}

impl std::fmt::Display for Duration {
    fn fmt(&self, f: &mut std::fmt::Formatter) -> std::fmt::Result {
        write!(f, "{}:{}:{}", self.hours, self.minutes, self.seconds)
    }
}
