CREATE TABLE users
(
    id         int PRIMARY KEY AUTO_INCREMENT,
    first_name text,
    last_name  text,
    username   VARCHAR(255) UNIQUE,
    password   text,
    email      VARCHAR(255) UNIQUE,
    mobile     text,
    role_id    int REFERENCES roles (id),
    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp
);

CREATE INDEX users_username_idx ON users (username);

CREATE INDEX users_email_idx ON users (email);

CREATE INDEX users_role_id_idx ON users (role_id);