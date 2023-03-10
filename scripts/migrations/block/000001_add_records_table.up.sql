create table if not exists records
(
    height      bigint                                                                                    not null comment 'height of block',
    hash        char(66)     default '0x0000000000000000000000000000000000000000000000000000000000000000' not null comment 'hash of block',
    parent_hash char(66)     default '0x0000000000000000000000000000000000000000000000000000000000000000' not null comment 'hash of parent block',
    timestamp   timestamp(6) default 0                                                                    not null comment 'timestamp',
    constraint records_pk primary key (hash)
) engine = InnoDB
  default charset = utf8mb4
  collate = utf8mb4_general_ci
    comment = 'block record list';

create index height_index on records (height);
create index timestamp_index on records (timestamp);
