CREATE TABLE test_tasks_list_go.tasks (
  id INT NOT NULL AUTO_INCREMENT,
  title VARCHAR(100) NOT NULL,
  description VARCHAR(500) NOT NULL,
  status VARCHAR(50) NOT NULL,
  user_id INT NOT NULL,
  PRIMARY KEY (id),
  INDEX user_id_idx (user_id ASC) VISIBLE,
  CONSTRAINT user_id FOREIGN KEY (user_id) REFERENCES tasks_list_go.users (id) ON DELETE NO ACTION ON UPDATE NO ACTION
);