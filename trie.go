package trie

import (
	"container/list"
	"fmt"
)

type Trie struct {
	value    []interface{}
	children map[string]*Trie
}

func NewTrie() *Trie {
	t := new(Trie)
	t.value = nil
	t.children = make(map[string]*Trie)
	return t
}

func (t *Trie) Add(key string, value interface{}) {

	if len(key) == 0 {
		t.addValue(value)
		return
	}

	// get the first character of our string
	//r, size := utf8.DecodeRuneInString(key)
	c := string(key[0])

	// if the first char of our string is in the map
	next, ok := t.children[c]

	// if we didn't have a child for that key then create a new node
	if !ok {
		next = NewTrie()
		t.children[c] = next
	}

	next.Add(key[1:], value)
}

func (t *Trie) addValue(value interface{}) {
	if t.value == nil {
		t.value = []interface{}{value}
	} else {
		t.value = append(t.value, value)
	}
}

func (t *Trie) Find(key string) (interface{}, error) {
	// get the first character of our string
	//r, size := utf8.DecodeRuneInString(key)
	c := string(key[0])

	next, ok := t.children[c]
	if !ok {
		return nil, fmt.Errorf("Key not found %q", c)
	}

	// if this is the last char then value is next.value
	if len(key) == 1 {
		// somehow we're missing the terminating value node
		if len(t.children) == 0 {
			return nil, fmt.Errorf("Key not found")
		} else {
			return next.value, nil
		}
	}

	return next.Find(key[1:])
}

func (t *Trie) MatchPartial(key string) ([]interface{}, error) {

	if len(key) == 0 {
		return t.fetchRemainder(), nil
	}

	c := string(key[0])

	next, ok := t.children[c]
	if !ok {
		return nil, fmt.Errorf("Key not found %q", c)
	}

	return next.MatchPartial(key[1:])
}

func (t *Trie) fetchRemainder() []interface{} {

	results := make([]interface{}, 0, 1)

	queue := list.New()
	queue.PushBack(t)

	var element *list.Element
	var node *Trie

	for queue.Len() > 0 {

		element = queue.Front()
		queue.Remove(element)
		node = element.Value.(*Trie)

		if node.value != nil {

			// save a cast of this for convenience
			// var value []interface{} = node.value.([]interface{})
			value := node.value

			// if results has insufficent capacity then grow it
			if cap(value) < len(value)+len(results) {
				newSlice := make([]interface{}, len(results), len(value)+len(results))
				copy(newSlice, results)
				results = newSlice
			}
			// extend the results slice with our value
			results = append(results, value...)
		}

		for _, child := range node.children {
			queue.PushBack(child)
		}

	}

	return results
}
