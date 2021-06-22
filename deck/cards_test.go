package deck

import (
	"fmt"
	"testing"
)

func ExmplCard() {
	fmt.Println(Card{Suit: Heart, Rank: Ace})
	fmt.Println(Card{Suit: Spade, Rank: Nine})
	fmt.Println(Card{Suit: Diamond, Rank: Two})
	fmt.Println(Card{Suit: Club, Rank: Jack})
	fmt.Println(Card{Suit: Joker})

	// Output:
	// Ace of Hearts
	// Nine of Spades
	// Two of Diamonds
	// Jack of Clubs
	// Joker
}

func TestDeck(t *testing.T) {
	cards := New()
	if len(cards) != 13*4 {
		t.Error("Wrong number of cards in a new deck.")
	}
}

func TestDefaultSort(t *testing.T) {
	cards := New(DefaultSort)
	exp := Card{Rank: Ace, Suit: Spade}
	if cards[0] != exp {
		t.Error("Expected Ace of Spades, recieved: ", cards[0])
	}
}

func TestJokers(t *testing.T) {
	cards := New(Jokers(3))
	count := 0
	for _, c := range cards {
		if c.Suit == Joker {
			count++
		}
	}
	if count != 3 {
		t.Error("Expected 3 Jokers in a deck, got:", count)
	}
}

func TestFilter(t *testing.T) {
	filterTwoThree := func(card Card) bool {
		return card.Rank == Two || card.Rank == Three
	}
	cards := New(Filter(filterTwoThree))
	for _, c := range cards {
		if c.Rank == Two || c.Rank == Three {
			t.Error("Twos or Threes are not expected in the deck.")
		}
	}
}

func TestDecks(t *testing.T) {
	cards := New(Decks(3))
	if len(cards) != 13*4*3 {
		t.Error("Expected 156 cards, got:", len(cards))
	}
}

func TestShuffle(t *testing.T) {
	seed := int64(0)
	cards := New(Shuffle(seed))
	orig := New()
	first := orig[40]
	second := orig[35]
	if cards[0] != first {
		t.Errorf("Expected the first card to be %s, got: %s", first, cards[0])
	}
	if cards[1] != second {
		t.Errorf("Expected the second card to be %s, got: %s", second, cards[1])
	}
}
