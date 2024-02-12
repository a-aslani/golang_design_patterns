package flyweight

import (
	"errors"
	"time"
)

const (
	TeamA = iota
	TeamB
)

type Player struct {
	Name         string
	Surname      string
	PreviousTeam uint64
	Photo        []byte
}

type Match struct {
	Date          time.Time
	VisitorID     uint64
	LocalID       uint64
	LocalScore    byte
	VisitorScore  byte
	LocalShoots   uint16
	VisitorShoots uint16
}

type HistoricalData struct {
	Year          uint8
	LeagueResults []Match
}

type Team struct {
	ID             uint64
	Name           string
	Shield         []byte
	Players        []Player
	HistoricalData []HistoricalData
}

func getTeamFactory(team int) (Team, error) {
	switch team {
	case TeamA:
		return Team{
			ID:   1,
			Name: "TEAM_A",
		}, nil
	case TeamB:
		return Team{
			ID:   2,
			Name: "TEAM_B",
		}, nil
	default:
		return Team{}, errors.New("team not found")
	}
}

type teamFlyweightFactory struct {
	createdTeams map[int]*Team
}

func NewTeamFactory() teamFlyweightFactory {
	return teamFlyweightFactory{
		createdTeams: make(map[int]*Team),
	}
}

func (t *teamFlyweightFactory) GetTeam(teamName int) (*Team, error) {

	tm, ok := t.createdTeams[teamName]
	if ok {
		return tm, nil
	}

	team, err := getTeamFactory(teamName)
	if err != nil {
		return nil, err
	}
	t.createdTeams[teamName] = &team

	return t.createdTeams[teamName], nil
}

func (t *teamFlyweightFactory) GetNumberOfObjects() int {
	return len(t.createdTeams)
}
