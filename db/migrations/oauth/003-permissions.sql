insert into scopes(name, state, description)
values ('api.auth.account.read', 0b01, 'read current logged-in account'),
       ('api.auth.account.profile.read', 0b01, 'read current logged-in profile'),
       ('api.auth.account.profile.update', 0b01, 'update current profile'),
       ('api.auth.password.update', 0b00, 'update password access'),
       ('api.auth.password.reset', 0b00, 'reset password access'),
       ('api.auth.accounts.manage', 0b10, 'access to manage existing accounts in system'),
       ('api.auth.register', 0b00, 'access to register')
ON DUPLICATE KEY
    UPDATE updated_at=now();

insert ignore into roles(name, state, description)
values ('user-base', 0b1, 'default base user permissions');

insert ignore into roles(name, state, description)
values ('user', 0b0, 'user permissions');

insert ignore into roles(name, state, description)
values ('admin', 0b0, 'admin permissions in admin panel');

insert ignore into role_scopes(role_id, scope_id)
values ((select id from roles where name = 'admin'),
        (select id from scopes where name = 'api.auth.accounts.manage'));

insert ignore into role_scopes(role_id, scope_id)
values ((select id from roles where name = 'user-base'),
        (select id from scopes where name = 'api.auth.account.profile.read'));

insert ignore into role_scopes(role_id, scope_id)
values ((select id from roles where name = 'user'),
        (select id from scopes where name = 'api.auth.account.read')),
       ((select id from roles where name = 'user'),
        (select id from scopes where name = 'api.auth.account.profile.update')),
       ((select id from roles where name = 'user'),
        (select id from scopes where name = 'api.auth.password.update'));
