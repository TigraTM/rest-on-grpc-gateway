BEGIN;

create schema payment;

create table if not exists "payment".accounts
(
    id             serial primary key,
    create_at      timestamp default now(),
    update_at      timestamp default now(),
    account_number text    not null,
    balance        decimal not null,
    currency       text      default 'RUB',
    user_id        integer not null,

    constraint user_id_account_number_constraint unique (user_id, account_number),
    constraint account_number_constraint unique (account_number),
    constraint balance_constraint check (balance >= 0)
);

create table if not exists "payment".payment_history
(
    id           serial primary key,
    create_at    timestamp default now(),
    update_at    timestamp default now(),
    sum          decimal not null,
    old_balance  decimal not null,
    company_name text    not null,
    category     text    not null,
    account_id   integer not null,

    constraint accounts_id_fk foreign key (account_id)
        references "payment".accounts (id)
);

END;