// Code generated by jtd-codegen for Go v0.2.1. DO NOT EDIT.

package qg

import (
	"encoding/json"
	"fmt"
)

type Qg = interface{}

type Command struct {
	// Value can be the following types:
	//  - [CommandBeginGame] (BeginGame)
	//  - [CommandEndGame] (EndGame)
	//  - [CommandJeopardyChooseQuestion] (JeopardyChooseQuestion)
	//  - [CommandJeopardyPlayerJudgment] (JeopardyPlayerJudgment)
	//  - [CommandJeopardyPressButton] (JeopardyPressButton)
	//  - [CommandJoinGame] (JoinGame)
	Value valueCommand

	t string
}

func (v Command) MarshalJSON() ([]byte, error) {
	switch value := v.Value.(type) {
	case CommandBeginGame:
		return json.Marshal(struct {
			T string `json:"type"`
			CommandBeginGame
		}{"BeginGame", value})
	case CommandEndGame:
		return json.Marshal(struct {
			T string `json:"type"`
			CommandEndGame
		}{"EndGame", value})
	case CommandJeopardyChooseQuestion:
		return json.Marshal(struct {
			T string `json:"type"`
			CommandJeopardyChooseQuestion
		}{"JeopardyChooseQuestion", value})
	case CommandJeopardyPlayerJudgment:
		return json.Marshal(struct {
			T string `json:"type"`
			CommandJeopardyPlayerJudgment
		}{"JeopardyPlayerJudgment", value})
	case CommandJeopardyPressButton:
		return json.Marshal(struct {
			T string `json:"type"`
			CommandJeopardyPressButton
		}{"JeopardyPressButton", value})
	case CommandJoinGame:
		return json.Marshal(struct {
			T string `json:"type"`
			CommandJoinGame
		}{"JoinGame", value})
	default:
		panic("unreachable")
	}
}

func (v *Command) UnmarshalJSON(b []byte) error {
	var t struct {
		T string `json:"type"`
	}
	if err := json.Unmarshal(b, &t); err != nil {
		return err
	}

	var value valueCommand
	var err error

	switch t.T {
	case "BeginGame":
		var v CommandBeginGame
		err = json.Unmarshal(b, &v)
		value = v
	case "EndGame":
		var v CommandEndGame
		err = json.Unmarshal(b, &v)
		value = v
	case "JeopardyChooseQuestion":
		var v CommandJeopardyChooseQuestion
		err = json.Unmarshal(b, &v)
		value = v
	case "JeopardyPlayerJudgment":
		var v CommandJeopardyPlayerJudgment
		err = json.Unmarshal(b, &v)
		value = v
	case "JeopardyPressButton":
		var v CommandJeopardyPressButton
		err = json.Unmarshal(b, &v)
		value = v
	case "JoinGame":
		var v CommandJoinGame
		err = json.Unmarshal(b, &v)
		value = v
	default:
		err = fmt.Errorf("bad Type value: %s", t.T)
	}

	if err != nil {
		return err
	}

	v.t = t.T
	v.Value = value
	return nil
}

type valueCommand interface {
	isCommand()
}

func (CommandBeginGame) isCommand()              {}
func (CommandEndGame) isCommand()                {}
func (CommandJeopardyChooseQuestion) isCommand() {}
func (CommandJeopardyPlayerJudgment) isCommand() {}
func (CommandJeopardyPressButton) isCommand()    {}
func (CommandJoinGame) isCommand()               {}

// CommandBeginGame is sent by a client to begin a game.
type CommandBeginGame struct {
}

// CommandEndGame is sent by a client to end the current game. The server
// will respond with an EventGameEnded. Only game admins (including the
// host) can end the game.
type CommandEndGame struct {
	// declareWinner determines whether the game should be ended with a
	// winner or not. If true, the game will be ended with a winner. If
	// false, the game will be ended abruptly.
	DeclareWinner bool `json:"declareWinner"`
}

// CommandJeopardyChooseQuestion is sent by a player to choose a question.
// The server must do validation to ensure that the player is allowed to
// choose the question.
type CommandJeopardyChooseQuestion struct {
	Category int32 `json:"category"`
	Question int32 `json:"question"`
}

