USE test;

CREATE TABLE users (
  user_id INT PRIMARY KEY AUTO_INCREMENT,
  username VARCHAR(50) NOT NULL,
  email VARCHAR(100) NOT NULL,
  password VARCHAR(255) NOT NULL
);

CREATE TABLE posts (
  post_id INT PRIMARY KEY AUTO_INCREMENT,
  user_id INT NOT NULL,
  title VARCHAR(255) NOT NULL,
  content TEXT NOT NULL,
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (user_id) REFERENCES users(user_id)
);

CREATE TABLE comments (
  comment_id INT PRIMARY KEY AUTO_INCREMENT,
  user_id INT NOT NULL,
  post_id INT NOT NULL,
  content TEXT NOT NULL,
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (user_id) REFERENCES users(user_id),
  FOREIGN KEY (post_id) REFERENCES posts(post_id)
);

CREATE TABLE categories (
  category_id INT PRIMARY KEY AUTO_INCREMENT,
  name VARCHAR(50) NOT NULL
);

CREATE TABLE post_categories (
  post_id INT NOT NULL,
  category_id INT NOT NULL,
  PRIMARY KEY (post_id, category_id),
  FOREIGN KEY (post_id) REFERENCES posts(post_id),
  FOREIGN KEY (category_id) REFERENCES categories(category_id)
);


INSERT INTO users (username, email, password) VALUES
  ('john.doe', 'john.doe@example.com', 'password123'),
  ('jane.doe', 'jane.doe@example.com', 'password456'),
  ('user1', 'user1@example.com', 'password1'),
  ('user2', 'user2@example.com', 'password2'),
  ('user3', 'user3@example.com', 'password3'),
  ('user4', 'user4@example.com', 'password4'),
  ('user5', 'user5@example.com', 'password5'),
  ('user6', 'user6@example.com', 'password6'),
  ('user7', 'user7@example.com', 'password7'),
  ('user8', 'user8@example.com', 'password8'),
  ('user9', 'user9@example.com', 'password9'),
  ('user10', 'user10@example.com', 'password10'),
  ('bob.smith', 'bob.smith@example.com', 'password789');

INSERT INTO posts (user_id, title, content) VALUES
  (1, 'My First Post', 'This is my first blog post.'),
  (2, 'My Second Post', 'This is my second blog post.'),
  (1, 'My Third Post', 'This is my third blog post.'),
  (1, 'My First Post', 'This is my first blog post.'),
  (2, 'My Second Post', 'This is my second blog post.'),
  (1, 'My Third Post', 'This is my third blog post.'),
  (2, 'My Fourth Post', 'This is my fourth blog post.'),
  (1, 'My Fifth Post', 'This is my fifth blog post.'),
  (2, 'My Sixth Post', 'This is my sixth blog post.'),
  (1, 'My Seventh Post', 'This is my seventh blog post.'),
  (2, 'My Eighth Post', 'This is my eighth blog post.'),
  (1, 'My Ninth Post', 'This is my ninth blog post.'),
  (2, 'My Tenth Post', 'This is my tenth blog post.');

INSERT INTO comments (user_id, post_id, content) VALUES
  (1, 1, 'Great post!'),
  (2, 1, 'I totally agree.'),
  (3, 2, 'Thanks for sharing.');

INSERT INTO categories (name) VALUES
  ('Programming'),
  ('Travel'),
  ('Food');

INSERT INTO post_categories (post_id, category_id) VALUES
  (1, 1),
  (2, 2),
  (3, 3),
  (1, 3);
