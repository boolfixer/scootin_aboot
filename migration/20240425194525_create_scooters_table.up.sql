USE scootin_aboot;

CREATE TABLE scooters(
    id BINARY(16) NOT NULL,
    name VARCHAR(255) NOT NULL,
    latitude SMALLINT UNSIGNED NOT NULL,
    longitude SMALLINT UNSIGNED NOT NULL,
    PRIMARY KEY(id)
) DEFAULT CHARACTER SET utf8mb4 COLLATE `utf8mb4_unicode_ci` ENGINE = InnoDB;