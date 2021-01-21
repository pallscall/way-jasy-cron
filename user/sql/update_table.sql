ALTER TABLE `users` ADD `public_key` VARCHA(500) NOT NULL DEFAULT '' COMMENT '公钥';
ALTER TABLE `users` ADD `private_key` VARCHA(1024) NOT NULL DEFAULT '' COMMENT '私钥';
alter table users modify password varchar(500) default '' not null comment '密码';
