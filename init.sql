create table if not exists delivery
(
    id      integer generated always as identity
        constraint id
            primary key,
    name    text not null,
    phone   text not null,
    zip     text not null,
    city    text not null,
    address text not null,
    region  text not null,
    email   text not null
);

alter table delivery
    owner to alex;

create table if not exists payment
(
    id            integer generated always as identity
        constraint payment_pk
            primary key,
    transaction   text    not null,
    request_id    text    not null,
    currency      text    not null,
    provider      text    not null,
    amount        integer not null,
    payment_dt    integer not null,
    bank          text    not null,
    delivery_cost integer not null,
    goods_total   integer not null,
    custom_fee    integer not null
);

alter table payment
    owner to alex;

create table if not exists "order"
(
    order_uid          text    not null
        constraint order_pk
            primary key,
    track_number       text    not null,
    entry              text    not null,
    delivery_id        integer not null
        constraint order_delivery_fk
            references delivery
            on delete cascade,
    payment_id         integer not null
        constraint order_payment_fk
            references payment
            on delete cascade,
    locale             text    not null,
    internal_signature text    not null,
    customer_id        text    not null,
    delivery_service   text    not null,
    shardkey           text    not null,
    sm_id              integer not null,
    date_created       text    not null,
    oof_shard          text    not null
);

alter table "order"
    owner to alex;

create table if not exists item
(
    id           integer generated always as identity
        constraint item_pk
            primary key,
    order_uid    text    not null
        constraint item_order_fk
            references "order",
    chrt_id      integer not null,
    track_number text    not null,
    price        integer not null,
    rid          text    not null,
    name         text    not null,
    sale         integer not null,
    size         text    not null,
    total_price  integer not null,
    nm_id        integer not null,
    brand        text    not null,
    status       integer not null
);

alter table item
    owner to alex;

