-- profiles
create table profiles
(
    id         bigint unsigned not null primary key AUTO_INCREMENT,
    updated_at datetime        null on update CURRENT_TIMESTAMP,
    index profiles_updated_at_index (updated_at),
    constraint profiles_accounts_id_fk
        foreign key (id) references accounts (id)
);
