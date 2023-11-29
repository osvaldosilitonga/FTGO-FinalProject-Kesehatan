package web

type WebResponse struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Data   any    `json:"data"`
}

type ErrWebResponse struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Detail any    `json:"detail"`
}
