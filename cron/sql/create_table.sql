CREATE TABLE `jobs`
(
    `id`            int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增 id',
    `name`          varchar(64)      NOT NULL DEFAULT '' COMMENT '任务名称',
    `url`           varchar(128)     NOT NULL DEFAULT '' COMMENT '调用路径',
    `spec`          varchar(64)      NOT NULL DEFAULT '' COMMENT '定时器',
    `status`        int(1)           NOT NULL DEFAULT 1 COMMENT '任务状态',
    `comment`       varchar(128)     NOT NULL DEFAULT '' COMMENT '任务描述',
    `creator`       varchar(32)      NOT NULL DEFAULT 'admin' COMMENT '创建者',
    `updater`       varchar(32)      NOT NULL DEFAULT 'admin' COMMENT '更新者',
    `method`        VARCHAR(10)      NOT NULL DEFAULT 'GET' COMMENT '请求方法',
    `body`          text             NOT NULL COMMENT '请求体',
    `header`        varchar(1024)    NOT NULL DEFAULT '' COMMENT '请求头',
    `stoppable`     int(1)           NOT NULL DEFAULT 0 COMMENT '任务失败是否停止',
    `ctime`         datetime         NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `mtime`         datetime         NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
    PRIMARY KEY (`id`),
    INDEX idx_name (name (64)),
    INDEX idx_creator (`creator`),
    KEY `ix_mtime` (`mtime`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8 COMMENT ='定时任务表';

CREATE TABLE `machines`
(
    `id`            int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增 id',
    `host`          varchar(15)      NOT NULL DEFAULT '' COMMENT '主机ip',
    `port`          int              NOT NULL DEFAULT 22 COMMENT '端口',
    `username`      varchar(64)      NOT NULL DEFAULT '' COMMENT '登录名',
    `password`      varchar(64)      NOT NULL DEFAULT '' COMMENT '登录密码',
    `comment`       varchar(128)     NOT NULL DEFAULT '' COMMENT '任务描述',
    `command`       text             NOT NULL DEFAULT '' COMMENT '任务指令',
    `status`        int(1)           NOT NULL DEFAULT 1  COMMENT '任务状态',
    `ctime`         datetime         NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `mtime`         datetime         NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
    PRIMARY KEY (`id`),
    INDEX idx_host (`host`),
    KEY `ix_mtime` (`mtime`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8 COMMENT ='shell任务表';
