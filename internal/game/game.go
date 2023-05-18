package game

import "github.com/yannis94/black-jack/internal/helpers"

type Game struct {
    tour int
    Pool *helpers.Queue
}
