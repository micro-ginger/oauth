-- registers
create table registers
(
    id              bigint unsigned not null primary key AUTO_INCREMENT,
    created_at      datetime        not null default CURRENT_TIMESTAMP,
    updated_at      datetime        null on update CURRENT_TIMESTAMP,
    account_id      bigint unsigned not null,
    hashed_password varchar(72)     null,
    constraint registers_accounts_id_fk
        foreign key (account_id) references accounts (id),
    index accounts_created_at_index (created_at),
    index accounts_updated_at_index (updated_at),
    unique index accounts_account_id_uindex (account_id)
);
