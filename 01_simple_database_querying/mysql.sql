-- INITIALIZE DATA --
drop database if exists i_view_stockbit;
create database i_view_stockbit;
use i_view_stockbit;

create table USER (
  ID serial primary key,
  UserName varchar (255) unique not null,
  Parent int not null
);

insert into USER (UserName, Parent)
values ('Ali', 2);

insert into USER (UserName, Parent)
values ('Budi', 0);

insert into USER (UserName, Parent)
values ('Cecep', 1);

-- SELECT QUERY --
select 
  u.ID,
  u.UserName,
  u2.UserName as "ParentUserName"
from USER u
left join USER u2
  on u.Parent = u2.ID
