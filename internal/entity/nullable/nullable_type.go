package nullable

import (
	"database/sql"
	"encoding/json"
)

type NullString struct {
	sql.NullString
}

func (s NullString) MarshalJSON() ([]byte, error) {
	if s.Valid {
		return json.Marshal(s.String)
	}

	return []byte("null"), nil
}

func (s *NullString) UnmarshalJSON(b []byte) error {
	s.Valid = string(b) != "null"
	return json.Unmarshal(b, &s.String)
}

type NullTime struct {
	sql.NullTime
}

func (t NullTime) MarshalJSON() ([]byte, error) {
	if t.Valid {
		return json.Marshal(t.Time)
	}

	return []byte("null"), nil
}

func (t *NullTime) UnmarshalJSON(b []byte) error {
	t.Valid = string(b) != "null"
	return json.Unmarshal(b, &t.Time)
}
