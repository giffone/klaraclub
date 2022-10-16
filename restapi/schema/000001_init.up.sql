create sequence if not exists my_sequence
    increment 1
    start 1
    minvalue 1
;

create table if not exists users
(
    id   integer default nextval('my_sequence'::regclass) not null
        primary key,
    data varchar
);