package main

import (
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)


type deck []string

func (d deck) printCards() string{
	var returnString string
	for _, card := range d {
		//fmt.Printf("index %d has value %q", i, card)
		returnString = returnString +" "+ card
	}
	return returnString
}

func NewDeck() deck {
	cards := deck{}
	cardValue := []string{"ace","two", "three", "four"}
	cardSuits := []string{"space","Hearts", "Diamond", "clubs "}

	for _, cv := range cardValue {
		for _, cs := range cardSuits {
			cards = append(cards,cv+" of "+cs) 
		}
	}
	return cards
}

func (d deck) Deal(numberToDeal int) (d1,d2 deck){
	d1 = d[:numberToDeal]
	d2 = d[numberToDeal:]
	return d1,d2
}



 func (d deck) toString() string {
	deckToStringSlice := []string(d)
	deckToString := strings.Join(deckToStringSlice, ",")
	return deckToString
 }

 func (d deck)SaveToDrive(fileName string) error {
	deckString := d.toString()
	byteToSave := []byte(deckString)
	//defer byteToSave.close

	err := os.WriteFile(fileName,byteToSave,0666)
	return err
}

func readFromDrive(fileName string)deck{ 
	file, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	deckString := string(file)
	deckStringSlice := strings.Split(deckString,",")
return deck(deckStringSlice)
}

func (d deck) shuffle()  {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		newPosition :=	r.Intn(len(d)-1)
		d[i], d[newPosition] = d[newPosition], d[i] 
	}
}