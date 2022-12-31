create table if not exists txns
(
    hash       char(66) default '0x0000000000000000000000000000000000000000000000000000000000000000' not null comment 'hash of tx',
    `from`     char(66) default '0x0000000000000000000000000000000000000000000000000000000000000000' not null comment 'from',
    `to`       char(66) default '0x0000000000000000000000000000000000000000000000000000000000000000' not null comment 'to',
    block_hash char(66) default '0x0000000000000000000000000000000000000000000000000000000000000000' not null comment 'confirmed from block hash',
    constraint txns_pk primary key (hash)
) engine = InnoDB
  default charset = utf8mb4
  collate = utf8mb4_general_ci
    comment = 'transaction';

create index idx_block_hash on txns (block_hash);
