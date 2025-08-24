package sqlite

import (
	"database/sql"
	"fmt"
	"log/slog"

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
	InsertNewMessageSQL = `INSERT INTO messages(message_id, chat_id, phone, content) VALUES (?, ?, ?, ?)`
	GetMessagesCountSQL = `SELECT COUNT(1) FROM messages;`
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
		return nil, fmt.Errorf("SaveMessage: %v", err)
	}
	slog.Debug("Messages table created")

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
	count := 0
	err := s.DB.QueryRow(GetMessagesCountSQL).Scan(&count)
	if err != nil {
		return count, fmt.Errorf("GetMessagesCount: %v", err)
	}
	return count, nil
}
