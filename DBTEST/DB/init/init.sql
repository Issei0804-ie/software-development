CREATE DATABASE information;
USE information;
CREATE TABLE users(id int primary key auto_increment, name varchar(255), age int);
INSERT INTO users(name,age) VALUES("hoge",20);
INSERT INTO users(name,age) VALUES("fuga",22);
INSERT INTO users(name,age) VALUES("piyo",15);

