package sqlite

import (
	"database/sql"
	"fmt"
	"log/slog"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/telepuz/smsmanager/internal"
	"github.com/telepuz/smsmanager/internal/config"
)

const (
	DDLMessagesTableSQL = `CREATE TABLE IF NOT EXISTS messages (
		"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"message_id" INT64,
		"chat_id" INT64,
		"phone" TEXT,
		"content" TEXT,
		"dt" DATETIME DEFAULT current_timestamp
		);`
	DDLSendSmsTableSQL = `CREATE TABLE IF NOT EXISTS send_sms (
		"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"phone" TEXT,
		"username" TEXT,
		"text" INT64,
		"dt" DATETIME DEFAULT current_timestamp
		);`
	InsertNewMessageSQL  = `INSERT INTO messages(message_id, chat_id, phone, content) VALUES (?, ?, ?, ?)`
	GetMessagesCountSQL  = `SELECT COUNT(1) FROM messages;`
	IsItTimeToSendSmsSQL = `SELECT COUNT(1) FROM send_sms
		WHERE username = ?
		AND dt > ?`
	SaveSendSmsTimeSQL = `INSERT INTO send_sms(phone, username, text) VALUES (?, ?, ?)`
)

type SQLite struct {
	DB *sql.DB
}

func New(cfg *config.Config) (*SQLite, error) {
	slog.Debug(
		fmt.Sprintf(
			"Creating sqlite database: %s",
			cfg.Storage.DBFilePath,
		),
	)
	db, err := sql.Open("sqlite3", cfg.Storage.DBFilePath)
	if err != nil {
		return nil, fmt.Errorf("NewSQLiteStorage: %v", err)
	}

	slog.Debug("Creating messages table...")
	statement, err := db.Prepare(DDLMessagesTableSQL)
	if err != nil {
		return nil, fmt.Errorf("NewSQLiteStorage: %v", err)
	}
	_, err = statement.Exec()
	if err != nil {
		return nil, fmt.Errorf("NewSQLiteStorage: %v", err)
	}
	slog.Debug("messages table was created")

	slog.Debug("Creating send_sms table...")
	statement, err = db.Prepare(DDLSendSmsTableSQL)
	if err != nil {
		return nil, fmt.Errorf("NewSQLiteStorage: %v", err)
	}
	_, err = statement.Exec()
	if err != nil {
		return nil, fmt.Errorf("NewSQLiteStorage: %v", err)
	}
	slog.Debug("send_sms table was created")

	return &SQLite{DB: db}, nil
}

func (s *SQLite) DatabaseClose() error {
	return s.DB.Close()
}

func (s *SQLite) SaveMessage(message internal.Message, chatID int64) error {
	slog.Debug("Inserting message record...")
	statement, err := s.DB.Prepare(InsertNewMessageSQL)
	if err != nil {
		return fmt.Errorf("SaveMessage: %v", err)
	}
	_, err = statement.Exec(
		message.Index,
		chatID,
		message.Phone,
		message.Content,
	)
	if err != nil {
		return fmt.Errorf("SaveMessage: %v", err)
	}
	return nil
}

func (s *SQLite) GetMessagesCount() (int, error) {
	slog.Debug("Select message count...")
	count := 0
	err := s.DB.QueryRow(GetMessagesCountSQL).Scan(&count)
	if err != nil {
		return count, fmt.Errorf("GetMessagesCount: %v", err)
	}
	return count, nil
}

func (s *SQLite) IsItTimeToSendSms(username string, dt time.Duration) (bool, error) {
	slog.Debug("Select send_sms record...")

	count := 0
	filterDate := time.Now().UTC().Add(-1 * dt).Format("2006-01-02 15:04:05")

	err := s.DB.QueryRow(
		IsItTimeToSendSmsSQL,
		username,
		filterDate).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("IsItTimeToSendSms: %v", err)
	}

	if count > 0 {
		return false, nil
	}
	return true, nil
}

func (s *SQLite) SaveSendSmsTime(phone_number, username, text string) error {
	slog.Debug("Inserting send_sms record...")
	statement, err := s.DB.Prepare(SaveSendSmsTimeSQL)
	if err != nil {
		return fmt.Errorf("SaveSendSmsTime: %v", err)
	}
	_, err = statement.Exec(
		phone_number,
		username,
		text,
	)
	if err != nil {
		return fmt.Errorf("SaveSendSmsTime: %v", err)
	}
	return nil
}
