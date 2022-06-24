CREATE TABLE roles (
    id int PRIMARY KEY AUTO_INCREMENT,
    access_level int NOT NULL,
    name varchar(25) UNIQUE,
    created_at timestamp,
    updated_at timestamp
);