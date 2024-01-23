package services

import (
	"errors"
	"fmt"

	"github.com/labora/labora-golang/ejercicios_integratorios/movies/models"
)

func CreateMovie(newMovie models.Movie) (int, error) {
	if *newMovie.Name == "" || *newMovie.Genre == "" || newMovie.ReleaseDate.IsZero() || *newMovie.Price == 0 {
		return -1, errors.New("All fields must be completed.")
	}

	stmt, err := DbConnection.Prepare("INSERT INTO movie(name, genre, release_date, acquired_date, stock, price) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id")
	if err != nil {
		return -1, err
	}

	var newMovieId int

	err = stmt.QueryRow(newMovie.Name, newMovie.Genre, newMovie.ReleaseDate, newMovie.AcquiredDate, newMovie.Stock, newMovie.Price).Scan(&newMovieId)

	if err != nil {
		return -1, err
	}

	return newMovieId, nil
}

func DeleteMovie(id int) error {
	stmt, err := DbConnection.Prepare("DELETE FROM movie WHERE id = $1")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}

func buildUpdateQuery(updatedMovie models.Movie) (string, error) {
	query := "UPDATE movie SET "
	paramsSize := 1
	if updatedMovie.Name != nil {
		query += fmt.Sprintf("name = $%d, ", paramsSize)
		paramsSize++
	}
	if updatedMovie.Genre != nil {
		query += fmt.Sprintf("genre = $%d, ", paramsSize)
		paramsSize++
	}
	if updatedMovie.ReleaseDate != nil {
		query += fmt.Sprintf("release_date = $%d, ", paramsSize)
		paramsSize++
	}
	if updatedMovie.AcquiredDate != nil {
		query += fmt.Sprintf("acquired_date = $%d, ", paramsSize)
		paramsSize++
	}
	if updatedMovie.Stock != nil {
		query += fmt.Sprintf("stock = $%d, ", paramsSize)
		paramsSize++
	}
	if updatedMovie.Price != nil {
		query += fmt.Sprintf("price = $%d, ", paramsSize)
		paramsSize++
	}
	if paramsSize == 1 {
		return "", fmt.Errorf("You must modify at least one field.")
	}

	query = query[:len(query)-2] + fmt.Sprintf(" WHERE id = $%d", paramsSize)
	return query, nil
}

func getNonNullFields(movie models.Movie) []interface{} {
	var nonNullFields []interface{}

	if movie.Id != nil {
		nonNullFields = append(nonNullFields, movie.Id)
	}
	if movie.Name != nil {
		nonNullFields = append(nonNullFields, movie.Name)
	}
	if movie.Genre != nil {
		nonNullFields = append(nonNullFields, movie.Genre)
	}
	if movie.ReleaseDate != nil {
		nonNullFields = append(nonNullFields, movie.ReleaseDate)
	}
	if movie.AcquiredDate != nil {
		nonNullFields = append(nonNullFields, movie.AcquiredDate)
	}
	if movie.Stock != nil {
		nonNullFields = append(nonNullFields, movie.Stock)
	}
	if movie.Price != nil {
		nonNullFields = append(nonNullFields, movie.Price)
	}

	return nonNullFields
}

func UpdateMovie(id int, updatedMovie models.Movie) error {
	query, err := buildUpdateQuery(updatedMovie)
	fmt.Println(query)
	if err != nil {
		return err
	}

	stmt, err := DbConnection.Prepare(query)
	if err != nil {
		return err
	}

	fields := getNonNullFields(updatedMovie)
	fields = append(fields, id)
	_, err = stmt.Exec(fields...)
	if err != nil {
		return err
	}

	return nil
}

func GetAllMovies() ([]models.Movie, error) {
	stmt, err := DbConnection.Prepare("SELECT * FROM movie")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movies []models.Movie

	for rows.Next() {
		var movie models.Movie
		err := rows.Scan(&movie.Id, &movie.Name, &movie.Genre, &movie.ReleaseDate, &movie.AcquiredDate, &movie.Stock, &movie.Price)
		if err != nil {
			return nil, err
		}
		movies = append(movies, movie)
	}

	// Check for errors from iterating over rows
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return movies, nil
}

func GetMoviesByName(name string) ([]models.Movie, error) {
	// Prepara la declaración SQL con la cláusula WHERE para buscar por nombre
	stmt, err := DbConnection.Prepare("SELECT * FROM movie WHERE name ILIKE $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	// Ejecuta la declaración SQL con el nombre proporcionado como argumento
	param := fmt.Sprintf("%%%s%%", name)
	rows, err := stmt.Query(param)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movies []models.Movie

	for rows.Next() {
		var movie models.Movie
		err := rows.Scan(&movie.Id, &movie.Name, &movie.Genre, &movie.ReleaseDate, &movie.AcquiredDate, &movie.Stock, &movie.Price)
		if err != nil {
			return nil, err
		}
		movies = append(movies, movie)
	}

	// Verifica errores al iterar sobre las filas
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return movies, nil
}

func GetMovieByID(id int) ([]models.Movie, error) {
	// Prepara la declaración SQL con la cláusula WHERE para buscar por ID
	stmt, err := DbConnection.Prepare("SELECT * FROM movie WHERE id = $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	// Ejecuta la declaración SQL con el ID proporcionado como argumento
	rows, err := stmt.Query(id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movies []models.Movie

	for rows.Next() {
		var movie models.Movie
		err := rows.Scan(&movie.Id, &movie.Name, &movie.Genre, &movie.ReleaseDate, &movie.AcquiredDate, &movie.Stock, &movie.Price)
		if err != nil {
			return nil, err
		}
		movies = append(movies, movie)
	}

	// Verifica errores al iterar sobre las filas
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return movies, nil
}
