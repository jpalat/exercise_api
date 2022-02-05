package models

import (
	"database/sql"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func ConnectDatabase() error {
	db, err := sql.Open("sqlite3", "./exercises.db")
	if err != nil {
		return err
	}

	DB = db
	return nil
}

type Record struct {
	Id       int    `json:"id"`
	Date     string `json:"date"`
	Exercise string `json:"exercise"`
	SetID    string `json:"setId"`
	Count    int
	Rest     int `json:"rest"`
	Style    string
}

func GetExercises(count int) ([]Record, error) {
	rows, err := DB.Query("Select id, date, exercise, setid, count, rest, style FROM records LIMIT " + strconv.Itoa(count))

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	workout := make([]Record, 0)
	for rows.Next() {
		set := Record{}
		err = rows.Scan(&set.Id, &set.Date, &set.Exercise, &set.SetID, &set.Count, &set.Rest, &set.Style)

		if err != nil {
			return nil, err
		}

		workout = append(workout, set)
	}

	err = rows.Err()

	if err != nil {
		return nil, err
	}

	return workout, err
}

// func GetPersonById(id string) (Person, error) {
// 	stmt, err := DB.Prepare("Select id, first_name, last_name, email, ip_address from people where id = ?")

// 	if err != nil {
// 		return Person{}, err
// 	}

// 	person := Person{}

// 	sqlErr := stmt.QueryRow(id).Scan(&person.Id, &person.FirstName, &person.LastName, &person.Email, &person.IpAddress)

// 	if sqlErr != nil {
// 		if sqlErr == sql.ErrNoRows {
// 			return Person{}, nil
// 		}
// 		return Person{}, sqlErr
// 	}
// 	return person, nil
// }
