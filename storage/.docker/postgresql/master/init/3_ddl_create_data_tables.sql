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

create table if not exists login_password
(
    id serial constraint login_password_pk primary key,
    name varchar(255) not null,
    login varchar(255) not null,
    password varchar(255) not null,
    note text not null,
    user_id int not null constraint login_password_user_id_fk references users,
    created_at timestamp not null,
    last_access_at timestamp
);

alter table login_password owner to pw_user;

create table if not exists credit_card
(
    id serial constraint credit_card_pk primary key,
    name varchar(255) not null,
    number varchar(30) not null,
    expiration varchar(10) not null,
    cvv varchar(6) not null,
    owner varchar(100) not null,
    note text not null,
    user_id int not null constraint credit_card_user_id_fk references users,
    created_at timestamp not null,
    last_access_at timestamp
);

alter table credit_card owner to pw_user;

create table if not exists text_record
(
    id serial constraint text_record_pk primary key,
    name varchar(255) not null,
    text text not null,
    note text not null,
    user_id int not null constraint text_record_user_id_fk references users,
    created_at timestamp not null,
    last_access_at timestamp
);

alter table text_record owner to pw_user;

create table if not exists binary_record
(
    id serial constraint binary_record_pk primary key,
    name varchar(255) not null,
    binary text not null,
    note text not null,
    user_id int not null constraint binary_record_user_id_fk references users,
    created_at timestamp not null,
    last_access_at timestamp
);

alter table binary_record owner to pw_user;