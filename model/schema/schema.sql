create table scm_account (
    user_id int auto_increment primary key,
    user_name varchar(255) not null,
    user_account varchar(255) not null,
    user_email varchar(255) not null,
    user_password varchar(255) not null,
    unique (user_email, user_account)
) charset = utf8mb4;

create table scm_product (
    product_id int auto_increment primary key,
    product_unit varchar(255) null,
    manufacturer_id int null,
    product_name varchar(255) null,
    reg_datetime timestamp(6) not null default current_timestamp(6),
    reg_user_id int null,
    mod_datetime timestamp(6) not null default current_timestamp(6) on update current_timestamp(6),
    mod_user_id int null,
    constraint scm_product_ibfk_1 foreign key (manufacturer_id) references scm_account (user_id),
    constraint scm_product_ibfk_2 foreign key (reg_user_id) references scm_account (user_id),
    constraint scm_product_ibfk_3 foreign key (mod_user_id) references scm_account (user_id),
    unique (product_unit)
) charset = utf8mb4;

create table scm_product_size (
    row_id int auto_increment primary key,
    product_id int not null,
    size decimal(3, 1) not null,
    barcode_unit varchar(255) not null
) charset = utf8mb4;

create table scm_order (
    order_id int auto_increment primary key,
    order_date datetime(6),
    etd datetime(6),
    order_memo text,
    product_id int,
    reg_datetime timestamp(6) not null default current_timestamp(6),
    reg_user_id int,
    mod_datetime timestamp(6) not null default current_timestamp(6) on update current_timestamp(6),
    mod_user_id int,
    constraint scm_order_ibfk_1 foreign key (product_id) references scm_product (product_id),
    constraint scm_order_ibfk_2 foreign key (reg_user_id) references scm_account (user_id),
    constraint scm_order_ibfk_3 foreign key (mod_user_id) references scm_account (user_id)
) charset = utf8mb4;