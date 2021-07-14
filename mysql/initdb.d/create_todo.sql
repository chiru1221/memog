create table todo(
    id int auto_increment,
    user_id int,
    task varchar(100),
    date varchar(100),
    deadline varchar(100),
    index(id)
);