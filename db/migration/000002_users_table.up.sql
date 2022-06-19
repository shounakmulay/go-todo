CREATE TABLE users
(
    id         int PRIMARY KEY AUTO_INCREMENT,
    first_name text,
    last_name  text,
    username   VARCHAR(255) UNIQUE,
    password   text,
    email      VARCHAR(255) UNIQUE,
    mobile     text,
    token      text,
    role_id    int,
    FOREIGN KEY (role_id) REFERENCES roles (id),
    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp
);

CREATE INDEX users_role_id_idx ON users (role_id);