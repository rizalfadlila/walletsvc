create table wallet
(
    id          varchar(255)   default ''::character varying not null
        constraint id
        primary key,
    owned_by    varchar(255)   default ''::character varying not null,
    balance     numeric(18, 2) default 0                     not null,
    enabled_at  varchar(13)    default ''::character varying,
    disabled_at varchar(13)    default ''::character varying not null,
    status      varchar(15)    default ''::character varying not null
);

alter table wallet
    owner to postgres;

