package ent

import (
	"context"
	log "github.com/sirupsen/logrus"
	pb "way-jasy-cron/cron-logger/api"
	"way-jasy-cron/cron-logger/internal/model/ent"
	"way-jasy-cron/cron-logger/internal/model/ent/logger"
	"way-jasy-cron/cron-logger/internal/model/ent_ex"
)

func (m *Manager) ListLog(ctx context.Context, req *ent_ex.ListLoggerOptions) (logs []*ent.Logger, total int, err error){
	total, err = m.Client.Logger.Query().Where(logger.OperatorEQ(req.Operator)).Count(ctx)
	if err != nil {
		log.Error("count logs total err:(%v)", err)
		return nil, 0, err
	}
	logs, err = m.Client.Logger.Query().Where(logger.OperatorEQ(req.Operator)).
		Offset(req.OffSet()).Limit(req.Limit()).Order(ent.Desc(logger.FieldID)).All(ctx)
	if err != nil {
		log.Error("list logs err:(%v)", err)
		return nil, 0, err
	}
	return
}

func (m *Manager) CreateLog(ctx context.Context, req *pb.WriteLogReq) (*pb.NoReply, error) {
	_, err := m.Client.Logger.Create().SetLog(req.Opt).SetOperator(req.Operator).Save(ctx)
	return &pb.NoReply{}, err
}
