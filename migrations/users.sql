create table users(mail varchar(100) not null, f_name varchar(50), l_name varchar(100), pass varchar(255) not null,
active integer default 1, created_at integer, updated_at integer, deleted_at integer);

-- creating index(s)
create index idx_user_mail on users(mail);
