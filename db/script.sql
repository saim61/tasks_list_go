CREATE SCHEMA IF NOT EXISTS tasks_list_go;

CREATE TABLE tasks_list_go.users (
  id INT NOT NULL AUTO_INCREMENT,
  email VARCHAR(100) NOT NULL,
  password VARCHAR(100) NOT NULL,
  PRIMARY KEY (id),
  UNIQUE INDEX id_UNIQUE (id ASC) VISIBLE,
  UNIQUE INDEX email_UNIQUE (email ASC) VISIBLE,
  UNIQUE INDEX password_UNIQUE (password ASC) VISIBLE
);

CREATE TABLE tasks_list_go.tasks (
  id INT NOT NULL AUTO_INCREMENT,
  title VARCHAR(100) NOT NULL,
  description VARCHAR(500) NOT NULL,
  status VARCHAR(50) NOT NULL,
  user_id INT NOT NULL,
  PRIMARY KEY (id),
  INDEX user_id_idx (user_id ASC) VISIBLE,
  CONSTRAINT user_id FOREIGN KEY (user_id) REFERENCES tasks_list_go.users (id) ON DELETE NO ACTION ON UPDATE NO ACTION
);

INSERT INTO tasks_list_go.tasks (id, title, description, status)
VALUES ('1', 'Task 1', 'This is a testing task', 'open'),
  ('2', 'Kitchen', 'Wash dishes', 'open'),
  (
    '3',
    'Laundry',
    'Wash and fold your clothes',
    'open'
  ),
  (
    '4',
    'Work',
    'Write some code and rock the world!',
    'open'
  ),
  ('5', 'Chill', 'Play some Elden Ring!', 'open');