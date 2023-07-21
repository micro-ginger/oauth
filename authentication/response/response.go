package response

type Response interface {
	SetChallenge(challenge string)
}

type BaseResponse struct {
	State     string `json:"state,omitempty"`
	Challenge string `json:"challenge,omitempty"`
	Remaining uint   `json:"remaining,omitempty"`
	Detail    any   `json:"detail,omitempty"`
}

func (r *BaseResponse) SetChallenge(challenge string) {
	r.Challenge = challenge
}
