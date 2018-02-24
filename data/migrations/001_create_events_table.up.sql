CREATE TABLE events (
    seq         bigserial PRIMARY KEY NOT NULL,
    type        varchar(50) NOT NULL,
    data        bytea NOT NULL
);
