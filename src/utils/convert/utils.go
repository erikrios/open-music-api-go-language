package convert

import "database/sql"

func ToNullString(value *string) sql.NullString {
	if value == nil {
		return sql.NullString{Valid: false}
	}
	return sql.NullString{Valid: true, String: *value}
}

func ToNullInt16(value *uint16) sql.NullInt16 {
	if value == nil {
		return sql.NullInt16{Valid: false}
	}
	return sql.NullInt16{Valid: true, Int16: int16(*value)}

}

func FromNullString(value sql.NullString) *string {
	if value.Valid {
		return &value.String
	}
	return nil
}

func FromNullInt16(value sql.NullInt16) *uint16 {
	if value.Valid {
		result := uint16(value.Int16)
		return &result
	}
	return nil
}
