package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// create a new type of 'deck'
// which is a slice of strings
// deck == []string; deck 是一個string 的 slice
type deck []string

func newDeck() deck { // 規定function 的 return 是 type deck
	cards := deck{}

	cardSuits := []string{"Spades", "Dimands", "Heart", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) { // (deck,deck) return 2個值
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",") // deck to string slice to one string with ,
}

func (d deck) saveToFile(filename string) error { // return type error
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		// option #1 - log the error and return a call to newDeck()
		// option #2 - Log the error and entirely quit the program
		fmt.Println("Error:", err)
		os.Exit(1) // quit the program
	}

	s := strings.Split(string(bs), ",")
	return deck(s) // slice string to dekc type
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)           // math/rand.Intn random number in [0,n), len(): lenght of slice
		d[i], d[newPosition] = d[newPosition], d[i] // d[i] 的值 跟 d[newPosition]的值 互換
	}
}
