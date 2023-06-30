//Data definitions

pub enum RaceStatusValue {
    Open,
    Invitational,
    Pending,
    InProgress,
    Finished,
    Cancelled,
}

pub enum EntrantStatusValue {
    Requested,
    Invited,
    Declined,
    Ready,
    NotReady,
    InProgress,
    Done,
    Dnf,
    Dq,
}

pub struct RaceStatus {
    pub value: RaceStatusValue,
    pub verbose_value: String,
    pub help_text: String,
}

pub struct EntrantStatus {
    pub value: EntrantStatusValue,
    pub verbose_value: String,
    pub help_text: String,
}

pub struct Category {
    pub name: String,
    pub short_name: String,
    pub slug: String,
    pub url: String,
    pub data_url: String,
}

pub struct Goal {
    pub name: String,
    pub custom: bool,
}

pub struct User {
    pub full_name: String,
    pub name: String,
}

pub struct Entrant {
    pub user: User,
    pub status: EntrantStatus,
    pub finish_time: Option<String>,
    pub place: u32,
    pub place_ordinal: String,
}

pub struct Race {
    pub name: String,
    pub category: Category,
    pub status: RaceStatus,
    pub url: String,
    pub data_url: String,
    pub goal: Goal,
    pub info: String,
    pub entrants_count: u32,
    pub entrants_count_finished: u32,
    pub entrants_count_inactive: u32,
    pub opened_at: String,
    pub started_at: String,
    pub time_limit: String,
}

pub struct RaceDetail {
    pub name: String,
    pub category: Category,
    pub status: RaceStatus,
    pub url: String,
    pub data_url: String,
    pub goal: Goal,
    pub info: String,
    pub entrants_count: u32,
    pub entrants_count_finished: u32,
    pub entrants_count_inactive: u32,
    pub opened_at: String,
    pub started_at: String,
    pub time_limit: String,
    pub entrants: Vec<Entrant>,
}
