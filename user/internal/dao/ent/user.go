package ent

import (
	"context"
	"way-jasy-cron/user/ecode"
	"way-jasy-cron/user/internal/model/ent"
	"way-jasy-cron/user/internal/model/ent/user"
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

func (m *Manager) QueryUserInfo(ctx context.Context, accessKey string) (*ent.User, error){
	var u []*ent.User
	err := m.Client.User.Query().Where(user.UsernameEQ(accessKey)).
		Select(user.FieldUsername, user.FieldEmail, user.FieldTel).Scan(ctx, &u)
	if ent.IsNotFound(err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return u[0], nil
}

func (m *Manager) Register(ctx context.Context, u *ent.User) error {
	_, err := m.Client.User.Update().SetEmail(u.Email).
		SetPassword(u.Password).SetTel(u.Tel).Where(user.UsernameEQ(u.Username)).Save(ctx)
	return err
}

func (m *Manager) GetPublicKey(ctx context.Context, username string) (string, error) {
	var v []struct {
		PublicKey  string    `json:"public_key"`
	}
	err := m.Client.User.Query().Where(user.UsernameEQ(username)).Select("public_key").Scan(ctx, &v)
	if err != nil {
		return "", err
	}
	if v == nil {
		return "", ecode.InvalidUser
	}
	return v[0].PublicKey, nil
}

func (m *Manager) SaveRsaKey(ctx context.Context, username,pubKey, priKey string) error {
	_, err := m.Client.User.Create().SetPublicKey(pubKey).SetPrivateKey(priKey).SetUsername(username).Save(ctx)
	return err
}