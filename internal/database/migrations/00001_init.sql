-- +goose Up
CREATE TABLE lists (
   id   BIGSERIAL PRIMARY KEY,
   name text  NOT NULL,
   description TEXT NOT NULL DEFAULT ''
);

CREATE TABLE tasks (
   id   BIGSERIAL PRIMARY KEY,
   list_id BIGINT NOT NULL,
   description  TEXT NOT NULL,
   created_at TIMESTAMP NOT NULL DEFAULT NOW(),
   completed_at TIMESTAMP,

   CONSTRAINT fk_list_id FOREIGN KEY(list_id) REFERENCES lists(id) ON DELETE CASCADE
);
