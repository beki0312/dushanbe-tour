create table name_infrastructure (
                                     id          bigserial primary key,
                                     name        text not null,
                                     title       text,
                                     address     text not null,
                                     phone       text not null,
                                     link        text,
                                     email       text,
                                     description text,
                                     image text,
                                     active bool default true,
                                     infrastructure bigint references infrastructure(id),
                                     language bigint references language(id),
                                     created_at  timestamp default CURRENT_TIMESTAMP,
                                     updated     timestamp
);