create table links(user varchar(100) not null, file_location varchar, temp_link varchar, created_at integer);

-- creating index(s)
create index idx_link on links(user, temp_link);
