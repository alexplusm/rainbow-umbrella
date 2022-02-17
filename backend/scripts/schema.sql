CREATE TABLE IF NOT EXISTS users (
    "user_id"           VARCHAR NOT NULL, -- TODO: use BIG SERIAL ?
    "login"             VARCHAR NOT NULL, -- TODO: UNIQUE CONSTRAINT
    "hashed_password"   VARCHAR NOT NULL,

    "first_name"        VARCHAR     NOT NULL,
    "last_name"         VARCHAR     NOT NULL,
    "birthday"          DATE        NOT NULL,
    "gender"            VARCHAR     NOT NULL,
    "city"              VARCHAR     NOT NULL,

--      "interests"             ... MANY TO MANY ?
--      "friends"               ... MANY TO MANY ?
--      "friendship_requests"   ... MANY TO MANY ?

    "created_at"        TIMESTAMP NOT NULL,

    PRIMARY KEY ("user_id")
);

--      "interests"             ... MANY TO MANY ?
--      "friends"               ... MANY TO MANY ?
--      "friendship_requests"   ... MANY TO MANY ?
