package service

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"time"
	"way-jasy-cron/common/ecron"
	pb "way-jasy-cron/cron/api"
	"way-jasy-cron/cron/ecode"
	"way-jasy-cron/cron/internal/model/ent"
	"way-jasy-cron/cron/internal/model/ent_ex"
)

func (svc *Service) QueryJob(ctx context.Context, id int) (*ent.Job, error) {
	job, err := svc.ent.QueryJobByID(ctx, id)
	if err != nil {
		log.Error("QueryJobByID err: (%+v)", err)
		return nil, err
	}
	return job, nil
}

func (svc *Service) CreateJob(ctx context.Context, job *ent.Job) error {
	if _, err := ecron.ParseStandard(job.Spec); err != nil {
		return errors.Wrap(ecode.InvalidSpec, "job spec is incorrect！")
	}
	err := svc.ent.CreateJob(ctx, job)
	if err != nil {
		log.Error("CreateJob err: (%+v)", err)
		return err
	}
	if err := svc.WriteCronLogger(ctx, fmt.Sprintf("添加了任务%v", job),job.Creator);
			err != nil {
		log.Warn("Write create Log err:", err)
		return nil
	}
	return err
}

func (svc *Service) UpdateJob(ctx context.Context, job *ent.Job) error {
	err := svc.ent.UpdateJob(ctx, job)
	if err != nil {
		log.Error("CreateJob err: (%+v)", err)
	}
	j := svc.acquireJobExecutor(job)
	svc.cron.Remove(j.EntryID())
	if ent_ex.JobStatus(job.Status) == ent_ex.JobRunning {
		svc.cron.AddJob(j)
	}
	if err := svc.WriteCronLogger(ctx, fmt.Sprintf("更新了任务%v", job), job.Creator);
		err != nil {
		log.Warn("Write update Log err:", err)
		return nil
	}
	return err
}

// -1:delete 0:stop 1:run
func (svc *Service) SwitchJobStatus(ctx context.Context, id, opt int) error {
	j, _ := svc.ent.QueryJobByID(ctx, id)
	switch ent_ex.JobStatus(opt) {
	case ent_ex.JobDelete:
		if err := svc.ent.DeleteJob(ctx, id); err != nil {
			log.Error("delete job err: (%+v)", err)
			return err
		}
		svc.cron.Remove(ecron.EntryID(id))
		if err := svc.WriteCronLogger(ctx, fmt.Sprintf("删除了任务，任务id(%d)", id), j.Creator);
			err != nil {
			log.Warn("Write delete Log err:", err)
		}
	case ent_ex.JobStopping:
		if err := svc.ent.StopJob(ctx, id); err != nil {
			log.Error("stop job err: (%+v)", err)
			return err
		}
		svc.cron.Remove(ecron.EntryID(id))
		if err := svc.WriteCronLogger(ctx, fmt.Sprintf("关闭了任务，任务id(%d)", id), j.Creator);
			err != nil {
			log.Warn("Write stop Log err:", err)
		}
	case ent_ex.JobRunning:
		if err := svc.ent.StartJob(ctx, id); err != nil {
			log.Error("start job err: (%+v)", err)
			return err
		}
		j, err := svc.ent.QueryJobByID(ctx, id)
		if err != nil {
			log.Error("method: SwitchJobStatus#job QueryJobByID err:", err)
			return err
		}
		_, err = svc.cron.AddJob(svc.acquireJobExecutor(j))
		if err := svc.WriteCronLogger(ctx, fmt.Sprintf("开启了任务，任务id(%d)", id), j.Creator);
			err != nil {
			log.Warn("Write run Log err:", err)
		}
	default:
		return ecode.InvalidOption
	}
	return nil
}

func (svc *Service) ListJob(ctx context.Context, req *ent_ex.ListJobReq) (*ent_ex.ListJobResp, error) {
	options := req.ToListJobOptions()
	jobs, total, err := svc.ent.ListJob(ctx, options)
	if err != nil {
		return nil, err
	}
	return &ent_ex.ListJobResp{
		Jobs: jobs,
		ListBaseResp: ent_ex.ListBaseResp{
			Total: int64(total),
			PN:    req.PN,
			PS:    req.PS,
		},
	}, nil
}

func (svc *Service) InitJob() error{
	jobs, err := svc.ent.ListAllRunningJobs(context.TODO())
	if err != nil {
		log.Error("InitJob err:", err)
		return err
	}
	for _, job := range jobs {
		if _, err := svc.cron.AddJob(svc.acquireJobExecutor(job)); err != nil {
			log.Error("AddJob err:", err)
			return err
		}
	}
	svc.cron.Start()
	return nil
}

func (svc *Service) acquireJobExecutor(job *ent.Job) *ent_ex.JobExecutor{
	return svc.jobex.Create(job, svc.SwitchJobStatus, svc.http, svc.mail, svc.WriteCronLogger)
}

func (svc *Service) WriteCronLogger(ctx context.Context, msg, operator string) (err error){
	t := time.Now()
	req := &pb.WriteLogReq{
		Opt: msg,
		Operator: operator,
		Ctime: &t,
	}
	_, err = svc.WriteLog(ctx, req)
	return err
}