// CommandJeopardyPlayerJudgment is emitted by a game admin to indicate
// whether a player has answered a question correctly. The winning player is
// whoever the last EventJeopardyButtonPressed event indicated. That player
// will instantly receive the points for the question, and the game will let
// them choose the next category and question. If the player answered wrong,
// then the game will let others press the button.
type CommandJeopardyPlayerJudgment struct {
	Correct bool `json:"correct"`
}

// CommandJeopardyPressButton is emitted when a player presses the button
// during a question. It is only valid to emit this command when the game is
// in the question state.
type CommandJeopardyPressButton struct {
}

// CommandJoinGame is sent by a client to join a game. The client (or the
// user) supplies a game ID and a player name. The server will respond with
// an EventJoinedGame.
type CommandJoinGame struct {
	// adminPassword is the password of the admin of the game.
	AdminPassword *string `json:"adminPassword"`
	// gameID is the ID of the game to join.
	GameID GameID `json:"gameID"`
	// playerName is the wanted name of the user.
	PlayerName PlayerName `json:"playerName"`
}

// Error is returned on every API error.
type Error struct {
	// Message is the error message
	Message string `json:"message"`
}

type Event struct {
	// Value can be the following types:
	//  - [EventGameEnded] (GameEnded)
	//  - [EventJeopardyBeginQuestion] (JeopardyBeginQuestion)
	//  - [EventJeopardyButtonPressed] (JeopardyButtonPressed)
	//  - [EventJeopardyResumeButton] (JeopardyResumeButton)
	//  - [EventJeopardyTurnEnded] (JeopardyTurnEnded)
	//  - [EventJoinedGame] (JoinedGame)
	//  - [EventPlayerJoined] (PlayerJoined)
	Value valueEvent

	t string
}

func (v Event) MarshalJSON() ([]byte, error) {
	switch value := v.Value.(type) {
	case EventGameEnded:
		return json.Marshal(struct {
			T string `json:"type"`
			EventGameEnded
		}{"GameEnded", value})
	case EventJeopardyBeginQuestion:
		return json.Marshal(struct {
			T string `json:"type"`
			EventJeopardyBeginQuestion
		}{"JeopardyBeginQuestion", value})
	case EventJeopardyButtonPressed:
		return json.Marshal(struct {
			T string `json:"type"`
			EventJeopardyButtonPressed
		}{"JeopardyButtonPressed", value})
	case EventJeopardyResumeButton:
		return json.Marshal(struct {
			T string `json:"type"`
			EventJeopardyResumeButton
		}{"JeopardyResumeButton", value})
	case EventJeopardyTurnEnded:
		return json.Marshal(struct {
			T string `json:"type"`
			EventJeopardyTurnEnded
		}{"JeopardyTurnEnded", value})
	case EventJoinedGame:
		return json.Marshal(struct {
			T string `json:"type"`
			EventJoinedGame
		}{"JoinedGame", value})
	case EventPlayerJoined:
		return json.Marshal(struct {
			T string `json:"type"`
			EventPlayerJoined
		}{"PlayerJoined", value})
	default:
		panic("unreachable")
	}
}

func (v *Event) UnmarshalJSON(b []byte) error {
	var t struct {
		T string `json:"type"`
	}
	if err := json.Unmarshal(b, &t); err != nil {
		return err
	}

	var value valueEvent
	var err error

	switch t.T {
	case "GameEnded":
		var v EventGameEnded
		err = json.Unmarshal(b, &v)
		value = v
	case "JeopardyBeginQuestion":
		var v EventJeopardyBeginQuestion
		err = json.Unmarshal(b, &v)
		value = v
	case "JeopardyButtonPressed":
		var v EventJeopardyButtonPressed
		err = json.Unmarshal(b, &v)
		value = v
	case "JeopardyResumeButton":
		var v EventJeopardyResumeButton
		err = json.Unmarshal(b, &v)
		value = v
	case "JeopardyTurnEnded":
		var v EventJeopardyTurnEnded
		err = json.Unmarshal(b, &v)
		value = v
	case "JoinedGame":
		var v EventJoinedGame
		err = json.Unmarshal(b, &v)
		value = v
	case "PlayerJoined":
		var v EventPlayerJoined
		err = json.Unmarshal(b, &v)
		value = v
	default:
		err = fmt.Errorf("bad Type value: %s", t.T)
	}

	if err != nil {
		return err
	}

	v.t = t.T
	v.Value = value
	return nil
}

