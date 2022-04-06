CREATE TABLE IF NOT EXISTS `users` (
    `user_id`           SERIAL,
    `login`             VARCHAR(32)   NOT NULL UNIQUE,
    `hashed_password`   VARCHAR(255)  NOT NULL,

    `first_name`        VARCHAR(255)  NOT NULL,
    `last_name`         VARCHAR(255)  NOT NULL,
    `birthday`          DATE          NOT NULL,
    `gender`            VARCHAR(255)  NOT NULL,
    `city`              VARCHAR(255)  NOT NULL,

    `created_at`        TIMESTAMP     NOT NULL,

    PRIMARY KEY (`user_id`)
);

CREATE TABLE IF NOT EXISTS `friendships` (
    `friendship_id`         SERIAL,
    `requesting_user_id`    BIGINT UNSIGNED   NOT NULL,
    `targeting_user_id`     BIGINT UNSIGNED   NOT NULL,
    `status`                VARCHAR(255)      NOT NULL,

    `created_at`        TIMESTAMP     NOT NULL,
    `updated_at`        TIMESTAMP,

    CONSTRAINT `unique_users` UNIQUE(`requesting_user_id`, `targeting_user_id`),

    FOREIGN KEY (`requesting_user_id`)    REFERENCES `users`(`user_id`),
    FOREIGN KEY (`targeting_user_id`)     REFERENCES `users`(`user_id`),

    PRIMARY KEY (`friendship_id`)
);

CREATE TABLE IF NOT EXISTS `interests` (
    `interest_id`   SERIAL,
    `value`         VARCHAR(255) NOT NULL UNIQUE,

    PRIMARY KEY (`interest_id`)
);

CREATE TABLE IF NOT EXISTS `user_interests` (
    `user_interest_id`  SERIAL,
    `user_id`           BIGINT UNSIGNED   NOT NULL,
    `interest_id`       BIGINT UNSIGNED   NOT NULL,

    FOREIGN KEY (`user_id`)         REFERENCES `users`(`user_id`),
    FOREIGN KEY (`interest_id`)     REFERENCES `interests`(`interest_id`),

    PRIMARY KEY (`user_interest_id`)
);
