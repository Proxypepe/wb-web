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

INSERT INTO public.delivery (name, phone, zip, city, address, region, email)
VALUES ('Test Testov', '+9720000000', '2639809', 'Kiryat Mozkin', 'Ploshad Mira 15', 'Kraiot', 'test@gmail.com');
INSERT INTO public.payment (transaction, request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee)
VALUES ('b563feb7b2b84b6test', '', 'USD', 'wbpay', 1817, 1637907727, 'alpha', 1500, 317, 0);
INSERT INTO public.order (order_uid, track_number, entry, delivery_id, payment_id, locale, internal_signature, customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard)
VALUES ('b563feb7b2b84b6test', 'WBILMTESTTRACK', 'WBIL', 1, 1, 'en', '', 'test', 'meest', '9', 99, '2023-03-30 21:31:00.5773854 +0300 MSK m=+1.111687201', '1');
INSERT INTO public.item (order_uid, chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status)
VALUES ('b563feb7b2b84b6test', 9934930, 'WBILMTESTTRACK', 453, 'ab4219087a764ae0btest', 'Mascaras', 30, '0', 317, 2389212, 'Vivienne Sabo', 202);
INSERT INTO public.item (order_uid, chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status)
VALUES ('b563feb7b2b84b6test', 9934932, 'WBILMTESTTRACK', 433, 'ab4219087a764ae0btest', 'Mascaras', 34, '0', 315, 2389212, 'Vivienne Sabo', 202);

