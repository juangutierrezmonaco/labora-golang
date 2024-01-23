CREATE DATABASE labora_movies;

CREATE TABLE IF NOT EXISTS movie (
  id SERIAL PRIMARY KEY,
  name VARCHAR(150) NOT NULL,
  genre VARCHAR(150) NOT NULL,
  release_date DATE NOT NULL,
  acquired_date DATE,
  stock INT DEFAULT 1,
  price DECIMAL(10,2) NOT NULL
);