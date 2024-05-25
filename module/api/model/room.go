package model

import (
	"database/sql/driver"
	"fmt"
	"strings"
	"time"
)

type RoomModel struct {
	ID        int64      `json:"id,omitempty" gorm:"column:id"`
	Name      string     `json:"name,omitempty" gorm:"column:name"`
	CreatedAt *time.Time `json:"created_at,omitempty" gorm:"column:created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" gorm:"column:updated_at"`
}

func (RoomModel) TableName() string {
	return "rooms"
}

type SeatsSlide []string

// Value converts the EnumSlice to a format that can be stored in the database.
func (a SeatsSlide) Value() (driver.Value, error) {
	// Convert the []MyEnum slice to a string representation of the enum array.
	values := make([]string, len(a))
	for i, val := range a {
		values[i] = string(val)
	}
	return "{" + strings.Join(values, ",") + "}", nil
}

// Scan converts the database value back to an EnumSlice.
func (a *SeatsSlide) Scan(value interface{}) error {
	// Check if the value is nil or an empty array.
	if value == nil {
		*a = []string{}
		return nil
	}

	// Convert the database value to a string.
	strValue, ok := value.(string)
	if !ok {
		return fmt.Errorf("failed to scan EnumSlice: value is not a string")
	}

	// Remove the curly braces and split the string into individual values.
	strValue = strings.Trim(strValue, "{}")
	values := strings.Split(strValue, ",")

	// Assign the values to the EnumSlice.
	*a = make([]string, len(values))
	for i, val := range values {
		(*a)[i] = string(val)
	}
	return nil
}
