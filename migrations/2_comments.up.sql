CREATE TABLE IF NOT EXISTS comments (
    id SERIAL PRIMARY KEY,
    comment VARCHAR NOT NULL,
    comment_date DATE DEFAULT CURRENT_DATE,
    user_id BIGINT REFERENCES users(id)
);

INSERT INTO users VALUES(1, 'dev_test_user');
INSERT INTO comments (comment, user_id) 
VALUES("first test comment", "1");