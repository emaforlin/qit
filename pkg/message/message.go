package message

type CreateMessageDto struct {
	ID      string `json:"id,omitempty" validate:"omitempty,min=1,max=100,alphanum"`
	Payload any    `json:"payload" validate:"required"`
}
