package core

import (
	"errors"
	"math/rand"
	"time"

	"github.com/yannis94/black-jack/internal/helpers"
)

type Hoof struct {
    Cards *helpers.Queue
}

const USED_DECK int = 6

func BuildHoof() *Hoof {
    var cards []*Card
    deck := NewDeck()

    for i:=0; i < USED_DECK; i++ {
        cards = append(cards, deck.Cards...)
    }

    shuffuledDecks := make([]Card, len(cards))
    rand.Seed(time.Now().UTC().UnixNano())
    perm := rand.Perm(len(cards))

    for i, card := range perm {
        shuffuledDecks[card] = *cards[i]
    }

    hoofCardQueue := helpers.NewQueue()

    for j:=0; j < len(shuffuledDecks); j++ {
        hoofCardQueue.Enqueue(shuffuledDecks[j])
    }

    return &Hoof{ Cards: hoofCardQueue }
}

func (hoof *Hoof) GetCard() (Card, error) {
    data, err := hoof.Cards.Dequeue()

    if err != nil {
        return Card{}, err
    }

    if card, ok := data.(Card); !ok {
        return Card{}, errors.New("Not a card returned by hoof")
    } else {
        return card, nil
    }
}
