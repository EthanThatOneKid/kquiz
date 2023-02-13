// Code generated by jtd-codegen for Go v0.2.1. DO NOT EDIT.

package qg

import (
	"encoding/json"
	"fmt"
)

type Qg = interface{}

type Command struct {
	Type string

	JeopardyChooseQuestion CommandJeopardyChooseQuestion0

	JeopardyPlayerIsCorrect CommandJeopardyPlayerIsCorrect0

	JeopardyPressButton CommandJeopardyPressButton0

	JoinGame CommandJoinGame0
}

func (v Command) MarshalJSON() ([]byte, error) {
	switch v.Type {
	case "JeopardyChooseQuestion":
		return json.Marshal(struct { T string `json:"type"`; CommandJeopardyChooseQuestion0 }{ v.Type, v.JeopardyChooseQuestion })
	case "JeopardyPlayerIsCorrect":
		return json.Marshal(struct { T string `json:"type"`; CommandJeopardyPlayerIsCorrect0 }{ v.Type, v.JeopardyPlayerIsCorrect })
	case "JeopardyPressButton":
		return json.Marshal(struct { T string `json:"type"`; CommandJeopardyPressButton0 }{ v.Type, v.JeopardyPressButton })
	case "JoinGame":
		return json.Marshal(struct { T string `json:"type"`; CommandJoinGame0 }{ v.Type, v.JoinGame })
	}

	return nil, fmt.Errorf("bad Type value: %s", v.Type)
}

func (v *Command) UnmarshalJSON(b []byte) error {
	var t struct { T string `json:"type"` }
	if err := json.Unmarshal(b, &t); err != nil {
		return err
	}

	var err error
	switch t.T {
	case "JeopardyChooseQuestion":
		err = json.Unmarshal(b, &v.JeopardyChooseQuestion)
	case "JeopardyPlayerIsCorrect":
		err = json.Unmarshal(b, &v.JeopardyPlayerIsCorrect)
	case "JeopardyPressButton":
		err = json.Unmarshal(b, &v.JeopardyPressButton)
	case "JoinGame":
		err = json.Unmarshal(b, &v.JoinGame)
	default:
		err = fmt.Errorf("bad Type value: %s", t.T)
	}

	if err != nil {
		return err
	}

	v.Type = t.T
	return nil
}

type CommandJeopardyChooseQuestion0 struct {
	Data CommandJeopardyChooseQuestion `json:"data"`
}

type CommandJeopardyPlayerIsCorrect0 struct {
	Data CommandJeopardyPlayerIsCorrect `json:"data"`
}

type CommandJeopardyPressButton0 struct {
	Data CommandJeopardyPressButton `json:"data"`
}

type CommandJoinGame0 struct {
	Data CommandJoinGame `json:"data"`
}

// CommandJeopardyChooseQuestion is sent by a player to choose a question.
// The server must do validation to ensure that the player is allowed to
// choose the question.
type CommandJeopardyChooseQuestion struct {
	Category string `json:"category"`

	Question string `json:"question"`
}

// CommandJeopardyPlayerIsCorrect is emitted by a game moderator to indicate
// that a player has answered a question correctly. The winning player is
// whoever the last EventJeopardyButtonPressed event indicated. That player
// will instantly receive the points for the question, and the game will let
// them choose the next category and question.
type CommandJeopardyPlayerIsCorrect = interface{}

// CommandJeopardyPressButton is emitted when a player presses the button
// during a question. It is only valid to emit this command when the game is
// in the question state.
type CommandJeopardyPressButton = interface{}

// CommandJoinGame is sent by a client to join a game. The client (or the
// user) supplies a game ID and a player name. The server will respond with
// an EventJoinedGame.
type CommandJoinGame struct {
	// gameID is the ID of the game to join.
	GameID string `json:"gameID"`

	// playerName is the wanted name of the user.
	PlayerName PlayerName `json:"playerName"`
}

// Error is returned on every API error.
type Error struct {
	// Message is the error message
	Message string `json:"message"`
}

type Event struct {
	Type string

	GameEnded EventGameEnded0

	JeopardyBeginQuestion EventJeopardyBeginQuestion0

	JeopardyButtonPressed EventJeopardyButtonPressed0

	JeopardyResumeButton EventJeopardyResumeButton0

	JeopardyTurnEnded EventJeopardyTurnEnded0

	JoinedGame EventJoinedGame0

	PlayerJoined EventPlayerJoined0
}

func (v Event) MarshalJSON() ([]byte, error) {
	switch v.Type {
	case "GameEnded":
		return json.Marshal(struct { T string `json:"type"`; EventGameEnded0 }{ v.Type, v.GameEnded })
	case "JeopardyBeginQuestion":
		return json.Marshal(struct { T string `json:"type"`; EventJeopardyBeginQuestion0 }{ v.Type, v.JeopardyBeginQuestion })
	case "JeopardyButtonPressed":
		return json.Marshal(struct { T string `json:"type"`; EventJeopardyButtonPressed0 }{ v.Type, v.JeopardyButtonPressed })
	case "JeopardyResumeButton":
		return json.Marshal(struct { T string `json:"type"`; EventJeopardyResumeButton0 }{ v.Type, v.JeopardyResumeButton })
	case "JeopardyTurnEnded":
		return json.Marshal(struct { T string `json:"type"`; EventJeopardyTurnEnded0 }{ v.Type, v.JeopardyTurnEnded })
	case "JoinedGame":
		return json.Marshal(struct { T string `json:"type"`; EventJoinedGame0 }{ v.Type, v.JoinedGame })
	case "PlayerJoined":
		return json.Marshal(struct { T string `json:"type"`; EventPlayerJoined0 }{ v.Type, v.PlayerJoined })
	}

	return nil, fmt.Errorf("bad Type value: %s", v.Type)
}

