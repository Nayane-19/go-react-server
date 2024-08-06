// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package pgstore

import (
	"github.com/google/uuid"
)

type Message struct {
	ID            uuid.UUID `db:"id" json:"id"`
	RoomID        uuid.UUID `db:"room_id" json:"room_id"`
	Message       string    `db:"message" json:"message"`
	ReactionCount int64     `db:"reaction_count" json:"reaction_count"`
	Answered      bool      `db:"answered" json:"answered"`
}

type Room struct {
	ID    uuid.UUID `db:"id" json:"id"`
	Theme string    `db:"theme" json:"theme"`
}

type Teste struct {
	ID    uuid.UUID `db:"id" json:"id"`
	Theme string    `db:"theme" json:"theme"`
}