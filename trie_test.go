package trie

import "testing"

func TestAddChar(t *testing.T) {

	trie := NewTrie()
	trie.Add("a", 1)

	if len(trie.children) != 1 {
		t.Errorf("Trie has %s children, expected 1", len(trie.children))
	}

	if _, ok := trie.children["a"]; !ok {
		t.Error("Child 'a' not found in trie")
	}

	if trie.children["a"].value != 1 {
		t.Error("Child")
	}

}

func TestAddWord(t *testing.T) {

	trie := NewTrie()
	trie.Add("was", "here")

	child, ok := trie.children["w"]
	if !ok {
		t.Error("Child 'w' not found at trie root")
	}

	if child.value != nil {
		t.Errorf("Value for child should be nil, found %v", child.value)
	}

	child, ok = child.children["a"]
	if !ok {
		t.Error("Child 'a' not found at trie 1st child")
	}

	if child.value != nil {
		t.Errorf("Value for child should be nil, found %v", child.value)
	}

	child, ok = child.children["s"]
	if !ok {
		t.Error("Child 's' not found at trie 2nd child")
	}

	if child.value != "here" {
		t.Errorf("Incorrect value found for key 'was' -> %s\n", child.value)
	}

}

func TestFind(t *testing.T) {

	trie := NewTrie()
	trie.Add("golang", 1)

	val, err := trie.Find("python")
	if err == nil {
		t.Errorf("Expected an error, received none. val -> %v\n", val)
	}
	if val != nil {
		t.Errorf("Expected nil value, received %v", val)
	}

	val, err = trie.Find("golang")
	if err != nil {
		t.Errorf("Expected no error, received %s", err)
	}
	if val != 1 {
		t.Errorf("Expected value 1 for key 'golang', found %v", val)
	}

}

func TestMatchPartial(t *testing.T) {
	trie := NewTrie()
	trie.Add("golang", 1)
	trie.Add("python", 0)

	results, _ := trie.MatchPartial("golang")

	if len(results) != 1 {
		t.Errorf("Unexpected results len, expceted 1 got %d", len(results))
	}

}
