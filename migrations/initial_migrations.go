package migrations

const Initial = `
create database employeeservice;
create table if not exists public.employees
(
    id         serial primary key,
    name       varchar(30),
    lastname   varchar(30),
    patronymic varchar(30),
    birthday   date
);

alter table public.employees
    owner to postgres;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

create table if not exists public.files(
                                           id uuid default uuid_generate_v4() not null,
                                           filename varchar(30),
                                           insert_date date,
                                           data bytea not null
);

alter table public.files
    owner to postgres;


create table if not exists public.Reports(
                                             id serial,
                                             name varchar(30),
                                             text text,
                                             employee_id int references employees(id) on delete cascade
);

insert into employees (name) values ('abc') returning id;
`
