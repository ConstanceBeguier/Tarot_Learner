define({ "api": [
  {
    "type": "get",
    "url": "/hand/:id",
    "title": "Request Hand information.",
    "name": "GetHandEndpoint",
    "group": "Hand",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "Number",
            "optional": false,
            "field": "id",
            "description": "<p>Users unique ID.</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "List",
            "optional": false,
            "field": "cards",
            "description": "<p>List of cards.</p>"
          }
        ]
      }
    },
    "version": "0.0.0",
    "filename": "./rest.go",
    "groupTitle": "Hand"
  },
  {
    "type": "get",
    "url": "/newparty/available_seats",
    "title": "Request seats availability.",
    "name": "GetNewpartyAvailableseatsEndpoint",
    "group": "Newparty",
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "List",
            "optional": false,
            "field": "availableSeats",
            "description": "<p>List of booleans.</p>"
          }
        ]
      }
    },
    "version": "0.0.0",
    "filename": "./rest.go",
    "groupTitle": "Newparty"
  },
  {
    "type": "get",
    "url": "/newparty",
    "title": "Start a new party",
    "name": "GetNewpartyEndpoint",
    "group": "Newparty",
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "Boolean",
            "optional": false,
            "field": "succeed",
            "description": "<p>Does the party successfuly start.</p>"
          }
        ]
      }
    },
    "version": "0.0.0",
    "filename": "./rest.go",
    "groupTitle": "Newparty"
  },
  {
    "type": "get",
    "url": "/newparty/status",
    "title": "Request if all seats are ready.",
    "name": "GetNewpartyStatusEndpoint",
    "group": "Newparty",
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "Boolean",
            "optional": false,
            "field": "ready",
            "description": "<p>Readyness of the party.</p>"
          }
        ]
      }
    },
    "version": "0.0.0",
    "filename": "./rest.go",
    "groupTitle": "Newparty"
  },
  {
    "type": "post",
    "url": "/newparty/available_seats/:id",
    "title": "Take place in the 'id' seat.",
    "name": "PostNewpartyAvailableseatsEndpoint",
    "group": "Newparty",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "Number",
            "optional": false,
            "field": "id",
            "description": "<p>Users unique ID.</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "List",
            "optional": false,
            "field": "availableSeats",
            "description": "<p>List of booleans.</p>"
          }
        ]
      }
    },
    "version": "0.0.0",
    "filename": "./rest.go",
    "groupTitle": "Newparty"
  },
  {
    "type": "get",
    "url": "/table/:id",
    "title": "Request Hand information.",
    "name": "GetTableEndpoint",
    "group": "Table",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "Number",
            "optional": false,
            "field": "id",
            "description": "<p>Users unique ID.</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "[2]float32",
            "optional": false,
            "field": "scores",
            "description": "<p>Actual score of attacker/defender.</p>"
          },
          {
            "group": "Success 200",
            "type": "[NB_PLAYERS]Card",
            "optional": false,
            "field": "cards",
            "description": "<p>on the table.</p>"
          },
          {
            "group": "Success 200",
            "type": "Integer",
            "optional": false,
            "field": "playerTurn",
            "description": "<p>ID of the player turn.</p>"
          },
          {
            "group": "Success 200",
            "type": "Integer",
            "optional": false,
            "field": "firstPlayer",
            "description": "<p>ID of the first player who played.</p>"
          },
          {
            "group": "Success 200",
            "type": "Integer",
            "optional": false,
            "field": "trickNb",
            "description": "<p>Trick's number.</p>"
          },
          {
            "group": "Success 200",
            "type": "[NB_PLAYERS]int",
            "optional": false,
            "field": "isAttacker",
            "description": "<p>Return the attacker status of players.</p>"
          }
        ]
      }
    },
    "version": "0.0.0",
    "filename": "./rest.go",
    "groupTitle": "Table"
  },
  {
    "type": "get",
    "url": "/table/trick",
    "title": "Request Trick information.",
    "name": "GetTablePlayerTurnEndpoint",
    "group": "Table",
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "Boolean",
            "optional": false,
            "field": "playerTurn",
            "description": "<p>Current trick.</p>"
          }
        ]
      }
    },
    "version": "0.0.0",
    "filename": "./rest.go",
    "groupTitle": "Table"
  },
  {
    "type": "get",
    "url": "/table/:trick/:id",
    "title": "Get ready for the next trick.",
    "name": "GetTableTrickIdEndpoint",
    "group": "Table",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "Number",
            "optional": false,
            "field": "trick",
            "description": "<p>Trick Number.</p>"
          },
          {
            "group": "Parameter",
            "type": "Number",
            "optional": false,
            "field": "id",
            "description": "<p>Users unique ID.</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "Empty",
            "optional": false,
            "field": "Empty",
            "description": "<p>Empty brace.</p>"
          }
        ]
      }
    },
    "version": "0.0.0",
    "filename": "./rest.go",
    "groupTitle": "Table"
  },
  {
    "type": "post",
    "url": "/table/:id/:color/:number",
    "title": "Play a card.",
    "name": "PostTableEndpoint",
    "group": "Table",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "Number",
            "optional": false,
            "field": "id",
            "description": "<p>Users unique ID.</p>"
          },
          {
            "group": "Parameter",
            "type": "Number",
            "optional": false,
            "field": "color",
            "description": "<p>Color of the playing card.</p>"
          },
          {
            "group": "Parameter",
            "type": "Number",
            "optional": false,
            "field": "number",
            "description": "<p>Number of the playing card.</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "Boolean",
            "optional": false,
            "field": "succeed",
            "description": "<p>Does the card can be played.</p>"
          }
        ]
      }
    },
    "version": "0.0.0",
    "filename": "./rest.go",
    "groupTitle": "Table"
  }
] });
