package entities

import (
	"database/sql"
	"time"
)

type Usergroup struct {
	ID    int
	Name  string
	Email string
	Added time.Time
}

// Create inserts new Usergroup in DB
func (u *Usergroup) Create(db *sql.DB) error {
	return db.QueryRow("INSERT INTO usergroup (name, email, added) VALUES ($1, $2, CURRENT_TIMESTAMP); SELECT last_insert_rowid() FROM usergroup",
		u.Name, u.Email, u.Added).Scan(&u.ID)
}

// Update Usergroup in DB
func (u *Usergroup) Update(db *sql.DB) error {
	_, err := db.Exec("UPDATE usergroup SET email = $1, name = $2 WHERE id = $3",
		u.Email, u.Name, u.ID)
	return err
}

// Load Usergroup struct by id
func (u *Usergroup) Load(db *sql.DB) error {
	err := db.QueryRow("SELECT id, name, email, added FROM usergroup WHERE id = $1 LIMIT 1",
		u.ID).Scan(&u.ID, &u.Name, &u.Email, &u.Added)
	return err
}

// DeleteUsergroupUser from DB
func (u *Usergroup) Delete(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM usergroup WHERE id = $1", u.ID)
	return err
}
