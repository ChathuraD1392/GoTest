-- +goose Up
-- +goose StatementBegin
CREATE TABLE posts(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    createdAt DATETIME NOT NULL 
);
-- +goose StatementEnd
INSERT INTO posts (title, content, createdAt) VALUES 
('The Rise of AI', 'An in-depth look at how AI is shaping the future.', '2024-12-30 10:00:00');

INSERT INTO posts (title, content, createdAt) VALUES 
('Top 10 JavaScript Libraries', 'Discover the most popular JS libraries of 2024.', '2024-12-30 11:00:00');

INSERT INTO posts (title, content, createdAt) VALUES 
('Understanding Databases', 'A beginner-friendly guide to relational databases.', '2024-12-30 12:00:00');

INSERT INTO posts (title, content, createdAt) VALUES 
('Healthy Living Tips', 'Practical advice for a healthier lifestyle.', '2024-12-30 13:00:00');

INSERT INTO posts (title, content, createdAt) VALUES 
('Exploring the Metaverse', 'What is the metaverse and why does it matter?', '2024-12-30 14:00:00');

INSERT INTO posts (title, content, createdAt) VALUES 
('Mastering Python', 'Learn the secrets of Python programming.', '2024-12-30 15:00:00');

INSERT INTO posts (title, content, createdAt) VALUES 
('Travel in 2024', 'Top destinations to visit this year.', '2024-12-30 16:00:00');

INSERT INTO posts (title, content, createdAt) VALUES 
('The Crypto Craze', 'Is cryptocurrency still worth investing in?', '2024-12-30 17:00:00');

INSERT INTO posts (title, content, createdAt) VALUES 
('Work-Life Balance', 'Strategies to achieve balance in your life.', '2024-12-30 18:00:00');

INSERT INTO posts (title, content, createdAt) VALUES 
('Cooking at Home', 'Easy recipes for delicious meals.', '2024-12-30 19:00:00');

-- +goose Down
-- +goose StatementBegin
DROP TABLE posts;
-- +goose StatementEnd