type valueEvent interface {
	isEvent()
}

func (EventGameEnded) isEvent()             {}
func (EventJeopardyBeginQuestion) isEvent() {}
func (EventJeopardyButtonPressed) isEvent() {}
func (EventJeopardyResumeButton) isEvent()  {}
func (EventJeopardyTurnEnded) isEvent()     {}
func (EventJoinedGame) isEvent()            {}
func (EventPlayerJoined) isEvent()          {}

// EventGameEnded is emitted when the current game ends.
type EventGameEnded struct {
	Leaderboard Leaderboard `json:"leaderboard"`
}

// EventJeopardyBeginQuestion is emitted when a question begins within this
// Jeopardy game. It is usually emitted once the chooser player has chosen a
// category and value.
//
// Each category name and question value will map to a category and question
// within the game data. Note that a question may repeat across multiple
// categories.
type EventJeopardyBeginQuestion struct {
	Category int32      `json:"category"`
	Chooser  PlayerName `json:"chooser"`
	Points   float32    `json:"points"`
	Question int32      `json:"question"`
}

// EventJeopardyButtonPressed is emitted when any player had pressed a button
// on their device, voiding other players' buttons. This event is only
// emitted when the game is in the "question" state.
type EventJeopardyButtonPressed struct {
	PlayerName PlayerName `json:"playerName"`
}

// EventJeopardyResumeButton is emitted when the player can now continue to
// press the button whenever they are ready to answer the question. This
// could happen if the other player who pressed the button first got the
// question wrong.
//
// Note that if alreadyPressed is true, then the player has already pressed
// the button, so they cannot press it again.
type EventJeopardyResumeButton struct {
	AlreadyAnsweredPlayers []PlayerName `json:"alreadyAnsweredPlayers"`
}

// EventJeopardyTurnEnded is emitted when a turn ends or when the game first
// starts.
type EventJeopardyTurnEnded struct {
	Chooser     PlayerName  `json:"chooser"`
	Leaderboard Leaderboard `json:"leaderboard"`
}

// EventJoinedGame is emitted when the current player joins a game. It is a
// reply to CommandJoinGame and is only for the current player. Not to be
// confused with EventPlayerJoinedGame, which is emitted when any player
// joins the current game.
type EventJoinedGame struct {
	GameData *GameData `json:"gameData"`
	GameInfo GameInfo  `json:"gameInfo"`
	IsAdmin  bool      `json:"isAdmin"`
}

// EventPlayerJoined is emitted when a player joins the current game.
type EventPlayerJoined struct {
	PlayerName PlayerName `json:"playerName"`
}

// GameData is the game data. It contains all the information about the game.
type GameData struct {
	// Value can be the following types:
	//  - [GameDataJeopardy] (jeopardy)
	//  - [GameDataKahoot] (kahoot)
	Value valueGameData

	t string
}

func (v GameData) MarshalJSON() ([]byte, error) {
	switch value := v.Value.(type) {
	case GameDataJeopardy:
		return json.Marshal(struct {
			T string `json:"game"`
			GameDataJeopardy
		}{"jeopardy", value})
	case GameDataKahoot:
		return json.Marshal(struct {
			T string `json:"game"`
			GameDataKahoot
		}{"kahoot", value})
	default:
		panic("unreachable")
	}
}

func (v *GameData) UnmarshalJSON(b []byte) error {
	var t struct {
		T string `json:"game"`
	}
	if err := json.Unmarshal(b, &t); err != nil {
		return err
	}

	var value valueGameData
	var err error

	switch t.T {
	case "jeopardy":
		var v GameDataJeopardy
		err = json.Unmarshal(b, &v)
		value = v
	case "kahoot":
		var v GameDataKahoot
		err = json.Unmarshal(b, &v)
		value = v
	default:
		err = fmt.Errorf("bad Game value: %s", t.T)
	}

	if err != nil {
		return err
	}

	v.t = t.T
	v.Value = value
	return nil
}

type valueGameData interface {
	isGameData()
}

