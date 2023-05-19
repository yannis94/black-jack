package core

import (
	"log"

	"github.com/yannis94/black-jack/internal/helpers"
)

type Game struct {
    tour int
    pool *helpers.Queue
    hoof *Hoof
    dealer *Player
}

func NewGame(players []*Player, hoof *Hoof) *Game {
    playersQueue := helpers.NewQueue()

    for _, player := range players {
        playersQueue.Enqueue(player)
    }

    dealer := NewPlayer(0, "Dealer", 100000)
    playersQueue.Enqueue(dealer)

    return &Game{
        tour: 0,
        pool: playersQueue,
        hoof: hoof,
        dealer: dealer,
    }
}

func (game *Game) cardDistribution() {
    cardDistribQueue := game.pool
    
    for cardDistribQueue.Length > 0 {
        tmp, _ := cardDistribQueue.Dequeue()
        nextPlayer, ok := tmp.(*Player)

        if !ok {
            log.Println("No player fond")
            break
        }

        if nextPlayer.Wallet < 1 {
            cardDistribQueue.Dequeue()
            continue
        }

        card, _ := game.hoof.GetCard()

        nextPlayer.Hand = append(nextPlayer.Hand, card)

        if len(nextPlayer.Hand) < 2 {
            nextPlayer.Playing = true
            cardDistribQueue.Enqueue(nextPlayer)
        }
    }
}

func (game *Game) collectBet() {
    playersQueue := game.pool

    for playersQueue.Length > 0 {
        tmp, _ := playersQueue.Dequeue()
        nextPlayer, ok := tmp.(*Player)

        if !ok {
            log.Println("No player fond")
            break
        } else if nextPlayer.Username == "Dealer" {
            continue
        }

        //make player choice
                
    }
}

func (game *Game) newTour() {
    playersQueue := game.pool

    for playersQueue.Length > 0 {
        tmp, _ := playersQueue.Dequeue()
        nextPlayer, ok := tmp.(*Player)

        if !ok {
            log.Println("No player fond")
            break
        } else if nextPlayer.Username == "Dealer" {
            log.Println("make dealer next move function")
            continue
        }

                
    }
}

func (game *Game) Start() {
    game.cardDistribution()
    game.collectBet()

    for game.dealer.Playing {
        game.newTour()
    }
}
