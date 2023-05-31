CREATE TABLE pusher
(
    ID    UUID DEFAULT gen_random_uuid()  PRIMARY KEY,
    NAME  TEXT UNIQUE NOT NULL,
    EMAIL TEXT
);

CREATE TABLE sender
(
    ID         UUID DEFAULT gen_random_uuid()  PRIMARY KEY,
    LOGIN      TEXT UNIQUE NOT NULL,
    HTML_URL   TEXT        NOT NULL,
    AVATAR_URL TEXT        NOT NULL,
    URL        TEXT        NOT NULL
);

CREATE TABLE push
(
    ID            UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    BEFORE_COMMIT TEXT,
    AFTER_COMMIT  TEXT,
    PUSHER_ID     UUID REFERENCES pusher (ID),
    SENDER_ID     UUID REFERENCES sender (ID)
);
