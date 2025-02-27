create table files(file_name varchar(200), file_location varchar(255), file_md5 varchar,
created_at integer, updated_at integer, deleted_at integer, bucket varchar, owner varchar(100));

-- creating index(s)
create index idx_file_name on files(file_name);
create index idx_file_location on files(file_location);
