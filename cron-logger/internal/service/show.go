package service

import "context"

func (svc *Service) Show(ctx context.Context) error{
	return svc.ent.Show(ctx)
}
