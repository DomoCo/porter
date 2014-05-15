package request_handler

import (
	"api"
	"encoding/json"
	"engine/game_engine"
	"github.com/Tactique/golib/logger"
)

func respondSuccess(payload interface{}) []byte {
	return generateResponse(payload, api.STATUS_OK)
}

func respondFailure(payload interface{}) []byte {
	return generateResponse(payload, api.STATUS_FAILURE)
}

func respondMalformed(payload interface{}) []byte {
	return generateResponse(payload, api.STATUS_MALFORMED_RESPONSE)
}

func respondBadRequest(payload interface{}) []byte {
	return generateResponse(payload, api.STATUS_BAD_REQUEST)
}

func respondUnknownRequest(payload interface{}) []byte {
	return generateResponse(payload, api.STATUS_UNKNOWN_REQUEST)
}

func generateResponse(payload interface{}, status int) []byte {
	response, err := json.Marshal(api.ResponseType{Status: status, Payload: payload})
	if err != nil {
		logger.Warnf("Could not generate response: %s", err.Error())
		backupResponse, err := json.Marshal(api.ResponseType{
			Status:  api.STATUS_UNSERIALIZEABLE_RESPONSE,
			Payload: "problem"})
		if err != nil {
			return []byte("Really bad")
		}
		logger.Infof("Generating response with status %d", status)
		logger.Debugf("Full message %s", string(response))
		return backupResponse
	}
	logger.Infof("Generating response with status %d", status)
	logger.Debugf("Full message %s", string(response))
	return response
}

func newRequest(jsonRequest []byte) ([]byte, *game_engine.World) {
	var request api.NewRequest
	err := json.Unmarshal(jsonRequest, &request)
	if err != nil {
		return respondMalformed(nil), nil
	}
	game, err := game_engine.NewWorld(request.Uids, request.Debug)
	if err != nil {
		return respondBadRequest(err.Error()), nil
	} else {
		return respondSuccess(nil), game
	}
}

func viewWorldRequest(jsonRequest []byte, playerId int, game *game_engine.World) []byte {
	var request api.ViewPlayersRequest
	err := json.Unmarshal(jsonRequest, &request)
	if err != nil {
		return respondMalformed(nil)
	}
	response, err := game.ViewWorld(playerId)
	if err != nil {
		return respondBadRequest(err.Error())
	}
	return respondSuccess(response)
}

func viewPlayersRequest(jsonRequest []byte, playerId int, game *game_engine.World) []byte {
	var request api.ViewPlayersRequest
	err := json.Unmarshal(jsonRequest, &request)
	if err != nil {
		return respondMalformed(nil)
	}
	response, err := game.ViewPlayers(playerId)
	if err != nil {
		return respondBadRequest(err.Error())
	}
	return respondSuccess(response)
}

func viewUnitsRequest(jsonRequest []byte, playerId int, game *game_engine.World) []byte {
	var request api.ViewUnitsRequest
	err := json.Unmarshal(jsonRequest, &request)
	if err != nil {
		return respondMalformed(nil)
	}
	response, err := game.ViewUnits(playerId)
	if err != nil {
		return respondBadRequest(err.Error())
	}
	return respondSuccess(response)
}

func viewTerrainRequest(jsonRequest []byte, playerId int, game *game_engine.World) []byte {
	var request api.ViewTerrainRequest
	err := json.Unmarshal(jsonRequest, &request)
	if err != nil {
		return respondMalformed(nil)
	}
	response, err := game.ViewTerrain(playerId)
	if err != nil {
		return respondBadRequest(err.Error())
	}
	return respondSuccess(response)
}

func moveRequest(jsonRequest []byte, playerId int, game *game_engine.World) []byte {
	var request api.MoveRequest
	err := json.Unmarshal(jsonRequest, &request)
	if err != nil {
		return respondMalformed(nil)
	}
	response, err := game.MoveUnit(playerId, request.Move)
	if err != nil {
		return respondBadRequest(err.Error())
	}
	return respondSuccess(response)
}

func attackRequest(jsonRequest []byte, playerId int, game *game_engine.World) []byte {
	var request api.AttackRequest
	err := json.Unmarshal(jsonRequest, &request)
	if err != nil {
		return respondMalformed(nil)
	}
	response, err := game.Attack(playerId, request.Attacker, request.AttackIndex, request.Target)
	if err != nil {
		return respondBadRequest(err.Error())
	}
	return respondSuccess(response)
}

func endTurnRequest(jsonRequest []byte, playerId int, game *game_engine.World) []byte {
	var request api.MoveRequest
	err := json.Unmarshal(jsonRequest, &request)
	if err != nil {
		return respondMalformed(nil)
	}
	response, err := game.EndTurn(playerId)
	if err != nil {
		return respondBadRequest(err.Error())
	}
	return respondSuccess(response)
}

func exitRequest(jsonRequest []byte, playerId int, game *game_engine.World) []byte {
	var request api.ExitRequest
	err := json.Unmarshal(jsonRequest, &request)
	if err != nil {
		return respondMalformed(nil)
	}
	return respondSuccess(nil)
}