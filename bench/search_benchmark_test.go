package bench

import (
	"bufio"
	"log"
	"os"
	"strings"
	"testing"

	levenshtein "github.com/covrom/levenshteinsearch/pkg/levenshteinsearch"
)

var aliceWords []string

const maxSimilaritySearch = 10

func ensureAlice() error {

	if len(aliceWords) > 0 {
		return nil
	}

	file, err := os.Open("../assets/alice/alice.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		nextWord := strings.TrimSpace(scanner.Text())
		nextWord = strings.Replace(nextWord, `"`, "", -1)
		nextWord = strings.Replace(nextWord, `.`, "", -1)
		nextWord = strings.Replace(nextWord, `,`, "", -1)
		nextWord = strings.Replace(nextWord, `;`, "", -1)
		nextWord = strings.Replace(nextWord, `:`, "", -1)
		nextWord = strings.Replace(nextWord, `!`, "", -1)
		nextWord = strings.Replace(nextWord, `?`, "", -1)
		for _, word := range strings.Split(nextWord, " ") {
			word = strings.TrimSpace(word)
			word = strings.ToLower(word)
			aliceWords = append(aliceWords, word)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return nil
}

func BenchmarkOptimized1Word(b *testing.B) {

	if err := ensureAlice(); err != nil {
		log.Fatal(err)
	}

	dict := levenshtein.CreateDictionary()

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		for i := 0; i < maxSimilaritySearch; i++ {
			dict.SearchAll("rabbit", i)
		}
	}
}

func BenchmarkOptimized3Word(b *testing.B) {

	if err := ensureAlice(); err != nil {
		log.Fatal(err)
	}

	dict := levenshtein.CreateDictionary()

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		for i := 0; i < maxSimilaritySearch; i++ {
			dict.SearchAll("rabbit", i)
		}
	}
	for i := 0; i < b.N; i++ {
		for i := 0; i < maxSimilaritySearch; i++ {
			dict.SearchAll("eart", i)
		}
	}
	for i := 0; i < b.N; i++ {
		for i := 0; i < maxSimilaritySearch; i++ {
			dict.SearchAll("the", i)
		}
	}
}
