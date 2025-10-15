package dto

type InitInput struct {
	ToUserID uint   `json:"to_user_id" binding:"required"`
	SDPOffer string `json:"sdp_offer" binding:"required"`
}

type InitResponse struct {
	SessionID string `json:"session_id"`
}

type ResponseInput struct {
	SessionID string `json:"session_id"`
	Accepted  bool   `json:"accepted"`
	SDPAnswer string `json:"sdp_answer"`
}

type Pending struct {
	SessionID  string `json:"session_id"`
	FromUserID uint   `json:"from_user_id"`
}

type PendingOutput struct {
	PendingSessions []Pending `json:"pending_sessions"`
}
