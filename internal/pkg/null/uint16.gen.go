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

var _ json.Unmarshaler = (*Uint16)(nil)

// Uint16 is a nullable type.
type Uint16 struct {
	Value uint16
	Set   bool // if json object has this field
}

// NewUint16 creates a new uint16.
func NewUint16(v uint16) Uint16 {
	return Uint16{
		Value: v,
		Set:   true,
	}
}

func (t Uint16) Ptr() *uint16 {
	if t.Set {
		return &t.Value
	}

	return nil
}

// Default return default value its value is Null or not Set.
func (t Uint16) Default(v uint16) uint16 {
	if t.Set {
		return t.Value
	}

	return v
}

func (t Uint16) Interface() any {
	if t.Set {
		return t.Value
	}

	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (t *Uint16) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}

	t.Set = true
	return json.UnmarshalNoEscape(data, &t.Value) //nolint:wrapcheck
}