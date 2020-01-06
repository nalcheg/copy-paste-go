CREATE TABLE users
(
    id   SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE tickets
(
    id         SERIAL PRIMARY KEY,
    creator_id INTEGER REFERENCES users (id) NOT NULL,
    subject    TEXT                          NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE      NOT NULL
);

CREATE TABLE comments
(
    id         SERIAL PRIMARY KEY,
    ticket_id  INTEGER REFERENCES tickets (id) NOT NULL,
    user_id    INTEGER REFERENCES users (id)   NOT NULL,
    comment    TEXT                            NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE        NOT NULL
);

---- create above / drop below ----

DROP TABLE users;
DROP TABLE tickets;
DROP TABLE comments;
