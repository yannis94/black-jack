package core

import "errors"

type Player struct {
    Id int
    Username string
    Hand []Card
    Wallet int
    Bet int
    Playing bool
}

func NewPlayer(id int, username string, wallet int) *Player {
    return &Player{ Id: id, Username: username, Wallet: wallet }
}

func (p *Player) UpdateWallet(amount int) error {
    if p.Wallet + amount < 0 {
        return errors.New("Player could not have a negative wallet.")
    }

    p.Wallet = p.Wallet + amount
    return nil
}

func (p *Player) MakeBet(amount int) error {
    err := p.UpdateWallet(amount)

    if err != nil {
        return err
    }

    p.Bet = amount

    return nil
}