func (GameDataJeopardy) isGameData() {}
func (GameDataKahoot) isGameData()   {}

type GameDataJeopardy struct {
	Data JeopardyGameData `json:"data"`
}

type GameDataKahoot struct {
	Data KahootGameData `json:"data"`
}

// GameID is the unique identifier for a game. Each player must type this
// code to join the game.
type GameID = string

type GameInfo struct {
	// Value can be the following types:
	//  - [GameInfoJeopardy] (jeopardy)
	Value valueGameInfo

	t string
}

func (v GameInfo) MarshalJSON() ([]byte, error) {
	switch value := v.Value.(type) {
	case GameInfoJeopardy:
		return json.Marshal(struct {
			T string `json:"type"`
			GameInfoJeopardy
		}{"jeopardy", value})
	default:
		panic("unreachable")
	}
}

func (v *GameInfo) UnmarshalJSON(b []byte) error {
	var t struct {
		T string `json:"type"`
	}
	if err := json.Unmarshal(b, &t); err != nil {
		return err
	}

	var value valueGameInfo
	var err error

	switch t.T {
	case "jeopardy":
		var v GameInfoJeopardy
		err = json.Unmarshal(b, &v)
		value = v
	default:
		err = fmt.Errorf("bad Type value: %s", t.T)
	}

	if err != nil {
		return err
	}

	v.t = t.T
	v.Value = value
	return nil
}

type valueGameInfo interface {
	isGameInfo()
}

func (GameInfoJeopardy) isGameInfo() {}

type GameInfoJeopardy struct {
	Data JeopardyGameInfo `json:"data"`
}

type GameType string

const (
	GameTypeJeopardy GameType = "jeopardy"
	GameTypeKahoot   GameType = "kahoot"
)

// JeopardyCategory is a category in a Jeopardy game.
type JeopardyCategory struct {
	// name is the name of the category.
	Name string `json:"name"`
	// questions are the questions in the category.
	Questions []JeopardyQuestion `json:"questions"`
}

// JeopardyGameData is the game data for a Jeopardy game.
type JeopardyGameData struct {
	Categories []JeopardyCategory `json:"categories"`
	// score_multiplier is the score multiplier for each question. The
	// default is 100.
	ScoreMultiplier *float32 `json:"score_multiplier,omitempty"`
}

// JeopardyGameInfo is the initial information for a Jeopardy game. This type
// contains no useful information about the entire game data, so it's used to
// send to players the first time they join.
type JeopardyGameInfo struct {
	Categories      []string `json:"categories"`
	NumQuestions    int32    `json:"numQuestions"`
	ScoreMultiplier float32  `json:"scoreMultiplier"`
}

// JeopardyQuestion is a question in a Jeopardy game.
type JeopardyQuestion struct {
	// question is the question.
	Question string `json:"question"`
}

// KahootGameData is the game data for a Kahoot game.
type KahootGameData struct {
	// questions are the questions in the game.
	Questions []KahootQuestion `json:"questions"`
	// time_limit is the time limit for each question. The format is in
	// Go's time.Duration, e.g. 10s for 10 seconds.
	TimeLimit string `json:"time_limit"`
}

// KahootQuestion is a question in a Kahoot game.
type KahootQuestion struct {
	// answers are the possible answers.
	Answers []string `json:"answers"`
	// question is the question.
	Question string `json:"question"`
}

// Leaderboard is a list of players and their scores.
type Leaderboard = []LeaderboardEntry

type LeaderboardEntry struct {
	PlayerName string  `json:"playerName"`
	Score      float32 `json:"score"`
}

// PlayerName is the name of a player.
type PlayerName = string

type RequestGetGame struct {
	GameID string `json:"gameID"`
}

type RequestGetJeopardyGame struct {
	GameID string `json:"gameID"`
}

type RequestNewGame struct {
	AdminPassword string   `json:"admin_password"`
	Data          GameData `json:"data"`
}

type ResponseGetGame struct {
	GameType GameType `json:"gameType"`
}

type ResponseGetJeopardyGame struct {
	Info JeopardyGameInfo `json:"info"`
}

type ResponseNewGame struct {
	GameID   string   `json:"gameID"`
	GameType GameType `json:"gameType"`
}
