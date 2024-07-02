// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package dbrepo

import (
	"time"
)

type Devices struct {
	ID               int32     `db:"id" json:"id"`
	ProjectID        int32     `db:"project_id" json:"project_id"`
	Name             string    `db:"name" json:"name"`
	PhoneNumber      *string   `db:"phone_number" json:"phone_number"`
	Description      *string   `db:"description" json:"description"`
	IsActive         bool      `db:"is_active" json:"is_active"`
	CustomAttributes []byte    `db:"custom_attributes" json:"custom_attributes"`
	CreatedAt        time.Time `db:"created_at" json:"created_at"`
	UpdatedAt        time.Time `db:"updated_at" json:"updated_at"`
	Jid              *string   `db:"jid" json:"jid"`
	ClientDeviceID   string    `db:"client_device_id" json:"client_device_id"`
}

type Projects struct {
	ID          int32     `db:"id" json:"id"`
	Name        string    `db:"name" json:"name"`
	Alias       string    `db:"alias" json:"alias"`
	Description *string   `db:"description" json:"description"`
	Settings    []byte    `db:"settings" json:"settings"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}
