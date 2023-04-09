package migrations

const initial = `
create table public.employees
(
    id         serial primary key 
        primary key,
    name       varchar(30),
    lastname   varchar(30),
    patronymic varchar(30),
    birthday   date
);

alter table public.employees
    owner to postgres;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
create table public.files(
    id uuid default uuid_generate_v4() not null,
    filename varchar(30),
    insert_date date,
    data bytea not null
);

alter table public.files
    owner to postgres;


create table Reports(
    id serial,
    name varchar(30),
    text text,
    employee_id int references employees(id) on delete cascade
)
`
