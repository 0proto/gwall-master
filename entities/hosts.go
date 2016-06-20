package entities

import (
	"database/sql"
	"log"
	"time"
)

type Host struct {
	ID      int
	Name    string
	Address string
	Enabled bool
	Added   time.Time
}

// Create inserts new Host in DB
func (h *Host) Create(db *sql.DB) error {
	return db.QueryRow("INSERT INTO host (name, address, enabled, added) VALUES ($1, $2, $3, CURRENT_TIMESTAMP); SELECT last_insert_rowid() FROM host",
		h.Name, h.Address, h.Enabled).Scan(&h.ID)
}

// Update Host in DB
func (h *Host) Update(db *sql.DB) error {
	_, err := db.Exec("UPDATE host SET name = $1, address = $2, enabled = $3 WHERE id = $4",
		h.Name, h.Address, h.Enabled, h.ID)
	return err
}

func LoadAllHosts(ids string, db *sql.DB) []Host {
	var result []Host
	query := "SELECT id, name, address, enabled, added FROM host WHERE id IN ($1);"
	if ids == "" {
		query = "SELECT id, name, address, enabled, added FROM host;"
	}
	rows, err := db.Query(query, ids)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		host := Host{}
		if err := rows.Scan(&host.ID, &host.Name, &host.Address, &host.Enabled, &host.Added); err != nil {
			log.Fatal(err)
		}
		result = append(result, host)
	}

	return result
}

// Load Host struct by id
func (h *Host) Load(db *sql.DB) error {
	err := db.QueryRow("SELECT id, name, address, enabled, added FROM host WHERE id = $1 LIMIT 1",
		h.ID).Scan(&h.ID, &h.Name, &h.Address, &h.Enabled, &h.Added)
	return err
}

// Delete Host from DB
func (h *Host) Delete(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM host WHERE id = $1", h.ID)
	return err
}
