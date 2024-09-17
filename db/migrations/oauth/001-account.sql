-- accounts
create table accounts (
    id bigint unsigned not null primary key AUTO_INCREMENT,
    created_at datetime not null default CURRENT_TIMESTAMP,
    updated_at datetime null on update CURRENT_TIMESTAMP,
    internal_status bigint unsigned not null default 0,
    status bigint unsigned not null,
    hashed_password binary(72) null,
    index accounts_created_at_index (created_at),
    index accounts_updated_at_index (updated_at),
    index accounts_internal_status_index (internal_status),
    index accounts_status_index (status)
);