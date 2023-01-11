package main


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
			cards = append(cards,cv+" 0f "+cs) 
		}
	}
	return cards
}

func (d deck) Deal(numberToDeal int) (d1,d2 deck){
	d1 = d[:numberToDeal]
	d2 = d[numberToDeal:]
	return d1,d2
}
