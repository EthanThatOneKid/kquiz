// Code generated by jtd-codegen for TypeScript v0.2.1

export type Qg = any;

export type Command =
  | CommandBeginGame
  | CommandEndGame
  | CommandJeopardyChooseQuestion
  | CommandJeopardyPlayerJudgment
  | CommandJeopardyPressButton
  | CommandJoinGame;

/**
 * CommandBeginGame is sent by a client to begin a game.
 */
export interface CommandBeginGame {
  type: "BeginGame";
}

/**
 * CommandEndGame is sent by a client to end the current game. The server
 * will respond with an EventGameEnded. Only game admins (including the
 * host) can end the game.
 */
export interface CommandEndGame {
  type: "EndGame";

  /**
   * declareWinner determines whether the game should be ended with a
   * winner or not. If true, the game will be ended with a winner. If
   * false, the game will be ended abruptly.
   */
  declareWinner: boolean;
}

/**
 * CommandJeopardyChooseQuestion is sent by a player to choose a question.
 * The server must do validation to ensure that the player is allowed to
 * choose the question.
 */
export interface CommandJeopardyChooseQuestion {
  type: "JeopardyChooseQuestion";
  category: number;
  question: number;
}

/**
 * CommandJeopardyPlayerJudgment is emitted by a game admin to indicate
 * whether a player has answered a question correctly. The winning player is
 * whoever the last EventJeopardyButtonPressed event indicated. That player
 * will instantly receive the points for the question, and the game will let
 * them choose the next category and question. If the player answered wrong,
 * then the game will let others press the button.
 */
export interface CommandJeopardyPlayerJudgment {
  type: "JeopardyPlayerJudgment";
  correct: boolean;
}

/**
 * CommandJeopardyPressButton is emitted when a player presses the button
 * during a question. It is only valid to emit this command when the game is
 * in the question state.
 */
export interface CommandJeopardyPressButton {
  type: "JeopardyPressButton";
}

/**
 * CommandJoinGame is sent by a client to join a game. The client (or the
 * user) supplies a game ID and a player name. The server will respond with
 * an EventJoinedGame.
 */
export interface CommandJoinGame {
  type: "JoinGame";

  /**
   * adminPassword is the password of the admin of the game.
   */
  adminPassword: string | null;

  /**
   * gameID is the ID of the game to join.
   */
  gameID: GameId;

  /**
   * playerName is the wanted name of the user.
   */
  playerName: PlayerName;
}

/**
 * Error is returned on every API error.
 */
export interface Error {
  /**
   * Message is the error message
   */
  message: string;
}

export type Event =
  | EventError
  | EventGameEnded
  | EventGameStarted
  | EventJeopardyBeginQuestion
  | EventJeopardyButtonPressed
  | EventJeopardyResumeButton
  | EventJeopardyTurnEnded
  | EventJoinedGame
  | EventPlayerJoined;

export interface EventError {
  type: "Error";
  error: Error;
}

/**
 * EventGameEnded is emitted when the current game ends.
 */
export interface EventGameEnded {
  type: "GameEnded";
  leaderboard: Leaderboard;
}

/**
 * EventGameStarted is emitted when the game starts. It contains no data and
 * is only meant to be used to trigger the client to start the game.
 */
export interface EventGameStarted {
  type: "GameStarted";
}

/**
 * EventJeopardyBeginQuestion is emitted when a question begins within this
 * Jeopardy game. It is usually emitted once the chooser player has chosen a
 * category and value.
 *
 * Each category name and question value will map to a category and question
 * within the game data. Note that a question may repeat across multiple
 * categories.
 */
export interface EventJeopardyBeginQuestion {
  type: "JeopardyBeginQuestion";
  category: number;
  chooser: PlayerName;
  points: number;
  question: string;
}

/**
 * EventJeopardyButtonPressed is emitted when any player had pressed a button
 * on their device, voiding other players' buttons. This event is only
 * emitted when the game is in the "question" state.
 */
export interface EventJeopardyButtonPressed {
  type: "JeopardyButtonPressed";
  playerName: PlayerName;
}

/**
 * EventJeopardyResumeButton is emitted when the player can now continue to
 * press the button whenever they are ready to answer the question. This
 * could happen if the other player who pressed the button first got the
 * question wrong.
 *
 * Note that if alreadyPressed is true, then the player has already pressed
 * the button, so they cannot press it again.
 */
