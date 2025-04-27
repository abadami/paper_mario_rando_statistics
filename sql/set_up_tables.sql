CREATE TABLE IF NOT EXISTS Races (
    id serial PRIMARY KEY,
    name text,
    category_name text,
    category_short_name text,
    url text,
    goal_name text,
    started_at date
);

CREATE TABLE IF NOT EXISTS Entrants (id serial PRIMARY KEY, name text, full_name text);

CREATE TABLE IF NOT EXISTS RaceEntrants (
    id serial PRIMARY KEY,
    race_id integer,
    entrant_id integer,
    finish_time text,
    place integer,
    place_ordinal text,
    status text,
    FOREIGN KEY (race_id) REFERENCES Races (id),
    FOREIGN KEY (entrant_id) REFERENCES Entrants (id)
);

CREATE TABLE IF NOT EXISTS TaskLog (
    id serial PRIMARY KEY,
    date_ran date,
    races_fetched integer,
    successful boolean
);
