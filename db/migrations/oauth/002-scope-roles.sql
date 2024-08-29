-- scopes
create table scopes (
    id bigint unsigned auto_increment,
    created_at datetime not null default CURRENT_TIMESTAMP,
    updated_at datetime null on update CURRENT_TIMESTAMP,
    name varchar(32) not null,
    state tinyint unsigned null,
    description varchar(128) null,
    constraint scopes_pk primary key (id),
    unique index scopes_name_uidx (name),
    index scopes_state_idx (state)
);
-- roles
create table roles (
    id bigint unsigned auto_increment,
    created_at datetime not null default CURRENT_TIMESTAMP,
    updated_at datetime null on update CURRENT_TIMESTAMP,
    name varchar(32) not null,
    state tinyint unsigned null,
    description varchar(128) null,
    constraint roles_pk primary key (id),
    unique index roles_name_uidx (name),
    index roles_state_idx (state)
);
-- role scopes
create table role_scopes (
    role_id bigint unsigned not null,
    scope_id bigint unsigned not null,
    constraint role_scopes_roles_id_fk foreign key (role_id) references roles (id) on delete cascade,
    constraint role_scopes_scopes_id_fk foreign key (scope_id) references scopes (id) on delete cascade
);
-- account roles
create table account_roles (
    account_id bigint unsigned not null,
    role_id bigint unsigned not null,
    is_authorized bool null,
    constraint account_roles_account_id_fk foreign key (account_id) references accounts (id) on delete cascade,
    constraint account_roles_roles_id_fk foreign key (role_id) references roles (id) on delete cascade,
    index account_roles_is_authorized_idx (is_authorized),
    unique index account_roles_account_id_role_id_uidx (account_id, role_id)
);
-- account roles
create table account_scopes (
    account_id bigint unsigned not null,
    scope_id bigint unsigned not null,
    is_authorized bool null,
    constraint account_scopes_account_id_fk foreign key (account_id) references accounts (id) on delete cascade,
    constraint account_scopes_scopes_id_fk foreign key (scope_id) references scopes (id) on delete cascade,
    index account_scopes_is_authorized_idx (is_authorized)
);