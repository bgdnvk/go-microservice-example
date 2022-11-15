CREATE TABLE IF NOT EXISTS comments (
    id SERIAL PRIMARY KEY,
    comment VARCHAR NOT NULL,
    comment_date DATE DEFAULT CURRENT_DATE,
    user_id BIGINT REFERENCES users(id)
);

INSERT INTO users(name) VALUES('dev_test_user');
INSERT INTO comments (comment) VALUES('first test comment');