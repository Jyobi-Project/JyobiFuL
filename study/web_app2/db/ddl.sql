-- noinspection SqlNoDataSourceInspectionForFile

-- noinspection SqlDialectInspectionForFile

create database go_example;

use go_example;

create table users(
    id int auto_increment primary key,
    name varchar(255),
    age int,
    address varchar(255),
    update_at DATETIME DEFAULT CURRENT_TIMESTAMP
);