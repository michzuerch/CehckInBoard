
-- +migrate Up
CREATE TABLE person (id int);
-- +migrate Down
DROP TABLE person;
