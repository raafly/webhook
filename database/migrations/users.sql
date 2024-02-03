-- DROP TABLE if EXISTS 
CREATE TABLE users(
    id serial,
    UUID varchar(50) not null primary key,
    email varchar(50) not null unique,
    name varchar(50) not null,
    password varchar(50) not null,
    token varchar(50),
    email_verified_at timestamp default current_timestamp,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp
)

-- token for access all api