package levenshteinsearch

import "testing"

func TestSearch(t *testing.T) {

	dict := CreateDictionary()

	dict.Put("banana")
	dict.Put("orange")
	dict.Put("monkey")
	dict.Put("бананы")
	dict.Put("оранжевый")
	dict.Put("обезьяна")

	result := dict.SearchAll("banana", 1)
	for word := range result {
		if word != "banana" {
			t.Error("Expected to find 'banana' with a distance of 1")
		}
	}

	result = dict.SearchAll("banan", 1)
	for word := range result {
		if word != "banana" {
			t.Error("Expected to find 'banan' with a distance of 1")
		}
	}

	result = dict.SearchAll("a", 5)
	if len(result) != 2 {
		t.Error("Expected to find 'banana' and 'orange' with a distance of 5")
	}

	result = dict.SearchAll("a", 6)
	if len(result) != 4 {
		t.Error("Expected to find 'banana', 'orange' and 'monkey' with a distance of 6")
	}

	result = dict.SearchAll("банан", 1)
	for word := range result {
		if word != "бананы" {
			t.Error("Expected to find 'банан' with a distance of 1")
		}
	}

	result = dict.SearchAll("ан", 2)
	if len(result) != 0 {
		t.Error("Expected to find 'ан' with a distance of 2")
	}

	result = dict.SearchAll("а", 6)
	if len(result) != 4 {
		t.Error("Expected to find 'а' with a distance of 2")
	}
}
