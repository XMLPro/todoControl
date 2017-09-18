CREATE TABLE tasks (
	task_id INTEGER PRIMARY KEY,
	content TEXT,
	place_id INTEGER, -- いい感じのが思いつかなかったのでなかったらフロント側で処理する
	important INTEGER
);

CREATE TABLE places (
	place_id INTEGER PRIMARY KEY,
	place_name TEXT,
	important INTEGER
);
