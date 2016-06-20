package entities

import (
	"database/sql"
	"strconv"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

const (
	PortScanAlert = iota
	AnomalyAlert  = iota
)

// Alert db representation
type Alert struct {
	ID        int
	AlertType int
	HostID    int
	Priority  int
	Title     string
	Message   string
	Date      time.Time
}

func serializeIds(groups []int) string {
	var result string

	for _, i := range groups {
		j := strconv.Itoa(i)
		result += "," + j
	}

	result = result[1:]
	return result
}

func deserializeIds(groups string) []int {
	var result []int
	var err error

	sgroups := strings.Split(groups, ",")

	for i, v := range sgroups {
		result[i], err = strconv.Atoi(v)
		if err != nil {
			return nil
		}
	}

	return result
}

// Create inserts new Alert in DB
func (a *Alert) Create(db *sql.DB) error {
	return db.QueryRow("INSERT INTO alert (host_id, priority, title, message, date) VALUES ($1, $2, $3, $4, CURRENT_TIMESTAMP); SELECT last_insert_rowid() FROM alert",
		a.HostID, a.Priority, a.Title, a.Message, a.Date).Scan(&a.ID)
}

/*
// Update PlayerReward in DB
func (a *Alert) Update(db *sql.DB) error {
	_, err := conf.Db.Exec("UPDATE alert SET rewardDay = $1, lastRewardDate = $2 WHERE id = $3",
		pr.RewardDay, pr.LastRewardDate, pr.ID)
	return err
}
*/

// Load values int Alert struct by id
func (a *Alert) Load(db *sql.DB) error {
	err := db.QueryRow("SELECT id, host_id, priority, title, message, date FROM alert WHERE id = $1 LIMIT 1",
		a.ID).Scan(&a.ID, &a.HostID, &a.Priority, &a.Title, &a.Message, &a.Date)
	return err
}

// Delete Alert from DB
func (a *Alert) Delete(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM alert WHERE playerID = $1", a.ID)
	return err
}
