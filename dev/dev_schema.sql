DROP TABLE IF EXISTS tasks;
DROP TABLE IF EXISTS places;

CREATE TABLE tasks (
    task_id INTEGER PRIMARY KEY,
    content TEXT NOT NULL,
    place_id INTEGER DEFAULT 0, 
    important INTEGER,
    done BOOLEAN DEFAULT 0
);

CREATE TABLE places (
    place_id INTEGER PRIMARY KEY,
    place_name TEXT NOT NULL,
    important INTEGER
); 
