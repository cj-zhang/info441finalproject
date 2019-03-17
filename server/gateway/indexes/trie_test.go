package indexes

import (
	"strconv"
	"testing"
)

//TODO: implement automated tests for your trie data structure
func TestTrie(t *testing.T) {
	tester := NewTrie()
	tester.Add("hello", 1)
	tester.Add("hello", 1)
	tester.Add("hell", 2)
	tester.Add("hellp", 3)
	tester.Add("hello", 4)

	results := tester.ReturnPrefixMatches(1, "hell")
	for _, item := range results {
		t.Log("hell " + strconv.FormatInt(item, 10))
	}
	results = tester.ReturnPrefixMatches(2, "hello")
	for _, item := range results {
		t.Log("hello " + strconv.FormatInt(item, 10))
	}

	tester.Delete("hello", 1)
	results = tester.ReturnPrefixMatches(2, "hello")
	for _, item := range results {
		t.Log("hellod " + strconv.FormatInt(item, 10))
	}

	tester.Delete("hell", 2)
	results = tester.ReturnPrefixMatches(1, "hell")
	for _, item := range results {
		t.Log("helld " + strconv.FormatInt(item, 10))
	}

}
