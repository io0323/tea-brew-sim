/*
TeaBrewSim データベース管理
SQLiteを使用したデータ永続化
*/
package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

// BrewingRecord 抽出記録
type BrewingRecord struct {
	ID          int64     `json:"id"`
	TeaType     string    `json:"tea_type"`
	Temperature float64   `json:"temperature"`
	Time        int       `json:"time"`
	Rating      float64   `json:"rating"`
	Comment     string    `json:"comment"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// Database データベース管理構造体
type Database struct {
	db *sql.DB
}

// NewDatabase 新しいデータベース接続を作成
func NewDatabase(dbPath string) (*Database, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("データベース接続に失敗: %v", err)
	}

	database := &Database{db: db}
	
	// テーブルを作成
	if err := database.createTables(); err != nil {
		return nil, fmt.Errorf("テーブル作成に失敗: %v", err)
	}

	return database, nil
}

// createTables テーブルを作成
func (d *Database) createTables() error {
	query := `
	CREATE TABLE IF NOT EXISTS brewing_records (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		tea_type TEXT NOT NULL,
		temperature REAL NOT NULL,
		time INTEGER NOT NULL,
		rating REAL DEFAULT 0.0,
		comment TEXT DEFAULT '',
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);

	CREATE TABLE IF NOT EXISTS tea_types (
		id TEXT PRIMARY KEY,
		name TEXT NOT NULL,
		min_temp REAL NOT NULL,
		max_temp REAL NOT NULL,
		min_time INTEGER NOT NULL,
		max_time INTEGER NOT NULL,
		description TEXT DEFAULT '',
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);

	CREATE INDEX IF NOT EXISTS idx_brewing_records_tea_type ON brewing_records(tea_type);
	CREATE INDEX IF NOT EXISTS idx_brewing_records_created_at ON brewing_records(created_at);
	`

	_, err := d.db.Exec(query)
	return err
}

// Close データベース接続を閉じる
func (d *Database) Close() error {
	return d.db.Close()
}

// SaveBrewingRecord 抽出記録を保存
func (d *Database) SaveBrewingRecord(record *BrewingRecord) error {
	query := `
	INSERT INTO brewing_records (tea_type, temperature, time, rating, comment, created_at, updated_at)
	VALUES (?, ?, ?, ?, ?, ?, ?)
	`

	now := time.Now()
	result, err := d.db.Exec(query, record.TeaType, record.Temperature, record.Time, 
		record.Rating, record.Comment, now, now)
	if err != nil {
		return fmt.Errorf("抽出記録の保存に失敗: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("IDの取得に失敗: %v", err)
	}

	record.ID = id
	record.CreatedAt = now
	record.UpdatedAt = now

	return nil
}

// GetBrewingRecords 抽出記録一覧を取得
func (d *Database) GetBrewingRecords(limit, offset int) ([]*BrewingRecord, error) {
	query := `
	SELECT id, tea_type, temperature, time, rating, comment, created_at, updated_at
	FROM brewing_records
	ORDER BY created_at DESC
	LIMIT ? OFFSET ?
	`

	rows, err := d.db.Query(query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("抽出記録の取得に失敗: %v", err)
	}
	defer rows.Close()

	var records []*BrewingRecord
	for rows.Next() {
		record := &BrewingRecord{}
		err := rows.Scan(&record.ID, &record.TeaType, &record.Temperature, 
			&record.Time, &record.Rating, &record.Comment, 
			&record.CreatedAt, &record.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("レコードのスキャンに失敗: %v", err)
		}
		records = append(records, record)
	}

	return records, nil
}

// GetBrewingRecordByID IDで抽出記録を取得
func (d *Database) GetBrewingRecordByID(id int64) (*BrewingRecord, error) {
	query := `
	SELECT id, tea_type, temperature, time, rating, comment, created_at, updated_at
	FROM brewing_records
	WHERE id = ?
	`

	record := &BrewingRecord{}
	err := d.db.QueryRow(query, id).Scan(&record.ID, &record.TeaType, 
		&record.Temperature, &record.Time, &record.Rating, 
		&record.Comment, &record.CreatedAt, &record.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("記録が見つかりません")
		}
		return nil, fmt.Errorf("抽出記録の取得に失敗: %v", err)
	}

	return record, nil
}

// UpdateBrewingRecord 抽出記録を更新
func (d *Database) UpdateBrewingRecord(record *BrewingRecord) error {
	query := `
	UPDATE brewing_records
	SET tea_type = ?, temperature = ?, time = ?, rating = ?, comment = ?, updated_at = ?
	WHERE id = ?
	`

	now := time.Now()
	result, err := d.db.Exec(query, record.TeaType, record.Temperature, 
		record.Time, record.Rating, record.Comment, now, record.ID)
	if err != nil {
		return fmt.Errorf("抽出記録の更新に失敗: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("更新行数の取得に失敗: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("更新対象の記録が見つかりません")
	}

	record.UpdatedAt = now
	return nil
}

// DeleteBrewingRecord 抽出記録を削除
func (d *Database) DeleteBrewingRecord(id int64) error {
	query := `DELETE FROM brewing_records WHERE id = ?`

	result, err := d.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("抽出記録の削除に失敗: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("削除行数の取得に失敗: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("削除対象の記録が見つかりません")
	}

	return nil
}

// GetStatistics 統計情報を取得
func (d *Database) GetStatistics() (map[string]interface{}, error) {
	stats := make(map[string]interface{})

	// 総抽出回数
	var totalBrews int
	err := d.db.QueryRow("SELECT COUNT(*) FROM brewing_records").Scan(&totalBrews)
	if err != nil {
		return nil, fmt.Errorf("総抽出回数の取得に失敗: %v", err)
	}
	stats["total_brews"] = totalBrews

	// 平均評価
	var avgRating float64
	err = d.db.QueryRow("SELECT AVG(rating) FROM brewing_records WHERE rating > 0").Scan(&avgRating)
	if err != nil {
		return nil, fmt.Errorf("平均評価の取得に失敗: %v", err)
	}
	stats["average_rating"] = avgRating

	// 人気のお茶の種類
	var favoriteTea string
	err = d.db.QueryRow(`
		SELECT tea_type FROM brewing_records 
		GROUP BY tea_type 
		ORDER BY COUNT(*) DESC 
		LIMIT 1
	`).Scan(&favoriteTea)
	if err != nil {
		return nil, fmt.Errorf("人気のお茶の取得に失敗: %v", err)
	}
	stats["favorite_tea"] = favoriteTea

	return stats, nil
}

// GetBrewingRecordsByTeaType お茶の種類別の抽出記録を取得
func (d *Database) GetBrewingRecordsByTeaType(teaType string, limit int) ([]*BrewingRecord, error) {
	query := `
	SELECT id, tea_type, temperature, time, rating, comment, created_at, updated_at
	FROM brewing_records
	WHERE tea_type = ?
	ORDER BY created_at DESC
	LIMIT ?
	`

	rows, err := d.db.Query(query, teaType, limit)
	if err != nil {
		return nil, fmt.Errorf("抽出記録の取得に失敗: %v", err)
	}
	defer rows.Close()

	var records []*BrewingRecord
	for rows.Next() {
		record := &BrewingRecord{}
		err := rows.Scan(&record.ID, &record.TeaType, &record.Temperature, 
			&record.Time, &record.Rating, &record.Comment, 
			&record.CreatedAt, &record.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("レコードのスキャンに失敗: %v", err)
		}
		records = append(records, record)
	}

	return records, nil
}
