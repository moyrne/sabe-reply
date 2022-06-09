CREATE TABLE `chat_record`
(
    `id`          integer     not null auto_increment,
    `created_at`  datetime    not null,

    `kind`        varchar(20) not null comment '{"personal":"个人","group":"群聊"}',

    `sender`      varchar(50) not null,
    `receiver`    varchar(50) not null,
    `content`     text        not null,
    `raw_content` text        not null,
    primary key (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;