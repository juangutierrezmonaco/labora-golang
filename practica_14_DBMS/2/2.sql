CREATE TABLE movie(
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  created DATE,
  stock BIGINT,
  price FLOAT NOT NULL
);

INSERT INTO movie(name, created, stock, price) VALUES (1,1704398539,20,60);
INSERT INTO movie VALUES ('2','El señor de los anillos',20);
INSERT INTO movie(id, name, price) VALUES (3,'Las dos torres',240000);
INSERT INTO movie(id, name, created) VALUES (4,'El retorno del rey',1600000000);