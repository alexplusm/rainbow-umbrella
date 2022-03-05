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

    PRIMARY KEY (user_id)
);

CREATE TABLE IF NOT EXISTS `friendships` (
    `friendship_id`         SERIAL,
    `requesting_user_id`    INT         NOT NULL,
    `targeting_user_id`     INT         NOT NULL,
    `status`                VARCHAR     NOT NULL,

    PRIMARY KEY (`friendship_id`),

    CONSTRAINT `unique_users` UNIQUE(`requesting_user_id`, `targeting_user_id`),

    FOREIGN KEY (`requesting_user_id`),
    FOREIGN KEY (`targeting_user_id`)
);

--      "interests"             ... MANY TO MANY ?
