package handler

import "github.com/google/wire"

type Handler struct {
	User *User
}

var ProviderSet = wire.NewSet(NewUser, wire.Struct(new(Handler), "*"))
