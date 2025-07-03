package model

type ErrorResponse struct {
	Error struct {
		Code    int16  `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}
