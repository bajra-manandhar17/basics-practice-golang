package deck

import (
	"os"
	"testing"
)

/**
* The *testing.T is the test handler that we can use to do things like log out errors if we find any.
* We can also use it to tell Go that something went wrong.
*
* The Test prefix is important because it tells Go that this is a test that we want to run.
 */

func TestNewDeck(t *testing.T) {
	d := NewDeck()

	/** If the length of the deck is not 52, then something went wrong */
	if len(d) != 52 {
		t.Errorf("Expected a deck length of 52 but got %v", len(d))
	}

	/** If the first card is not Ace of Spades, then something went wrong */
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be Ace of Spades but got %v", d[0])
	}

	/** If the last card is not King of Clubs, then something went wrong */
	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected last card to be King of Clubs but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	/** Delete any existing files */
	_ = os.Remove("_decktesting")

	d := NewDeck()
	d.SaveToFile("_decktesting")

	/** Load the deck from the file */
	loadedDeck := NewDeckFromFile("_decktesting")

	/** If the length of the deck is not 52, then something went wrong */
	if len(loadedDeck) != 52 {
		t.Errorf("Expected a deck length of 52 but got %v", len(loadedDeck))
	}

	_ = os.Remove("_decktesting")
}
