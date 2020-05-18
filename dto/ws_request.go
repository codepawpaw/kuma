package dto

type WsRequest struct {
	Action  string `json:"action"`
	Channel string `json:"channel"`
}
