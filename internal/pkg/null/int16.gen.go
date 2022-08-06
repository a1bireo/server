// Code generated by ./internal/cmd/gen/null/main.go. DO NOT EDIT.

// SPDX-License-Identifier: AGPL-3.0-only
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published
// by the Free Software Foundation, version 3.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
// See the GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <https://www.gnu.org/licenses/>

package null

import (
	"github.com/goccy/go-json"
)

var _ json.Unmarshaler = (*Int16)(nil)

// Int16 is a nullable type.
type Int16 struct {
	Value int16
	Set   bool // if json object has this field
}

// NewInt16 creates a new int16.
func NewInt16(v int16) Int16 {
	return Int16{
		Value: v,
		Set:   true,
	}
}

func (t Int16) Ptr() *int16 {
	if t.Set {
		return &t.Value
	}

	return nil
}

// Default return default value its value is Null or not Set.
func (t Int16) Default(v int16) int16 {
	if t.Set {
		return t.Value
	}

	return v
}

func (t Int16) Interface() any {
	if t.Set {
		return t.Value
	}

	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (t *Int16) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}

	t.Set = true
	return json.UnmarshalNoEscape(data, &t.Value) //nolint:wrapcheck
}