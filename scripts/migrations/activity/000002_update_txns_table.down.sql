alter table txns
    drop column timestamp;

alter table txns
    drop column nonce;

alter table txns
    drop column data;

alter table txns
    drop column value;

alter table txns
    drop column events;
