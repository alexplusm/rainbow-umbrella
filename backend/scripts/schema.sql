CREATE TABLE IF NOT EXISTS users (
    "user_id"           VARCHAR NOT NULL,
    "login"             VARCHAR NOT NULL,
    "hashed_password"   VARCHAR NOT NULL,

    "first_name"                VARCHAR     NOT NULL,
    "last_name"                 VARCHAR     NOT NULL,
    "age"                       INT         NOT NULL,
    "city"                      VARCHAR     NOT NULL,

--     "interests" ... MANY TO MAY ?
    "created_at"        VARCHAR NOT NULL,

    PRIMARY KEY ("user_id")
);