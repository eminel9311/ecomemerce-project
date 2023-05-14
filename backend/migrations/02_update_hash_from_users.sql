-- +migrate Up
ALTER TABLE users
ADD COLUMN hash varchar(255);