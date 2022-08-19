create table card_payment(
                             goen_id int primary key,
                             card_account_number varchar(255),
                             expiry_date datetime

);
create table cash_desk(
                          goen_id int primary key,
                          id int,
                          name varchar(255),
                          is_opened bool ,
                          goen_belonged_store int,
                          goen_association_cashdeskes int

);
create table cash_payment(
                             goen_id int primary key,
                             balance float

);
create table cashier(
                        goen_id int primary key,
                        id int,
                        name varchar(255) ,
                        goen_worked_store int,
                        goen_cashiers int

);
create table item(
                     goen_id int primary key auto_increment,
                     barcode int,
                     name varchar(255),
                     price float,
                     stock_number int,
                     order_price float ,
                     goen_belonged_catalog int,
                     goen_items int,
                     goen_contained_items int

);
create table order_entry(
                            goen_id int primary key,
                            quantity int,
                            sub_amount float ,
                            goen_item int,
                            goen_contained_entries int

);
create table order_product(
                              goen_id int primary key,
                              id int,
                              time datetime,
                              order_status int,
                              amount float ,
                              goen_supplier int

);
create table payment(
                        goen_id int primary key,
                        amount_tendered float ,
                        goen_belonged_sale int

);
create table product_catalog(
                                goen_id int primary key,
                                id int,
                                name varchar(255) ,
                                goen_productcatalogs int

);
create table sale(
                     goen_id int primary key,
                     time datetime,
                     is_complete bool,
                     amount float,
                     is_readyto_pay bool ,
                     goen_belongedstore int,
                     goen_belonged_cash_desk int,
                     goen_assoicated_payment int,
                     goen_sales int,
                     goen_contained_sales int

);
create table sales_line_item(
                                goen_id int primary key,
                                quantity int,
                                subamount float ,
                                goen_belonged_sale int,
                                goen_belonged_item int,
                                goen_contained_sales_line int

);
create table store(
                      goen_id int primary key,
                      id int,
                      name varchar(255),
                      address varchar(255),
                      is_opened bool

);
create table supplier(
                         goen_id int primary key,
                         id int,
                         name varchar(255)

);
alter table cash_desk
    add constraint foreign key (goen_belonged_store) references store (goen_id),
    add constraint foreign key (goen_association_cashdeskes)  references store (goen_id) 
;
alter table cashier
    add constraint foreign key (goen_worked_store) references store (goen_id),
    add constraint foreign key (goen_cashiers) references store (goen_id)
;
alter table item
    add constraint foreign key (goen_belonged_catalog) references product_catalog (goen_id),
    add constraint foreign key (goen_items) references store (goen_id),
    add constraint foreign key (goen_contained_items) references product_catalog (goen_id)
;
alter table order_entry
    add constraint foreign key (goen_item) references item (goen_id),
    add constraint foreign key (goen_contained_entries) references order_product (goen_id)
;
alter table order_product
    add constraint foreign key (goen_supplier) references supplier (goen_id)
;
alter table payment
    add constraint foreign key (goen_belonged_sale) references sale (goen_id)
;
alter table product_catalog
    add constraint foreign key (goen_productcatalogs) references store (goen_id)
;
alter table sale
    add constraint foreign key (goen_belongedstore) references store (goen_id),
    add constraint foreign key (goen_belonged_cash_desk) references cash_desk (goen_id),
    add constraint foreign key (goen_assoicated_payment) references payment (goen_id),
    add constraint foreign key (goen_sales) references store (goen_id),
    add constraint foreign key (goen_contained_sales) references cash_desk (goen_id)
;
alter table sales_line_item
    add constraint foreign key (goen_belonged_sale) references sale (goen_id),
    add constraint foreign key (goen_belonged_item) references item (goen_id),
    add constraint foreign key (goen_contained_sales_line) references sale (goen_id)
;
