package margopher

import (
	"bytes"
	"strings"
)

type margopher struct {
	states map[[2]string][]string
}

// Margopher constructor
func New() *margopher {
	return &margopher{states: make(map[[2]string][]string)}
}

// Generate margopher senetence based on a given length
func (m *margopher) Generate() string {

	var sentence bytes.Buffer

	// Initialize prefix with a random key
	prefix := m.getRandomPrefix([2]string{"", ""})
	sentence.WriteString(strings.Join(prefix[:], " ") + " ")

	for {
		suffix := getRandomWord(m.states[prefix])
		sentence.WriteString(suffix + " ")

		// Break the loop if suffix ends in "." and senetenceLength is enough
		if isTerminalWord(suffix) {
			break
		}

		prefix = [2]string{prefix[1], suffix}
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
