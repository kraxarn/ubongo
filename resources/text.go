package resources

import (
	"embed"
	"fmt"
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
