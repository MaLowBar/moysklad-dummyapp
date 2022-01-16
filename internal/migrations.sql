create table baseapp
(
    id           serial,
    app_id       varchar,
    app_uid      varchar,
    account_id   varchar,
    status       varchar,
    access_token varchar,
    secret_key   varchar
);

create table dummyapp
(
    baseapp_id   integer,
    info_message varchar,
    store        varchar
);

