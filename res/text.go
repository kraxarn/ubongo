package res

import (
	"embed"
	"fmt"
	"math/rand"
	"strings"
)

//go:embed text/english-*.txt
var files embed.FS

func read(name string) []string {
	data, err := files.ReadFile(fmt.Sprintf("text/english-%s.txt", name))
	if err != nil {
		return []string{}
	}
	return strings.Split(string(data), "\n")
}

func readLine(name string, index int) string {
	words := read(name)
	if len(words) == 0 {
		return ""
	}
	return words[index%len(words)]
}

func Adjective(index int) string {
	return readLine("adjectives", index)
}

func Noun(index int) string {
	return readLine("nouns", index)
}

func RandomWord(seed int64) string {
	rand.Seed(seed)
	adjective := Adjective(rand.Int())
	noun := Noun(rand.Int())
	return fmt.Sprintf("%s %s", adjective, noun)
}
