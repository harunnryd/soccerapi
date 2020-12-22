package model

import "github.com/satori/uuid"

// Player is an `players` table abstractions.
type Player struct {
	Model
	TeamID uuid.UUID
	Name   string
}
