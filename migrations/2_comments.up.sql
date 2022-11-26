CREATE TABLE IF NOT EXISTS comments (
    id SERIAL PRIMARY KEY,
    comment VARCHAR NOT NULL,
    user_id BIGINT REFERENCES users(id)
);

INSERT INTO users(name) VALUES('dev_test_user');
INSERT INTO comments (comment, user_id) VALUES('first test comment 1', 1);