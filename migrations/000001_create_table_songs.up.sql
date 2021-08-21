-- Set timezone
SET TIMEZONE = "Asia/Bangkok";

--Create table songs
CREATE TABLE songs
(
    id          VARCHAR(21) PRIMARY KEY,
    title       TEXT    NOT NULL,
    year        INTEGER NOT NULL,
    performer   TEXT    NOT NULL,
    genre       TEXT,
    duration    TEXT,
    inserted_at TEXT    NOT NULL,
    updated_at  TEXT    NOT NULL,
);