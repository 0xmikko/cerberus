package users

import (
	"context"
	"github.com/MikaelLazarev/cerberus/server/core"
)

func (s *service) GetUserById(ctx context.Context, id core.ID) (*core.User, error) {
	return s.store.FindByID(id)
}
