CREATE TABLE gw.restaurants
(
    id     INT         NOT NULL AUTO_INCREMENT,
    name   VARCHAR(64) NOT NULL,
    login  VARCHAR(64) NOT NULL,
    tg_key VARCHAR(64) NOT NULL,

    PRIMARY KEY (id)
);

ALTER TABLE gw.restaurants
    AUTO_INCREMENT = 30001;