-- Drop tables if they already exist
DROP TABLE IF EXISTS posts;
DROP TABLE IF EXISTS BlogPosts;

-- Create a new table in the db
CREATE TABLE posts (
  id       SERIAL PRIMARY KEY,
  title    VARCHAR(200) NOT NULL,
  body     VARCHAR(200) NOT NULL
);

-- Seed the table 
INSERT INTO posts
  (title, body)
VALUES
  ('Blog Post 1', 'Hello Go!');