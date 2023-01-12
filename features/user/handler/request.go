package handler

import "api/features/user"

type LoginRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type RegisterRequest struct {
	Nama     string `json:"nama" form:"nama"`
	Email    string `json:"email" form:"email"`
	Alamat   string `json:"alamat" form:"alamat"`
	HP       string `json:"hp" form:"hp"`
	Password string `json:"password" form:"password"`
}
type UpdateRequest struct {
	Name    string `json:"name" form:"name"`
	Email   string `json:"email" form:"email"`
	Address string `json:"address" form:"address"`
	HP      string `json:"hp" form:"hp"`
}

func ToCore(data interface{}) *user.Core {
	res := user.Core{}

	switch data.(type) {
	case LoginRequest:
		cnv := data.(LoginRequest)
		res.Email = cnv.Email
		res.Password = cnv.Password
	case RegisterRequest:
		cnv := data.(RegisterRequest)
		res.Email = cnv.Email
		res.Nama = cnv.Nama
		res.Alamat = cnv.Alamat
		res.HP = cnv.HP
		res.Password = cnv.Password
	case UpdateRequest:
		cnv := data.(UpdateRequest)
		res.Nama = cnv.Name
		res.Email = cnv.Email
		res.HP = cnv.HP
		res.Alamat = cnv.Address
	default:
		return nil
	}

	return &res
}
