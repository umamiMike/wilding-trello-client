package main

import (
	"fmt"
	"log"
	"os"

	"github.com/adlio/trello"
)

type CardsColl struct {
	cards []*trello.Card
}

func main() {

	appKey := os.Getenv("TRELLO_APPKEY")
	token := os.Getenv("TRELLO_TOKEN")
	boardID := os.Args[1]

	client := trello.NewClient(appKey, token)

	board, err := client.GetBoard(boardID, trello.Defaults())
	if err != nil {
		log.Fatal("getBoard failed")
	}

	cards, err := board.GetCards(trello.Defaults())
	if err != nil {
		log.Fatal("get card failed ")

	}

	CardsColl{cards: cards}.Descriptions()

}

func (c CardsColl) Descriptions() {
	for _, v := range c.cards {
		fmt.Println("")
		fmt.Println("# Card - ", v.Name)
		fmt.Println(v.Desc)
	}
}
