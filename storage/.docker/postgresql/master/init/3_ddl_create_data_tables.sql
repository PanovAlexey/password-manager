create table if not exists users
(
    id serial constraint users_pk primary key,
    email varchar(255) not null,
    password varchar(255) not null,
    created_at timestamp not null,
    last_access_at timestamp
);

alter table users owner to pw_user;
create unique index if not exists users_email_uindex on users (email);
