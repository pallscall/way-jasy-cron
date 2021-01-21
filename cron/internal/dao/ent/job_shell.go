package ent

import (
	"context"
	log "github.com/sirupsen/logrus"
	"way-jasy-cron/cron/internal/model/ent"
	"way-jasy-cron/cron/internal/model/ent/machine"
	"way-jasy-cron/cron/internal/model/ent_ex"
)

func (m *Manager) ListShellJob(ctx context.Context, req *ent_ex.ListShellJobOptions) (jobs []*ent.Machine,total int, err error){
	total, err = m.Client.Machine.Query().Where(machine.StatusNEQ(int(ent_ex.ShellJobDelete))).Count(ctx)
	if err != nil {
		log.Error("method:ListShellJob#ent, count total err:(%v)", err)
		return nil, 0, err
	}
	jobs, err = m.Client.Machine.Query().Where(machine.StatusNEQ(int(ent_ex.ShellJobDelete))).Where(machine.HostContains(req.Host)).
		Where(machine.CommentContains(req.Comment)).
		Offset(req.OffSet()).Limit(req.Limit()).Order(ent.Desc(machine.FieldID)).All(ctx)
	if err != nil {
		log.Error("method:ListShellJob#ent, list shell jobs err:(%v)", err)
		return nil, 0, err
	}
	return
}

func (m *Manager) CreateShellJob(ctx context.Context, j *ent.Machine) error {
	_, err := m.Client.Machine.Create().
		SetHost(j.Host).SetUsername(j.Username).SetPassword(j.Password).SetComment(j.Comment).
		SetCommand(j.Command).SetStatus(j.Status).SetPort(j.Port).
		Save(ctx)
	return err
}

func (m *Manager) UpdateShellJob(ctx context.Context, j *ent.Machine) error {
	_, err := m.Client.Machine.Update().
		SetHost(j.Host).SetUsername(j.Username).SetPassword(j.Password).SetComment(j.Comment).
		SetCommand(j.Command).SetStatus(j.Status).SetPort(j.Port).Save(ctx)
	return err
}

func (m *Manager) DeleteShellJob(ctx context.Context, id int) error {
	_, err := m.Client.Machine.Update().SetStatus(int(ent_ex.ShellJobDelete)).Where(machine.IDEQ(id)).Save(ctx)
	return err
}

func (m *Manager) StartShellJob(ctx context.Context, id int) error {
	_, err := m.Client.Machine.Update().SetStatus(int(ent_ex.ShellJobDelete)).Where(machine.IDEQ(id)).Save(ctx)
	return err
}

func (m *Manager) StopShellJob(ctx context.Context, id int) error {
	_, err := m.Client.Machine.Update().SetStatus(int(ent_ex.ShellJobDelete)).Where(machine.IDEQ(id)).Save(ctx)
	return err
}