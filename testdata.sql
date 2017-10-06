DELETE from tasks;
DELETE from places;

insert into tasks(content, limit_time, workload, place_id, importance) 
            values("content1", "", 0, 0, 0);
insert into tasks(content, limit_time, workload, place_id, importance) 
            values("content2", "", 0, 1, 0);
insert into tasks(content, limit_time, workload, place_id, importance) 
            values("content3", "", 10, 1, 50);

insert into places(place_name, importance) values("university", 10);
insert into places(place_name, importance) values("friends", 50);
