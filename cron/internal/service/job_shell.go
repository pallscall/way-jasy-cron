package service

import (
	"context"
	log "github.com/sirupsen/logrus"
	"way-jasy-cron/cron/ecode"
	"way-jasy-cron/cron/internal/model/ent"
	"way-jasy-cron/cron/internal/model/ent_ex"
)

func (svc *Service) ListShellJob(ctx context.Context, req *ent_ex.ListShellJobReq) (*ent_ex.ListShellJobResp, error) {
	options := req.ToListShellJobOptions()
	jobs, total, err := svc.ent.ListShellJob(ctx, options)
	if err != nil {
		return nil, err
	}
	return &ent_ex.ListShellJobResp{
		Jobs: jobs,
		ListBaseResp: ent_ex.ListBaseResp{
			Total: int64(total),
			PN:    req.PN,
			PS:    req.PS,
		},
	}, nil
}


func (svc *Service) CreateShellJob(ctx context.Context, machine *ent.Machine) error {
	err := svc.ent.CreateShellJob(ctx, machine)
	if err != nil {
		log.Error("CreateJob err: (%+v)", err)
	}
	return err
}

func (svc *Service) UpdateShellJob(ctx context.Context, machine *ent.Machine) error {
	err := svc.ent.UpdateShellJob(ctx, machine)
	if err != nil {
		log.Error("CreateJob err: (%+v)", err)
	}
	if ent_ex.ShellJobStatus(machine.Status) == ent_ex.ShellJobRunning {
	}
	return err
}

// -1:delete 0:stop 1:run
func (svc *Service) SwitchShellJobStatus(ctx context.Context, id, opt int) error {
	switch ent_ex.ShellJobStatus(opt) {
	case ent_ex.ShellJobDelete:
		if err := svc.ent.DeleteShellJob(ctx, id); err != nil {
			log.Error("delete job err: (%+v)", err)
			return err
		}
	case ent_ex.ShellJobStopping:
		if err := svc.ent.StopShellJob(ctx, id); err != nil {
			log.Error("stop job err: (%+v)", err)
			return err
		}

	case ent_ex.ShellJobRunning:
		if err := svc.ent.StartShellJob(ctx, id); err != nil {
			log.Error("start job err: (%+v)", err)
			return err
		}
		//j, err := svc.ent.QueryShJobByID(ctx, id)
		//if err != nil {
		//	log.Error("method: SwitchJobStatus#job QueryJobByID err:", err)
		//	return err
		//}
	default:
		return ecode.InvalidOption
	}
	return nil
}
