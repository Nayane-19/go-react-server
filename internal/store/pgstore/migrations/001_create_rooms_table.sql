CREATE TABLE IF NOT EXISTS teste (
    "id"    uuid            PRIMARY KEY     NOT NULL    DEFAULT gen_random_uuid(),
    "theme" VARCHAR(255)                    NOT NULL
);

---- create above / drop below ----

DROP TABLE IF EXISTS rooms;