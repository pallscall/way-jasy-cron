package service

import (
	"context"
	"github.com/prometheus/common/log"
	pb "way-jasy-cron/cron-logger/api"
	"way-jasy-cron/cron-logger/internal/model/ent_ex"
)

func (s *Service) WriteLog(ctx context.Context, req *pb.WriteLogReq) (*pb.NoReply, error) {
	_, err := s.ent.CreateLog(ctx, req)
	if err != nil {
		log.Error("CreateLog err:", err)
	}
	return &pb.NoReply{}, err
}

func (s *Service) ListLog(ctx context.Context, req *ent_ex.ListLoggerReq) (resp *ent_ex.ListLoggerResp, err error) {
	options := req.ToListLoggerOptions()
	logs, total, err := s.ent.ListLog(ctx, options)
	if err != nil {
		return nil, err
	}
	return &ent_ex.ListLoggerResp{
		Logs: logs,
		ListBaseResp: ent_ex.ListBaseResp{
			Total: int64(total),
			PN:    req.PN,
			PS:    req.PS,
		},
	}, nil
}