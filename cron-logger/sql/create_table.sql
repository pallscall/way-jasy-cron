CREATE TABLE `loggers`
(
    `id`            int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增 id',
    `log`           varchar(256)     NOT NULL DEFAULT '' COMMENT '日志信息',
    `operator`      varchar(64)      NOT NULL DEFAULT '' COMMENT '操作人（日志创建者）',
    `ctime`         datetime         NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    PRIMARY KEY (`id`),
    INDEX idx_operator (operator (64)),
    KEY `ix_ctime` (`ctime`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8 COMMENT ='操作日志表';
