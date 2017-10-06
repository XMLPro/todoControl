CREATE TABLE tasks (
  id         INTEGER PRIMARY KEY,
  content    TEXT NOT NULL,
  limit_time TEXT DEFAULT "",
  workload   INTEGER,
  place_id   INTEGER DEFAULT 0,
  importance INTEGER
);

CREATE TABLE places (
  id         INTEGER PRIMARY KEY,
  place_name TEXT    NOT NULL,
  importance INTEGER NOT NULL
);
