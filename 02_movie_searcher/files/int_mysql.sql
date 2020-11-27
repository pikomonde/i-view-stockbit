-- INITIALIZE DATA --
drop database if exists i_view_stockbit_omdb;
create database i_view_stockbit_omdb;
use i_view_stockbit_omdb;

create table movie_log_query (
  id serial primary key,
  search_key varchar (255) not null,
  page_key int not null,
  title varchar(255) not null,
  year varchar(255) not null,
  imdb_id varchar(255) not null,
  type varchar(255) not null,
  poster_url varchar(255) not null,
  create_time timestamp not null
);
