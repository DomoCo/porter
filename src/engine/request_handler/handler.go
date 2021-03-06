package request_handler

import (
	"api"
	"bytes"
	"github.com/Tactique/golib/logger"
	"strconv"
)

type RequestHandler struct {
	sessionGameWrapper *GameWrapper
	gameRequest        map[string]func([]byte, int, *GameWrapper) []byte
}

func NewRequestHandler() *RequestHandler {
	return &RequestHandler{
		sessionGameWrapper: nil,
		gameRequest: map[string]func([]byte, int, *GameWrapper) []byte{
			api.COMMAND_EXIT:         exitRequest,
			api.COMMAND_MOVE:         moveRequest,
			api.COMMAND_TURN:         endTurnRequest,
			api.COMMAND_ATTACK:       attackRequest,
			api.COMMAND_VIEW_WORLD:   viewWorldRequest,
			api.COMMAND_VIEW_TERRAIN: viewTerrainRequest,
			api.COMMAND_VIEW_UNITS:   viewUnitsRequest,
			api.COMMAND_VIEW_PLAYERS: viewPlayersRequest}}
}

func (handler *RequestHandler) HandleRequest(request []byte) []byte {
	command, requestJson := splitOnce(request)
	logger.Infof("Got command %s and request json %s", string(command), string(requestJson))
	if handler.sessionGameWrapper == nil {
		if string(command) == api.COMMAND_NEW {
			response, game := newRequest(requestJson)
			handler.sessionGameWrapper = game
			logger.Infof("After new game request, game is now %t nil", (handler.sessionGameWrapper == nil))
			return buildResponse(command, response)
		} else {
			return buildResponse(command, respondUnknownRequest("Need new game request"))
		}
	} else {
		fun, ok := handler.gameRequest[string(command)]
		if ok {
			playerId, requestJsonNoPlayerId := splitOnce(requestJson)
			playerIdInt, err := strconv.Atoi(string(playerId))
			if err != nil {
				logger.Warnf("Not a playerId %s (%s)", playerId, err.Error())
				return buildResponse(command, respondMalformed("playerId not an int"))
			}
			logger.Infof("request for playerId %d", playerIdInt)
			response := fun(requestJsonNoPlayerId, playerIdInt, handler.sessionGameWrapper)
			return buildResponse(command, response)
		} else {
			logger.Warnf("Unknown Command %s", string(command))
			return buildResponse(command, respondUnknownRequest("Unknown command"))
		}
	}
}

func buildResponse(command []byte, response []byte) []byte {
	return append(append(command, byte(':')), response...)
}

func splitOnce(input []byte) ([]byte, []byte) {
	pieces := bytes.SplitN(input, []byte(":"), 2)
	if len(pieces) == 1 {
		return pieces[0], []byte{}
	} else if len(pieces) == 2 {
		return pieces[0], pieces[1]
	} else {
		return []byte{}, []byte{}
	}
}
