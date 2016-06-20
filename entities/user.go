package entities

import "log"
import "time"
import "database/sql"

type User struct {
	ID       int       `json:"id"`
	Login    string    `json:"login"`
	Email    string    `json:"email"`
	Password string    `json:"-"`
	Added    time.Time `json:"added"`
}

// Create inserts new User in DB
func (u *User) Create(db *sql.DB) error {
	return db.QueryRow("INSERT INTO user (login, email, password, added) VALUES ($1, $2, $3, CURRENT_TIMESTAMP); SELECT last_insert_rowid() FROM user",
		u.Login, u.Email, u.Password).Scan(&u.ID)
}

// Update User in DB
func (u *User) UpdateEmail(db *sql.DB) error {
	_, err := db.Exec("UPDATE user SET email = $1 WHERE id = $2",
		u.Email, u.ID)
	return err
}

// Load User struct by id
func (u *User) Load(db *sql.DB) error {
	err := db.QueryRow("SELECT id, login, email, added FROM user WHERE id = $1 LIMIT 1",
		u.ID).Scan(&u.ID, &u.Login, &u.Email, &u.Added)
	return err
}

func LoadAllUsers(ids string, db *sql.DB) []User {
	var result []User
	query := "SELECT id, login, email, added FROM user WHERE id IN ($1)"
	if ids == "" {
		query = "SELECT id, login, email, added FROM user;"
	}
	rows, err := db.Query(query, ids)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		user := User{}
		if err := rows.Scan(&user.ID, &user.Login, &user.Email, &user.Added); err != nil {
			log.Fatal(err)
		}
		result = append(result, user)
	}

	return result
}

// Delete User from DB
func (u *User) Delete(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM user WHERE id = $1", u.ID)
	return err
}
