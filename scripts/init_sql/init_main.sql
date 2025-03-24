
create table if not exists currencies
(
    currency_id serial primary key,
    name        text  not null,
    type_size   float not null
);

create table if not exists bids
(
    user_id       text      not null,
    currency_id   int       not null references currencies (currency_id),
    status        text  default 'created',
    create_date   timestamp not null,
    complete_date timestamp,
    min_price     float     not null,
    max_price     float     not null,
    amount_to_buy float     not null,
    bought_amount float default 0.0,
    buy_speed     float,
    avg_price     float
);

create table if not exists transactions
(
    user_id     text not null,
    currency_id int  not null references currencies (currency_id),
    type        text,
    amount      float,
    status      text default 'created',
    time        timestamp
);


create table if not exists user_balance
(
    user_id     text not null,
    currency_id int  not null references currencies (currency_id),
    balance     float
);

create table if not exists currency_rate
(
    from_id    int not null references currencies (currency_id),
    to_id      int not null references currencies (currency_id),
    cost       float,
    valid_from timestamp
);