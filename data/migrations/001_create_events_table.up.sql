CREATE TABLE events (
    seq         bigserial PRIMARY KEY NOT NULL,
    type        varchar(40) NOT NULL,
    data        bytea NOT NULL
);
