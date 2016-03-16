drop database airmeet;
create database airmeet;
use airmeet;

create table events(
  id int(10) auto_increment primary key,
  event_name varchar(64) not null,
  room_name varchar(64),
  description varchar(64),
  items varchar(64),
  major int(5),
  active boolean not null default true,
  created_at datetime,
  deleted_at datetime
)
