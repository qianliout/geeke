use stack;
create table name_codes
(
    id          bigint      not null auto_increment primary key,
    name        varchar(30) not null default '',
    code        varchar(10) not null default '',
    profile     bigint      not null default 0,
    cash_flow   bigint      not null default 0,
    balance     bigint      not null default 0,
    stock_price float       not null default 0,
    crawl_date  bigint      not null default 0,
    shsz        varchar(10) not null default '',
    unique key ` idx_code ` (` code `)
);
create table profiles
(
    id                 bigint not null auto_increment primary key,
    name               varchar(100),
    code               varchar(10),
    reporting_period   varchar(20),
    operate_all_income bigint,
    operate_income     bigint,
    operate_all_cost   bigint,
    operate_cost       bigint,
    tax                bigint,
    sales_expense      bigint,
    manage_expense     bigint,
    financial_expense  bigint,
    rd_expense         bigint,
    diluted_earn       float,
    unique key ` idx_code ` (` code `, ` reporting_period `)
);

create table balances
(
    id               bigint not null auto_increment primary key,
    name             varchar(20),
    code             varchar(10),
    reporting_period varchar(50),
    money_funds      bigint,
    trans_finance    bigint,
    stock            bigint,
    short_loan       bigint,
    long_loan        bigint,
    capital          bigint,
    unique key ` idx_code ` (` code `, ` reporting_period `)
);


