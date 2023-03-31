package click

import (
	"github.com/jmoiron/sqlx"
	"log"
	"time"
)

const connectionString = "tcp://localhost:9000?"

const getDataByUserId = `
SELECT * 
FROM default.my_first_table
WHERE user_id = ?;`

const addNewRow = `
INSERT INTO default.my_first_table
    (user_id, message, timestamp, metric) 
	values ($1, $2, $3, $4);
`

type ClickRepo struct {
	Db *sqlx.DB
}

type TestData struct {
	User_id   int
	Message   string
	Timestamp time.Time
	Metric    float32
}

func Read(searchId int) (result []TestData) {
	repo := ClickRepo{Db: sqlx.MustConnect("clickhouse", connectionString)}

	err := repo.Db.Select(&result, getDataByUserId, searchId)
	if err != nil {
		log.Println(err)
		return nil
	}

	return
}

func Write(user_id int, message string, metric float32) {
	repo := ClickRepo{Db: sqlx.MustConnect("clickhouse", connectionString)}

	tx, err := repo.Db.Begin()
	if err != nil {
		log.Println(err)
		return
	}

	defer tx.Rollback()
	_, err = tx.Exec(addNewRow, user_id, message, time.Now(), metric)
	if err != nil {
		log.Println(err)
		return
	}

	err = tx.Commit()
	if err != nil {
		log.Println(err)
	}
}
