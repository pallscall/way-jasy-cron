package ent

import (
	"context"
	"graduate/user/internal/model/ent"
	"graduate/user/internal/model/ent/user"
)

func (m *Manager) QueryUser(ctx context.Context, accessKey string) (*ent.User, error){
	u, err := m.Client.User.Query().Where(user.UsernameEQ(accessKey)).Only(ctx)
	if ent.IsNotFound(err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (m *Manager) Register(ctx context.Context, user *ent.User) error {
	_, err := m.Client.User.Create().SetUsername(user.Username).SetEmail(user.Email).
		SetPassword(user.Password).SetTel(user.Tel).Save(ctx)
	return err
}
