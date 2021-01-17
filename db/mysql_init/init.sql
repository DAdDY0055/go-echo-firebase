CREATE TABLE users (
    id CHAR(40) NOT NULL ,
    name VARCHAR(64) NOT NULL,
    email VARCHAR(255) NOT NULL,
    firebase_auth_id VARCHAR(255) NOT NULL,
    PRIMARY KEY (id)
);

-- CREATE TABLE users (
--     id INT(11) AUTO_INCREMENT NOT NULL,
--     name VARCHAR(64) NOT NULL,
--     PRIMARY KEY (id));
-- );
