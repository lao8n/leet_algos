

type Suit string

const (
	Heart    Suit = "heart"
	Clubs    Suit = "clubs"
	Diamonds Suit = "diamonds"
	Aces     Suit = "aces"
)

type Value int

const (
	One Value iota
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten 
	Jack 
	Queen
	King
)

type Card struct {
	suit Suit
	value Value
}

type Deck struct {
	cards [52]Card
}

func (d *Deck) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d.cards), func(i, j int) {
		d.card[i], d.card[j] = d.card[j], d.card[i]
	})
}