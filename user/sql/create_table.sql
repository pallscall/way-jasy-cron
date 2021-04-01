CREATE TABLE `jobs`
(
    `id`            int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增 id',
    `username`      varchar(64)      NOT NULL DEFAULT '' COMMENT '昵称',
    `password`      varchar(500)      NOT NULL DEFAULT '' COMMENT '密码',
    `public_key`    varchar (500)   NOT NULL DEFAULT '' COMMENT '公钥',
    `private_key`   varchar (500)   NOT NULL DEFAULT '' COMMENT '公钥';
    `tel`           char(11)         NOT NULL DEFAULT '' COMMENT '电话',
    `email`         varchar(64)      NOT NULL DEFAULT '' COMMENT '邮箱',
    `rtime`         datetime         NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '注册时间',
    PRIMARY KEY (`id`),
    INDEX idx_username (`username`),
    KEY `ix_ctime` (`ctime`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8 COMMENT ='用户表';
