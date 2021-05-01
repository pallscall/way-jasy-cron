alter table jobs add retry int(9) NOT NULL DEFAULT 0 COMMENT '失败重试次数，为0时任务停止';
alter table jobs add retry_temp int(9) NOT NULL DEFAULT 0 COMMENT '失败重试次数还原数';
alter table jobs add count int(9) NOT NULL DEFAULT 0 COMMENT '任务执行次数，为0时任务停止',
