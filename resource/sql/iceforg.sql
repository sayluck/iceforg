create table t_user
(
    id            bigint auto_increment comment '序号'
        primary key,
    name          varchar(20)                           not null,
    password      varchar(20)                           not null comment '用户密码',
    nick_name     varchar(30) default ''                null comment '用户昵称',
    create_time   timestamp   default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP comment '创建时间',
    last_mod_time timestamp                             null comment '修改时间',
    constraint t_user_UN
        unique (name)
)
    comment '用户信息';

