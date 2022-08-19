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
alter table item_contained_item
    add constraint foreign key (owner_goen_id) references item (goen_id) on delete cascade,
    add constraint foreign key (possession_goen_id) references item (goen_id) on delete cascade;
alter table item
    add constraint foreign key (belonged_item_goen_id) references item (goen_id) on delete set null;