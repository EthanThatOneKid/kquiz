import * as jtd from "jtd";

export default {
  definitions: {
    Command: {
      discriminator: "type",
      mapping: {
        EndGame: {
          properties: {
            data: {
              ref: "CommandEndGame",
            },
          },
        },
        JeopardyChooseQuestion: {
          properties: {
            data: {
              ref: "CommandJeopardyChooseQuestion",
            },
          },
        },
        JeopardyPlayerIsCorrect: {
          properties: {
            data: {
              ref: "CommandJeopardyPlayerIsCorrect",
            },
          },
        },
        JeopardyPressButton: {
          properties: {
            data: {
              ref: "CommandJeopardyPressButton",
            },
          },
        },
        JoinGame: {
          properties: {
            data: {
              ref: "CommandJoinGame",
            },
          },
        },
      },
    },
    CommandEndGame: {
      additionalProperties: false,
      metadata: {
        description:
          "CommandEndGame is sent by a client to end the current game. The server\nwill respond with an EventGameEnded. Only game moderators (including the\nhost) can end the game.\n",
      },
      optionalProperties: {},
      properties: {
        declareWinner: {
          metadata: {
            description:
              "declareWinner determines whether the game should be ended with a\nwinner or not. If true, the game will be ended with a winner. If\nfalse, the game will be ended abruptly.\n",
          },
          type: "boolean",
        },
      },
    },
    CommandJeopardyChooseQuestion: {
      additionalProperties: false,
      metadata: {
        description:
          "CommandJeopardyChooseQuestion is sent by a player to choose a question.\nThe server must do validation to ensure that the player is allowed to\nchoose the question.\n",
      },
      optionalProperties: {},
      properties: {
        category: {
          type: "string",
        },
        question: {
          type: "string",
        },
      },
    },
    CommandJeopardyPlayerIsCorrect: {
      metadata: {
        description:
          "CommandJeopardyPlayerIsCorrect is emitted by a game moderator to indicate\nthat a player has answered a question correctly. The winning player is\nwhoever the last EventJeopardyButtonPressed event indicated. That player\nwill instantly receive the points for the question, and the game will let\nthem choose the next category and question.\n",
      },
    },
    CommandJeopardyPressButton: {
      metadata: {
        description:
          "CommandJeopardyPressButton is emitted when a player presses the button\nduring a question. It is only valid to emit this command when the game is\nin the question state.\n",
      },
    },
    CommandJoinGame: {
      additionalProperties: false,
      metadata: {
        description:
          "CommandJoinGame is sent by a client to join a game. The client (or the\nuser) supplies a game ID and a player name. The server will respond with\nan EventJoinedGame.\n",
      },
      optionalProperties: {},
      properties: {
        gameID: {
          metadata: {
            description: "gameID is the ID of the game to join.",
          },
          type: "string",
        },
        moderatorPassword: {
          metadata: {
            description:
              "moderatorPassword is the password of the moderator of the game.",
          },
          nullable: true,
          type: "string",
        },
        playerName: {
          metadata: {
            description: "playerName is the wanted name of the user.",
          },
          ref: "PlayerName",
        },
      },
    },
    Error: {
      additionalProperties: false,
      metadata: {
        description: "Error is returned on every API error.\n",
      },
      optionalProperties: {},
      properties: {
        message: {
          metadata: {
            description: "Message is the error message",
          },
          type: "string",
        },
      },
    },
    Event: {
      discriminator: "type",
      mapping: {
        GameEnded: {
          properties: {
            data: {
              ref: "EventGameEnded",
            },
          },
        },
        JeopardyBeginQuestion: {
          properties: {
            data: {
              ref: "EventJeopardyBeginQuestion",
            },
          },
        },
        JeopardyButtonPressed: {
          properties: {
            data: {
              ref: "EventJeopardyButtonPressed",
            },
          },
        },
        JeopardyResumeButton: {
          properties: {
            data: {
              ref: "EventJeopardyResumeButton",
            },
          },
        },
        JeopardyTurnEnded: {
          properties: {
            data: {
              ref: "EventJeopardyTurnEnded",
            },
          },
        },
        JoinedGame: {
          properties: {
            data: {
              ref: "EventJoinedGame",
            },
          },
        },
        PlayerJoined: {
          properties: {
            data: {
              ref: "EventPlayerJoined",
            },
          },
        },
      },
    },
    EventGameEnded: {
      additionalProperties: false,
      metadata: {
        description: "EventGameEnded is emitted when the current game ends.\n",
      },
      optionalProperties: {},
      properties: {
        leaderboard: {
          ref: "Leaderboard",
        },
      },
    },
    EventJeopardyBeginQuestion: {
      additionalProperties: false,
      metadata: {
        description:
          "EventJeopardyBeginQuestion is emitted when a question begins within this\nJeopardy game. It is usually emitted once the chooser player has chosen a\ncategory and value.\n\nEach category name and question value will map to a category and question\nwithin the game data. Note that a question may repeat across multiple\ncategories.\n",
      },
      optionalProperties: {},
      properties: {
        category: {
          type: "string",
        },
        chooser: {
          ref: "PlayerName",
        },
        question: {
          type: "string",
        },
      },
    },
    EventJeopardyButtonPressed: {
      additionalProperties: false,
      metadata: {
        description:
          'EventJeopardyButtonPressed is emitted when any player had pressed a button\non their device, voiding other players\' buttons. This event is only\nemitted when the game is in the "question" state.\n',
      },
      optionalProperties: {},
      properties: {
        playerName: {
          ref: "PlayerName",
        },
      },
    },
    EventJeopardyResumeButton: {
      additionalProperties: false,
      metadata: {
        description:
          "EventJeopardyResumeButton is emitted when the player can now continue to\npress the button whenever they are ready to answer the question. This\ncould happen if the other player who pressed the button first got the\nquestion wrong.\n\nNote that if alreadyPressed is true, then the player has already pressed\nthe button, so they cannot press it again.\n",
      },
      optionalProperties: {},
      properties: {
        alreadyPressed: {
          type: "boolean",
        },
      },
    },
    EventJeopardyTurnEnded: {
      additionalProperties: false,
      metadata: {
        description:
          "EventJeopardyTurnEnded is emitted when a turn ends or when the game first\nstarts.\n",
      },
      optionalProperties: {},
      properties: {
        currentScore: {
          type: "float64",
        },
        isChooser: {
          type: "boolean",
        },
        leaderboard: {
          ref: "Leaderboard",
        },
      },
    },
    EventJoinedGame: {
      metadata: {
        description:
          "EventJoinedGame is emitted when the current player joins a game. It is a\nreply to CommandJoinGame and is only for the current player. Not to be\nconfused with EventPlayerJoinedGame, which is emitted when any player\njoins the current game.\n",
      },
    },
    EventPlayerJoined: {
      additionalProperties: false,
      metadata: {
        description:
          "EventPlayerJoined is emitted when a player joins the current game.\n",
      },
      optionalProperties: {},
      properties: {
        playerName: {
          ref: "PlayerName",
        },
      },
    },
    Game: {
      discriminator: "game",
      mapping: {
        jeopardy: {
          properties: {
            data: {
              ref: "Jeopardy",
            },
          },
        },
        kahoot: {
          properties: {
            data: {
              ref: "Kahoot",
            },
          },
        },
      },
      metadata: {
        description:
          "Game is the main game object. It contains all the information about the\ngame.\n",
      },
    },
    Jeopardy: {
      additionalProperties: false,
      metadata: {
        description: "JeopardyGame is the game data for a Jeopardy game.\n",
      },
      optionalProperties: {
        moderators: {
          metadata: {
            description: "moderators enables moderators being able to join.\n",
          },
          type: "boolean",
        },
        require_name: {
          metadata: {
            description:
              "require_name, if true, will require members to input a name before\nwe can participate.\n",
          },
          type: "boolean",
        },
        score_multiplier: {
          metadata: {
            description:
              "score_multiplier is the score multiplier for each question. The\ndefault is 100.\n",
          },
          type: "float64",
        },
      },
      properties: {
        categories: {
          elements: {
            ref: "JeopardyCategory",
          },
        },
      },
    },
    JeopardyCategory: {
      additionalProperties: false,
      metadata: {
        description: "JeopardyCategory is a category in a Jeopardy game.\n",
      },
      optionalProperties: {},
      properties: {
        name: {
          metadata: {
            description: "name is the name of the category.\n",
          },
          type: "string",
        },
        questions: {
          elements: {
            ref: "JeopardyQuestion",
          },
          metadata: {
            description: "questions are the questions in the category.\n",
          },
        },
      },
    },
    JeopardyGameInfo: {
      additionalProperties: false,
      metadata: {
        description:
          "JeopardyGameInfo is the initial information for a Jeopardy game. This type\ncontains no useful information about the entire game data, so it's used to\nsend to players the first time they join.\n",
      },
      optionalProperties: {},
      properties: {
        categories: {
          elements: {
            type: "string",
          },
        },
        numQuestions: {
          type: "int32",
        },
        players: {
          elements: {
            ref: "PlayerName",
          },
        },
        scoreMultiplier: {
          type: "int32",
        },
      },
    },
    JeopardyQuestion: {
      additionalProperties: false,
      metadata: {
        description: "JeopardyQuestion is a question in a Jeopardy game.\n",
      },
      optionalProperties: {},
      properties: {
        question: {
          metadata: {
            description: "question is the question.\n",
          },
          type: "string",
        },
      },
    },
    Kahoot: {
      additionalProperties: false,
      metadata: {
        description: "KahootGame is the game data for a Kahoot game.\n",
      },
      optionalProperties: {},
      properties: {
        questions: {
          elements: {
            ref: "KahootQuestion",
          },
          metadata: {
            description: "questions are the questions in the game.\n",
          },
        },
        time_limit: {
          metadata: {
            description:
              "time_limit is the time limit for each question. The format is in\nGo's time.Duration, e.g. 10s for 10 seconds.\n",
          },
          type: "string",
        },
      },
    },
    KahootQuestion: {
      additionalProperties: false,
      metadata: {
        description: "KahootQuestion is a question in a Kahoot game.\n",
      },
      optionalProperties: {},
      properties: {
        answers: {
          elements: {
            type: "string",
          },
          metadata: {
            description: "answers are the possible answers.\n",
          },
        },
        question: {
          metadata: {
            description: "question is the question.\n",
          },
          type: "string",
        },
      },
    },
    Leaderboard: {
      elements: {
        ref: "LeaderboardEntry",
      },
      metadata: {
        description: "Leaderboard is a list of players and their scores.\n",
      },
    },
    LeaderboardEntry: {
      additionalProperties: false,
      optionalProperties: {},
      properties: {
        playerName: {
          type: "string",
        },
        score: {
          type: "int32",
        },
      },
    },
    PlayerName: {
      metadata: {
        description: "PlayerName is the name of a player.\n",
      },
      type: "string",
    },
  },
} as jtd.Schema;
