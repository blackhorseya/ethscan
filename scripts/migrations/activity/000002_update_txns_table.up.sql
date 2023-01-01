alter table txns
    add timestamp timestamp default 0 not null comment 'block timestamp';

alter table txns
    add nonce bigint(6) default 0 not null comment 'nonce';

alter table txns
    add data text default '' not null comment 'data';

alter table txns
    add value text default '' not null comment 'value';

alter table txns
    add events json default '{}' not null comment 'events';
