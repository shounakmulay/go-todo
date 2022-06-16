CREATE TABLE todos
(
    id          int PRIMARY KEY AUTO_INCREMENT,
    user_id     int,
    title       text,
    description text,
    due_date    timestamp,
    done        tinyint(1),
    created_at  timestamp,
    updated_at  timestamp,
    deleted_at  timestamp
);

CREATE INDEX todos_user_id_done_idx ON todos (user_id, done);

CREATE INDEX todos_due_date_idx ON todos (due_date);