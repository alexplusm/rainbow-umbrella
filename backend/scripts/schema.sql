CREATE TABLE IF NOT EXISTS users (
    user_id           SERIAL,
    login             VARCHAR(32)   NOT NULL UNIQUE,
    hashed_password   VARCHAR(255)  NOT NULL,

    first_name        VARCHAR(255)  NOT NULL,
    last_name         VARCHAR(255)  NOT NULL,
    birthday          DATE          NOT NULL,
    gender            VARCHAR(255)  NOT NULL,
    city              VARCHAR(255)  NOT NULL,

    created_at        TIMESTAMP     NOT NULL,

    --      "interests"             ... MANY TO MANY ?
    --      "friends"               ... MANY TO MANY ?
    --      "friendship_requests"   ... MANY TO MANY ? // // // // //

    PRIMARY KEY (user_id)
);

--      "interests"             ... MANY TO MANY ?
--      "friends"               ... MANY TO MANY ?
--      "friendship_requests"   ... MANY TO MANY ?
