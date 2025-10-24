package service

import (
	"errors"
	"sync"

	"github.com/google/uuid"
	"github.com/lapanxd/volatus-api/internal/dto"
)

type HandshakeSession struct {
	FromUserID uint
	ToUserID   uint
	SDPOffer   string
}

// todo: persist sessions in database when needed (redis or postgres)
var (
	handshakeStore = make(map[string]HandshakeSession)
	storeMutex     = sync.Mutex{}
)

func InitHandshake(fromUserID, toUserID uint, sdpOffer string) (string, error) {
	if fromUserID == toUserID {
		return "", errors.New("cannot handshake with yourself")
	}

	for _, session := range handshakeStore {
		if session.FromUserID == fromUserID && session.ToUserID == toUserID {
			return "", errors.New("handshake request already exists")
		}
		if session.FromUserID == toUserID && session.ToUserID == fromUserID {
			return "", errors.New("handshake request already exists in reverse")
		}
	}

	sessionID := uuid.NewString()
	storeMutex.Lock()
	defer storeMutex.Unlock()
	handshakeStore[sessionID] = HandshakeSession{
		FromUserID: fromUserID,
		ToUserID:   toUserID,
		SDPOffer:   sdpOffer,
	}

	return sessionID, nil
}

func RespondHandshake(userID uint, sessionID string, accepted bool) (*HandshakeSession, error) {
	storeMutex.Lock()
	defer storeMutex.Unlock()

	session, found := handshakeStore[sessionID]
	if !found {
		return nil, errors.New("session not found")
	}

	if session.ToUserID != userID {
		return nil, errors.New("not your session")
	}

	delete(handshakeStore, sessionID)

	if !accepted {
		return nil, nil
	}

	return &session, nil
}

func GetPendingHandshakes(userID uint) []dto.Pending {
	storeMutex.Lock()
	defer storeMutex.Unlock()

	pending := make([]dto.Pending, 0)
	for id, session := range handshakeStore {
		if session.ToUserID == userID {
			pending = append(pending, dto.Pending{
				SessionID:  id,
				FromUserID: session.FromUserID,
				SDPOffer:   session.SDPOffer,
			})
		}
	}

	return pending
}
