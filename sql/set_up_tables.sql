CREATE TABLE IF NOT EXISTS Races (
    id integer PRIMARY KEY,
    name text,
    category_name text,
    category_short_name text,
    url text,
    goal_name text,
    started_at date,
);

CREATE TABLE IF NOT EXISTS Entrants (
    id integer PRIMARY KEY,
    name text,
    full_name text
);

CREATE TABLE IF NOT EXISTS RaceEntrants (
    id integer PRIMARY KEY,
    race_id integer,
    entrant_id integer,
    finish_time text,
    place integer,
    place_ordinal text,
    status text,
    FOREIGN KEY (race_id) REFERENCES Races(id),
    FOREIGN KEY (entrant_id) REFERENCES Entrants(id)
);

CREATE TABLE IF NOT EXISTS TaskLog (
    id integer PRIMARY KEY,
    date_ran date,
    races_fetched integer,
    successful boolean
);
