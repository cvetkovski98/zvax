DROP TABLE IF EXISTS keys;
CREATE TABLE keys (
    id serial PRIMARY KEY,
    holder varchar(255) NOT NULL UNIQUE,
    affiliation varchar(255) NOT NULL,
    value text NOT NULL
);