func (v *Event) UnmarshalJSON(b []byte) error {
	var t struct { T string `json:"type"` }
	if err := json.Unmarshal(b, &t); err != nil {
		return err
	}

	var err error
	switch t.T {
	case "GameEnded":
		err = json.Unmarshal(b, &v.GameEnded)
	case "JeopardyBeginQuestion":
		err = json.Unmarshal(b, &v.JeopardyBeginQuestion)
	case "JeopardyButtonPressed":
		err = json.Unmarshal(b, &v.JeopardyButtonPressed)
	case "JeopardyResumeButton":
		err = json.Unmarshal(b, &v.JeopardyResumeButton)
	case "JeopardyTurnEnded":
		err = json.Unmarshal(b, &v.JeopardyTurnEnded)
	case "JoinedGame":
		err = json.Unmarshal(b, &v.JoinedGame)
	case "PlayerJoined":
		err = json.Unmarshal(b, &v.PlayerJoined)
	default:
		err = fmt.Errorf("bad Type value: %s", t.T)
	}

	if err != nil {
		return err
	}

	v.Type = t.T
	return nil
}

type EventGameEnded0 struct {
	Data EventGameEnded `json:"data"`
}

type EventJeopardyBeginQuestion0 struct {
	Data EventJeopardyBeginQuestion `json:"data"`
}

type EventJeopardyButtonPressed0 struct {
	Data EventJeopardyButtonPressed `json:"data"`
}

type EventJeopardyResumeButton0 struct {
	Data EventJeopardyResumeButton `json:"data"`
}

type EventJeopardyTurnEnded0 struct {
	Data EventJeopardyTurnEnded `json:"data"`
}

type EventJoinedGame0 struct {
	Data EventJoinedGame `json:"data"`
}

type EventPlayerJoined0 struct {
	Data EventPlayerJoined `json:"data"`
}

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
	Category string `json:"category"`

	Chooser PlayerName `json:"chooser"`

	Question string `json:"question"`
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
	AlreadyPressed bool `json:"alreadyPressed"`
}

// EventJeopardyTurnEnded is emitted when a turn ends or when the game first
// starts.
type EventJeopardyTurnEnded struct {
	CurrentScore float64 `json:"currentScore"`

	IsChooser bool `json:"isChooser"`

	Leaderboard Leaderboard `json:"leaderboard"`
}

// EventJoinedGame is emitted when the current player joins a game. It is a
// reply to CommandJoinGame and is only for the current player. Not to be
// confused with EventPlayerJoinedGame, which is emitted when any player
// joins the current game.
type EventJoinedGame = Game

// EventPlayerJoined is emitted when a player joins the current game.
type EventPlayerJoined struct {
	PlayerName PlayerName `json:"playerName"`
}

// Game is the main game object. It contains all the information about the
// game.
type Game struct {
	Game string

	Jeopardy GameJeopardy

	Kahoot GameKahoot
}

func (v Game) MarshalJSON() ([]byte, error) {
	switch v.Game {
	case "jeopardy":
		return json.Marshal(struct { T string `json:"game"`; GameJeopardy }{ v.Game, v.Jeopardy })
	case "kahoot":
		return json.Marshal(struct { T string `json:"game"`; GameKahoot }{ v.Game, v.Kahoot })
	}

	return nil, fmt.Errorf("bad Game value: %s", v.Game)
}

func (v *Game) UnmarshalJSON(b []byte) error {
	var t struct { T string `json:"game"` }
	if err := json.Unmarshal(b, &t); err != nil {
		return err
	}

	var err error
	switch t.T {
	case "jeopardy":
		err = json.Unmarshal(b, &v.Jeopardy)
	case "kahoot":
		err = json.Unmarshal(b, &v.Kahoot)
	default:
		err = fmt.Errorf("bad Game value: %s", t.T)
	}

	if err != nil {
		return err
	}

	v.Game = t.T
	return nil
}

type GameJeopardy struct {
	Data Jeopardy `json:"data"`
}

type GameKahoot struct {
	Data Kahoot `json:"data"`
}

// JeopardyGame is the game data for a Jeopardy game.
type Jeopardy struct {
	Categories []JeopardyCategory `json:"categories"`

	// moderators enables moderators being able to join.
	Moderators *bool `json:"moderators,omitempty"`

	// require_name, if true, will require members to input a name before
	// we can participate.
	RequireName *bool `json:"require_name,omitempty"`

	// score_multiplier is the score multiplier for each question. The
	// default is 100.
	ScoreMultiplier *float64 `json:"score_multiplier,omitempty"`
}

// JeopardyCategory is a category in a Jeopardy game.
type JeopardyCategory struct {
	// name is the name of the category.
	Name string `json:"name"`

	// questions are the questions in the category.
	Questions []JeopardyQuestion `json:"questions"`
}

// JeopardyQuestion is a question in a Jeopardy game.
type JeopardyQuestion struct {
	// answers are the possible answers.
	Answers []string `json:"answers"`

	// correct_answer is the correct answer within the list of answers
	// above. The index starts at 1.
	CorrectAnswer int32 `json:"correct_answer"`

	// question is the question.
	Question string `json:"question"`
}

// KahootGame is the game data for a Kahoot game.
type Kahoot struct {
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
	PlayerName string `json:"playerName"`

	Score int32 `json:"score"`
}

// PlayerName is the name of a player.
type PlayerName = string
