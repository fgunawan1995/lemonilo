package model

type Login struct {
	UserIDOrEmail string `json:"user_id_or_email" validate:"required"`
	Password      string `json:"password" validate:"required"`
}

func (p Login) Validate() error {
	return validate.Struct(p)
}
