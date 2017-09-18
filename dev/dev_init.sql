DELETE from tasks;
DELETE from places;

-- places 
INSERT INTO places(place_name, important) VALUES("university", 100);
INSERT INTO places(place_name, important) VALUES("friend", 50);

-- tasks 
INSERT INTO tasks(content, place_id, important) VALUES("test_task", 1, 100);
INSERT INTO tasks(content, place_id, important) VALUES("test_task2", 2, 0);
INSERT INTO tasks(content, place_id, important) VALUES("non_place_task", 0, 50);
