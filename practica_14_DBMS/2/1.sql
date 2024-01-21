/* 1 */
CREATE TABLE project(
  id VARCHAR(255) PRIMARY KEY,
  name VARCHAR(255)
);

CREATE TABLE department(
  id VARCHAR(255) PRIMARY KEY,
  name VARCHAR(255)
);

CREATE TABLE employee(
  id SERIAL PRIMARY KEY,
  name VARCHAR(255),
  address VARCHAR(255),
  department_id VARCHAR(255) REFERENCES department(id)
);

CREATE TABLE assignment(
  id SERIAL PRIMARY KEY,
  employee_id INT REFERENCES employee(id),
  project_id VARCHAR(255) REFERENCES project(id)
);

INSERT INTO 
  project(name, id)
VALUES
  ('Desarrollo', 'p1'),
  ('Sistema de Gestión', 'p2'),
  ('Capacitación', 'p3');

INSERT INTO 
  department(name, id)
VALUES
  ('Producción', 'd1'),
  ('Ventas', 'd2'),
  ('RRHH', 'd3'),
  ('Contabilidad', 'd4');


INSERT INTO 
  employee(name, address, department_id)
VALUES
  ('Carlos', 'Direcc1', 'd1'),
  ('Pepe', 'Direcc2', 'd2'),
  ('Susana', 'Direcc3', 'd3'),
  ('Graciela', 'Direcc4', 'd4');

INSERT INTO 
  assignment(employee_id, project_id)
VALUES
  (1, 'p3'),
  (2, 'p1'),
  (3, 'p2');

/* 2 */
CREATE TABLE hobbie(
  id VARCHAR(255) PRIMARY KEY,
  name VARCHAR(255),
  employee_id INT REFERENCES employee(id)
);

/* 3 */
INSERT INTO 
  assignment(employee_id, project_id)
VALUES
  (4, 'p3');