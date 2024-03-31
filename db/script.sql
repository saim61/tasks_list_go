CREATE SCHEMA IF NOT EXISTS tasks_list_go;

CREATE TABLE IF NOT EXISTS tasks_list_go.tasks (
    id INT NOT NULL AUTO_INCREMENT,
    title VARCHAR(100) NOT NULL,
    description VARCHAR(500) NOT NULL,
    status VARCHAR(50) NOT NULL,
    PRIMARY KEY (id),
    UNIQUE INDEX id_unique (id ASC) VISIBLE
);

INSERT INTO tasks_list_go.tasks (id, title, description, status) 
VALUES 
('1', 'Task 1', 'This is a testing task', 'open'), 
('2', 'Kitchen', 'Wash dishes', 'open'), 
('3', 'Laundry', 'Wash and fold your clothes', 'open'), 
('4', 'Work', 'Write some code and rock the world!', 'open'), 
('5', 'Chill', 'Play some Elden Ring!', 'open');