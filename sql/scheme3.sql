create table item(
                     goen_id int primary key ,
                     barcode int,
                     name varchar(255),
                     price float,
                     stock_number int,
                     order_price float ,
                     belonged_item_goen_id int
);
create table item_contained_item (
                                     goen_id_1 int,
                                     goen_id_2 int,
                                     primary key (goen_id_1, goen_id_2)
);
alter table item_contained_item
    add constraint foreign key (goen_id_1) references item (goen_id) on delete cascade ,
    add constraint foreign key (goen_id_2) references item (goen_id) on delete cascade ;
alter table item
    add constraint foreign key  (belonged_item_goen_id) references item(goen_id) on delete set null;