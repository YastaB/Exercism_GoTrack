package tournament

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
)

type Team struct {
	name string
	MP   int // Match Played
	W    int // Win
	D    int // Draw
	L    int // Loss
	P    int // Points
}

func applyMatchCondition(matchScript string, teams map[string]*Team) error {
	fields := strings.Split(matchScript, ";")
	if len(fields) != 3 {
		return errors.New("invalid match script")
	}

	team1 := fields[0]
	team2 := fields[1]
	status := fields[2]

	if _, ok := teams[team1]; !ok {
		teams[team1] = &Team{name: team1}
	}
	if _, ok := teams[team2]; !ok {
		teams[team2] = &Team{name: team2}
	}

	switch status {
	case "win":
		teams[team1].W++
		teams[team1].MP++
		teams[team1].P += 3
		teams[team2].L++
		teams[team2].MP++
		break
	case "loss":
		teams[team1].L++
		teams[team1].MP++
		teams[team2].W++
		teams[team2].MP++
		teams[team2].P += 3
		break
	case "draw":
		teams[team1].P++
		teams[team1].MP++
		teams[team1].D++
		teams[team2].P++
		teams[team2].MP++
		teams[team2].D++
		break
	default:
		return errors.New("invalid match condition")
	}
	return nil
}

// Tally gets match scripts in text and outputs to the writer
// in the order of descending total points
func Tally(reader io.Reader, writer io.Writer) error {
	var teams = map[string]*Team{}
	ioReader := bufio.NewReader(reader)

	for {
		line, _ := ioReader.ReadString('\n')
		if len(line) == 0{
			break
		}
		if line == "\n" || line[0] == '#'{
			continue
		}
		// remove the end line break
		line = line[:len(line)-1]
		parseError := applyMatchCondition(line, teams)
		if parseError != nil{
			return parseError
		}
	}

	teamList := make([]*Team, 0, len(teams))
	for _, val := range teams{
		teamList = append(teamList, val)
	}

	sort.Slice(teamList, func(i, j int) bool {
		if teamList[i].P > teamList[j].P {
			return true
		}else if (teamList[i].P == teamList[j].P) && teamList[i].name < teamList[j].name{
			return true
		}
		return false
	})
	writer.Write([]byte(fmt.Sprintf("%-31v|%3v |%3v |%3v |%3v |%3v\n", "Team", "MP", "W", "D", "L","P")))
	for _, team := range teamList{
		writer.Write([]byte(fmt.Sprintf("%-31v|%3d |%3d |%3d |%3d |%3d\n", team.name, team.MP, team.W, team.D, team.L, team.P)))
	}
	return nil
}
