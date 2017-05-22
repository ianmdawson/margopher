package margopher

import "testing"

func TestGetRandomWord(t *testing.T) {
	if getRandomWord([]string{"1", "2", "3"}) == "" {
		t.Error("getRandomWord: it should return a string element from slice.")
	}
}

func TestIsTerminalWord(t *testing.T) {
	if isTerminalWord("Hey.") == false {
		t.Error("isTerminalWord: it should return true for words ending in period.")
	}
	if isTerminalWord("Hey;") == false {
		t.Error("isTerminalWord: it should return true for words ending in semicolon.")
	}
	if isTerminalWord("Hey?") == false {
		t.Error("isTerminalWord: it should return true for words ending in question mark.")
	}
	if isTerminalWord("Hey!") == false {
		t.Error("isTerminalWord: it should return true for words ending in exclamation point.")
	}
}

func TestReadText(t *testing.T) {
	m := New()
	m.ParseText("Cats are wonderful. Cats love tuna.")

	if m.states == nil {
		t.Error("ParseText: it should initialize states.")
	}
}
