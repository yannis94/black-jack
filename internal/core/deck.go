package core

type Card struct {
    Value string
    Family string
}

type Deck struct {
    Cards []*Card
}

func NewDeck() Deck {
    var cards []*Card
    families := []string{"club", "diamond", "heart", "spade"}
    values := []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}

    for _, family := range families {
        for _, val := range values {
            card := &Card{
                Value: val,
                Family: family,
            }
            cards = append(cards, card)
        }
    }

    return Deck{
        Cards: cards,
    }
}
