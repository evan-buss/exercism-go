package tournament

import (
	"encoding/csv"
	"io"
	"log"
)

// Team contains each teams record and info
type Team struct {
	name   string
	played int
	wins   int
	draws  int
	losses int
	points int
}

// Game contains the parsed data from the string
type Game struct {
	teams   [2]string
	outcome [2]string
}

var tournament map[string]Team

// Tally takes the tournament results and formats the output
func Tally(reader io.Reader, writer io.Writer) error {
	data, err := parse(reader)
	if err != nil {
		return err
	}
	display(data)
	return nil
}

func display(games []Game) {
	for _, game := range games {
		val1, ok1 := tournament[game.teams[0]]
		val2, ok2 := tournament[game.teams[1]]

		// Create new if they don't exist
		if !ok1 {
			tournament[game.teams[0]] = Team{name: game.teams[0]}
		}
		if !ok2 {
			tournament[game.teams[0]] = Team{name: game.teams[0]}
		}
	}
}

// Parse reads from the input and determines the outcomes of each game
func parse(reader io.Reader) ([]Game, error) {
	var output []Game //Record of all games
	csv := csv.NewReader(reader)
	csv.Comma = ';'

	for {
		record, err := csv.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
			return nil, err
		}

		output = append(output,
			Game{
				[2]string{
					record[0],
					record[1]},
				determineOutcomes(record[2])})
	}
	return output, nil
}

// Parses inputs and returns the outcomes
func determineOutcomes(s string) [2]string {
	var output [2]string
	switch s {
	case "win":
		output[0] = "win"
		output[1] = "loss"
	case "loss":
		output[0] = "loss"
		output[1] = "win"
	case "draw":
		output[0] = "draw"
		output[1] = "draw"
	}

	return output
}
