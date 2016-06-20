package entities

import "time"
import "database/sql"

type Notification struct {
	ID         int
	FromAlert  *Alert
	Recepients []User
	Date       time.Time
}

// Create inserts new Notification in DB
func (n *Notification) Create(db *sql.DB) error {
	var userIds []int
	for _, user := range n.Recepients {
		userIds = append(userIds, user.ID)
	}
	sRecepients := serializeIds(userIds)
	return db.QueryRow("INSERT INTO notification (from_alert, recepients, date) VALUES ($1, $2, CURRENT_TIMESTAMP); SELECT last_insert_rowid() FROM notification",
		n.FromAlert.ID, sRecepients).Scan(&n.ID)
}

/*
// Update Notification in DB
func (u *Notification) UpdateEmail(db *sql.DB) error {
	_, err := db.Exec("UPDATE user SET email = $1 WHERE id = $2",
		u.Email, u.ID)
	return err
}
*/

// Load Notification struct by id
func (n *Notification) Load(db *sql.DB) error {
	var fromAlertId int
	var sRecepients string
	err := db.QueryRow("SELECT id, from_alert, recepients, date FROM notification WHERE id = $1 LIMIT 1",
		n.ID).Scan(&n.ID, &fromAlertId, &sRecepients, &n.Date)
	al := Alert{ID: fromAlertId}
	al.Load(db)
	n.FromAlert = &al

	n.Recepients = LoadAllUsers(sRecepients, db)
	return err
}

// Delete Notification from DB
func (n *Notification) Delete(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM notification WHERE id = $1", n.ID)
	return err
}