export interface EventJeopardyResumeButton {
  type: "JeopardyResumeButton";
  alreadyAnsweredPlayers: PlayerName[];
}

/**
 * EventJeopardyTurnEnded is emitted when a turn ends or when the game first
 * starts.
 */
export interface EventJeopardyTurnEnded {
  type: "JeopardyTurnEnded";
  answered: JeopardyAnsweredQuestions;
  chooser: PlayerName;
  leaderboard: Leaderboard;
}

/**
 * EventJoinedGame is emitted when the current player joins a game. It is a
 * reply to CommandJoinGame and is only for the current player. Not to be
 * confused with EventPlayerJoinedGame, which is emitted when any player
 * joins the current game.
 */
export interface EventJoinedGame {
  type: "JoinedGame";
  gameData: GameData | null;
  gameID: string;
  gameInfo: GameInfo;
  isAdmin: boolean;
}

/**
 * EventPlayerJoined is emitted when a player joins the current game.
 */
export interface EventPlayerJoined {
  type: "PlayerJoined";
  playerName: PlayerName;
}

/**
 * GameData is the game data. It contains all the information about the game.
 */
export type GameData = GameDataJeopardy | GameDataKahoot;

export interface GameDataJeopardy {
  game: "jeopardy";
  data: JeopardyGameData;
}

export interface GameDataKahoot {
  game: "kahoot";
  data: KahootGameData;
}

/**
 * GameID is the unique identifier for a game. Each player must type this
 * code to join the game.
 */
export type GameId = string;

export type GameInfo = GameInfoJeopardy;

export interface GameInfoJeopardy {
  type: "jeopardy";
  data: JeopardyGameInfo;
}

export enum GameType {
  Jeopardy = "jeopardy",
  Kahoot = "kahoot",
}

export interface JeopardyAnsweredQuestion {
  category: number;
  player: PlayerName;
  question: number;
}

/**
 * JeopardyAnsweredQuestions is the list of answered questions for a player.
 */
export type JeopardyAnsweredQuestions = JeopardyAnsweredQuestion[];

/**
 * JeopardyCategory is a category in a Jeopardy game.
 */
export interface JeopardyCategory {
  /**
   * name is the name of the category.
   */
  name: string;

  /**
   * questions are the questions in the category.
   */
  questions: JeopardyQuestion[];
}

/**
 * JeopardyGameData is the game data for a Jeopardy game.
 */
export interface JeopardyGameData {
  categories: JeopardyCategory[];

  /**
   * score_multiplier is the score multiplier for each question. The
   * default is 100.
   */
  score_multiplier?: number;
}

/**
 * JeopardyGameInfo is the initial information for a Jeopardy game. This type
 * contains no useful information about the entire game data, so it's used to
 * send to players the first time they join.
 */
export interface JeopardyGameInfo {
  categories: string[];
  numQuestions: number;
  scoreMultiplier: number;
}

/**
 * JeopardyQuestion is a question in a Jeopardy game.
 */
export interface JeopardyQuestion {
  /**
   * question is the question.
   */
  question: string;
}

/**
 * KahootGameData is the game data for a Kahoot game.
 */
export interface KahootGameData {
  /**
   * questions are the questions in the game.
   */
  questions: KahootQuestion[];

  /**
   * time_limit is the time limit for each question. The format is in
   * Go's time.Duration, e.g. 10s for 10 seconds.
   */
  time_limit: string;
}

/**
 * KahootQuestion is a question in a Kahoot game.
 */
export interface KahootQuestion {
  /**
   * answers are the possible answers.
   */
  answers: string[];

  /**
   * question is the question.
   */
  question: string;
}

/**
 * Leaderboard is a list of players and their scores.
 */
export type Leaderboard = LeaderboardEntry[];

export interface LeaderboardEntry {
  playerName: string;
  score: number;
}

/**
 * PlayerName is the name of a player.
 */
export type PlayerName = string;

export interface RequestGetGame {
  gameID: string;
}

export interface RequestGetJeopardyGame {
  gameID: string;
}

export interface RequestNewGame {
  admin_password: string;
  data: GameData;
}

export interface ResponseGetGame {
  gameType: GameType;
}

export interface ResponseGetJeopardyGame {
  info: JeopardyGameInfo;
}

export interface ResponseNewGame {
  gameID: string;
  gameType: GameType;
}
