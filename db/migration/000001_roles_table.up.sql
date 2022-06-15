CREATE TABLE roles
(
    id           int PRIMARY KEY AUTO_INCREMENT,
    access_level int NOT NULL,
    name         text,
    created_at   timestamp,
    updated_at   timestamp,
    deleted_at   timestamp
);
