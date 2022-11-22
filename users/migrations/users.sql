DROP TABLE IF EXISTS users;
CREATE TABLE users
(
    id    serial PRIMARY KEY,
    name  varchar(255),
    email varchar(255) NOT NULL,
    phone varchar(255)
);
