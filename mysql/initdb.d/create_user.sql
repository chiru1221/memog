create table user(
    user_id int auto_increment,
    name varchar(20) unique,
    passwd binary(20),
    index(user_id)
);