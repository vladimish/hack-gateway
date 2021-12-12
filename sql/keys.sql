CREATE TABLE gw.keys
(
    id     INT NOT NULL AUTO_INCREMENT,
    tg_key VARCHAR(64),

    PRIMARY KEY (id)
);

ALTER TABLE gw.`keys`
    AUTO_INCREMENT = 30001;