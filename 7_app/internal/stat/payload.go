package stat

type GetStatResponse struct {
	Period        string `json:"period"`
	DirectionsSum uint   `json:"directions_sum"`
}
