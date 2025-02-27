create table sessions(user varchar(100), registration_token varchar, login_token varchar,
invite_token varchar, created_at integer, updated_at integer);

-- creating index(s)
create index idx_user_session on sessions(user);
create index idx_login_session on sessions(login_token);
