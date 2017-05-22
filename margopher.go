package margopher

import (
	"bytes"
	"strings"
)

type margopher struct {
	states map[[2]string][]string
}

// New Margopher constructor
func New() *margopher {
	return &margopher{states: make(map[[2]string][]string)}
}

// Generate margopher sentence based on a given length
func (m *margopher) Generate() string {

	var sentence bytes.Buffer

	// Initialize prefix with a random key
	prefix := m.getRandomPrefix([2]string{"", ""})
	sentence.WriteString(strings.Join(prefix[:], " ") + " ")

	for {
		suffix := getRandomWord(m.states[prefix])

		// Break the loop if suffix ends in "." and sentenceLength is enough or if
		// the suffix is nothing, assume the end of the sentence
		if isTerminalWord(suffix) || (len(suffix) == 0) {
			break
		}

		sentence.WriteString(suffix + " ")
		prefix = [2]string{prefix[1], suffix}
	}

	if !strings.HasSuffix(sentence.String(), ".") {
		sentence.Truncate(len(sentence.String()) - 1)
		sentence.WriteString(".")
	}

	return sentence.String()
}

// Readers
func (m *margopher) ReadText(text string) string {
	m.ParseText(text)

	return m.Generate()
}

func (m *margopher) ReadFile(filePath string) string {
	m.ParseFile(filePath)

	return m.Generate()
}

func (m *margopher) ReadURL(URL string) string {
	m.ParseURL(URL)

	return m.Generate()
}

func (m *margopher) ReadDictionary() map[[2]string][]string {
	return m.states
}

// Return a random prefix other than the one in the arguments
func (m *margopher) getRandomPrefix(prefix [2]string) [2]string {
	// By default, Go orders keys randomly for maps
	for key := range m.states {
		if key != prefix {
			prefix = key
			break
		}
	}

	return prefix
}
