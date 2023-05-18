package core

import (
	"errors"

)

type Table struct {
    Id int
    Sits []*Player
    MaxPlayer int
    Hoof *Hoof
}

func NewTable(id int) *Table {
    return &Table{
        Id: id,
        Hoof: BuildHoof(),
        MaxPlayer: 5,
        Sits: make([]*Player, 5),
    }
}

func (table *Table) AddPlayer(player *Player) error {
    if table.MaxPlayer == len(table.Sits) {
        return errors.New("Table is full, please go to another one.")
    }

    table.Sits = append(table.Sits, player)
    return nil
}

func (table *Table) RemovePlayer(id int) Player {
    var removedPlayer *Player

    for i, player := range table.Sits {
        if player.Id == id {
            removedPlayer = player
            table.Sits = append(table.Sits[:i], table.Sits[i+1:]...)
        }
    }

    return *removedPlayer
}
