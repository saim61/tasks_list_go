CREATE SCHEMA IF NOT EXISTS tasks_list_go;

CREATE TABLE IF NOT EXISTS tasks_list_go.tasks (
    id INT NOT NULL AUTO_INCREMENT,
    title VARCHAR(100) NOT NULL,
    description VARCHAR(500) NOT NULL,
    status VARCHAR(50) NOT NULL,
    PRIMARY KEY (id),
    UNIQUE INDEX id_unique (id ASC) VISIBLE
);

INSERT INTO tasks_list_go.tasks (id, title, description, status) VALUES ('1', 'Task 1', 'This is a testing task', 'open');
INSERT INTO tasks_list_go.tasks (id, title, description, status) VALUES ('2', 'Kitchen', 'Wash dishes', 'open');
INSERT INTO tasks_list_go.tasks (id, title, description, status) VALUES ('3', 'Laundry', 'Wash and fold your clothes', 'open');


