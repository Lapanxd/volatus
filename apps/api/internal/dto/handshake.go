package dto

type InitRequest struct {
	ToUserID uint `json:"to-user-id" binding:"required"`
}

type InitResponse struct {
	SessionID string `json:"session-id"`
}
