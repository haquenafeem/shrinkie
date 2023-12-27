package model

type ShrinkRequest struct {
	Url string `json:"url"`
}

type ShrinkResponse struct {
	RedirectTo string `json:"redirect_to"`
	HexValue   string `json:"hex_value"`
	FullURL    string `json:"full_url"`
	IsSuccess  bool   `json:"is_success"`
	Err        string `json:"err"`
}
