BEGIN;
CREATE TABLE members
(
    id serial primary key,
    name varchar(255) not null unique,
    email varchar(255) not null unique
);
COMMIT;