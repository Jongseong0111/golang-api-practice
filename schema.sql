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
    reg_datetime datetime(6) null,
    constraint scm_product_ibfk_1 foreign key (manufacturer_id) references scm_account (user_id),
    unique (product_unit)
) charset = utf8mb4;