package frenchspace

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestText(t *testing.T) {
	assert.Equal(t, "5\u202F% swag", Text("5% swag"))
	assert.Equal(t, "5\u202F% swag", Text("5 % swag"))
	assert.Equal(t, "yolo\u202F; swag", Text("yolo; swag"))
	assert.Equal(t, "yolo\u202F; swag", Text("yolo ; swag"))
	assert.Equal(t, "yolo\u202F! swag", Text("yolo ! swag"))
	assert.Equal(t, "yolo\u202F! swag", Text("yolo! swag"))
	assert.Equal(t, "yolo\u202F? swag", Text("yolo ? swag"))
	assert.Equal(t, "yolo\u202F? swag", Text("yolo? swag"))

	assert.Equal(t, "«\u202Fyolo\u202F» swag", Text("«yolo» swag"))
	assert.Equal(t, "«\u202Fyolo\u202F» swag", Text("« yolo » swag"))
	assert.Equal(t, "‹\u202Fyolo\u202F› swag", Text("‹yolo› swag"))
	assert.Equal(t, "‹\u202Fyolo\u202F› swag", Text("‹ yolo › swag"))
	assert.Equal(t, "“\u202Fyolo\u202F” swag", Text("“yolo” swag"))
	assert.Equal(t, "“\u202Fyolo\u202F” swag", Text("“ yolo ” swag"))

	assert.Equal(t, "yolo\u00A0:", Text("yolo:"))
	assert.Equal(t, "yolo\u00A0:", Text("yolo :"))

	assert.Equal(t, "Bonjour tout le monde\u202F! Ça va\u202F?",
		Text("Bonjour  tout  le monde!\n Ça va?"))
}

func TestSpaces(t *testing.T) {
	assert.Equal(t, "bonjour tout le monde", spaces(" \n \t   bonjour  tout  le monde \n"))
}
