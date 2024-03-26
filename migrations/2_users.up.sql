CREATE TABLE t_XXXX
(
    id              SERIAL PRIMARY KEY,
    email           VARCHAR(64) NOT NULL UNIQUE,
    nickname        VARCHAR(64) NOT NULL UNIQUE,
    password_hash   VARCHAR(60) NOT NULL,
    email_confirmed TIMESTAMP NULL,
    logo            VARCHAR(255)         DEFAULT NULL,
    is_block        BOOLEAN     NOT NULL DEFAULT FALSE,
    created         TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP
);