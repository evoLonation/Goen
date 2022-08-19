create table store
(
    id        int primary key,
    name      varchar(255),
    address   varchar(255),
    is_opened bool
);
create table product_catalog
(
    id       int primary key,
    name     varchar(255),
    store_id int,
    foreign key (store_id) references store (id)

);
create table cash_desk
(
    id        int primary key,
    name      varchar(255),
    is_opened bool,
    store_id  int,
    foreign key (store_id) references store (id)
);
create table payment
(
    goen_type       int not null check ( goen_type between 0 and 2),
    goen_id         int auto_increment primary key,
    amount_tendered double
);
create table cash_payment
(
    goen_id         int auto_increment primary key,
    balance         double,
    goen_payment_id int unique,
    foreign key (goen_payment_id) references payment (goen_id)
);
create table card_payment
(
    goen_id             int auto_increment primary key,
    card_account_number varchar(255),
    expiry_date         datetime,
    goen_payment_id     int unique,
    foreign key (goen_payment_id) references payment (goen_id)
);
create table sale
(
    goen_id        int auto_increment primary key,
    time           datetime,
    is_complete    bool,
    amount         double,
    is_readyto_pay bool,
    payment_id     int unique,
    store_id       int,
    cash_desk_id   int,
    foreign key (payment_id) references payment (goen_id),
    foreign key (store_id) references store (id),
    foreign key (cash_desk_id) references cash_desk(id)
);
create table cashier
(
    id       int primary key,
    name     varchar(255),
    store_id int,
    foreign key (store_id) references store (id)
);
create table item
(
    barcode            int primary key,
    name               varchar(255),
    price              double,
    stock_number       int,
    order_price        double,
    store_id           int,
    product_catalog_id int,
    foreign key (store_id) references store (id),
    foreign key (product_catalog_id) references product_catalog (id)
);
create table sales_line_item
(
    goen_id   int auto_increment primary key,
    quantity  int,
    subamount double,
    item_id   int ,
    sale_id   int ,
    foreign key (item_id) references item (barcode),
    foreign key (sale_id) references sale (goen_id)
);

create table supplier
(
    id   int primary key,
    name varchar(255)
);
create table order_product
(
    id           int primary key,
    time         datetime,
    order_status int,
    amount       double,
    supplier_id  int,
    foreign key (supplier_id) references supplier (id)
);
create table order_entry
(
    goen_id          int auto_increment primary key,
    quantity         int,
    sub_amount       double,
    item_id          int,
    order_product_id int,
    foreign key (item_id) references item (barcode),
    foreign key (order_product_id) references order_product (id)
);


