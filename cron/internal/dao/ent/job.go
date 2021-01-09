package ent

import (
	"context"
	log "github.com/sirupsen/logrus"
	"way-jasy-cron/cron/internal/model/ent"
	"way-jasy-cron/cron/internal/model/ent/job"
	"way-jasy-cron/cron/internal/model/ent_ex"
)

func (m *Manager) QueryJobByID(ctx context.Context, id int)(*ent.Job, error) {
	j, err := m.Client.Job.Query().Where(job.IDEQ(id)).Only(ctx)
	return j, err
}

func (m *Manager) CreateJob(ctx context.Context, j *ent.Job) error {
	_, err := m.Client.Job.Create().
		SetName(j.Name).SetBody(j.Body).SetComment(j.Comment).SetSpec(j.Spec).SetHeader(j.Header).
		SetMethod(j.Method).SetCreator(j.Creator).SetStoppable(j.Stoppable).
		SetStatus(j.Status).SetUpdater(j.Updater).SetURL(j.URL).
		Save(ctx)
	return err
}

func (m *Manager) UpdateJob(ctx context.Context, j *ent.Job) error {
	_, err := m.Client.Job.Update().
		SetName(j.Name).SetBody(j.Body).SetComment(j.Comment).SetSpec(j.Spec).SetHeader(j.Header).
		SetMethod(j.Method).SetCreator(j.Creator).SetStoppable(j.Stoppable).
		SetStatus(j.Status).SetUpdater(j.Updater).SetURL(j.URL).
		Where(job.IDEQ(j.ID)).Save(context.TODO())
	return err
}

func (m *Manager) DeleteJob(ctx context.Context, id int) error {
	_, err := m.Client.Job.Update().SetStatus(int(ent_ex.JobDelete)).Where(job.IDEQ(id)).Save(ctx)
	return err
}

func (m *Manager) StartJob(ctx context.Context, id int) error {
	_, err := m.Client.Job.Update().SetStatus(int(ent_ex.JobRunning)).Where(job.IDEQ(id)).Save(ctx)
	return err
}

func (m *Manager) StopJob(ctx context.Context, id int) error {
	_, err := m.Client.Job.Update().SetStatus(int(ent_ex.JobStopping)).Where(job.IDEQ(id)).Save(ctx)
	return err
}

func (m *Manager) ListJob(ctx context.Context, req *ent_ex.ListJobOptions) (jobs []*ent.Job, total int, err error) {
	total, err = m.Client.Job.Query().Where(job.StatusNEQ(int(ent_ex.JobDelete))).Count(ctx)
	if err != nil {
		log.Error("count total err:(%v)", err)
		return nil, 0, err
	}
	jobs, err = m.Client.Job.Query().Where(job.StatusNEQ(int(ent_ex.JobDelete))).Where(job.NameContains(req.Name)).
		Offset(req.OffSet()).Limit(req.Limit()).Order(ent.Desc(job.FieldID)).All(ctx)
	if err != nil {
		log.Error("list jobs err:(%v)", err)
		return nil, 0, err
	}
	return
}


func (m *Manager) ListAllRunningJobs(ctx context.Context) (jobs []*ent.Job, err error) {
	return m.Client.Job.Query().Where(job.StatusNEQ(int(ent_ex.JobDelete))).All(ctx)
}