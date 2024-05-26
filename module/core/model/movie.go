package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

type MovieModel struct {
	ID        string          `json:"id,omitempty" gorm:"column:id;default:uuid_generate_v4()"`
	Title     string          `json:"title,omitempty" gorm:"column:title"`
	Content   string          `json:"content,omitempty" gorm:"column:content"`
	Videos    *VideosDBStruct `json:"videos,omitempty" gorm:"column:videos"`
	CreatedAt *time.Time      `json:"created_at,omitempty" gorm:"column:created_at"`
	UpdatedAt *time.Time      `json:"updated_at,omitempty" gorm:"column:updated_at"`
}

func (MovieModel) TableName() string {
	return "movies"
}

type VideosDBStruct struct {
}

func (j *VideosDBStruct) GormDataType() string {
	return "jsonb"
}

// Value Marshal
func (j VideosDBStruct) Value() (driver.Value, error) {
	valueString, err := json.Marshal(j)
	return string(valueString), err
}

// Scan Unmarshal
func (j *VideosDBStruct) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	data, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(data, &j)
}
