create table item
(
    goen_id               int primary key,
    goen_in_all_instance  bool         not null default (false),
    barcode               int          not null default (0),
    name                  varchar(255) not null default (''),
    price                 float        not null default (0),
    stock_number          int          not null default (0),
    order_price           float        not null default (0),
    belonged_item_goen_id int
);
create table item_contained_item
(
    owner_goen_id      int,
    possession_goen_id int,
    primary key (owner_goen_id, possession_goen_id)
);
create table payment
(
    goen_id              int primary key,
    goen_in_all_instance bool  not null default (false),
    goen_type            int   not null check ( goen_type between 0 and 2),
    amount_tendered      float not null default (0)
);
create table card_payment
(
    goen_id             int primary key,
    card_account_number int      not null default (0),
    expiry_date         datetime not null default ('0001-01-01 00:00:00')
);
create table cash_payment
(
    goen_id int primary key,
    balance float not null default (0)
);



alter table item_contained_item
    add constraint foreign key (owner_goen_id) references item (goen_id) on delete cascade,
    add constraint foreign key (possession_goen_id) references item (goen_id) on delete cascade;
alter table item
    add constraint foreign key (belonged_item_goen_id) references item (goen_id) on delete set null;

alter table card_payment
    add constraint foreign key (goen_id) references payment (goen_id) on delete cascade;
alter table cash_payment
    add constraint foreign key (goen_id) references payment (goen_id) on delete cascade;