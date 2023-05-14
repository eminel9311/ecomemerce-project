
-- +migrate Up
CREATE TABLE "users" (
     "id" varchar(255) PRIMARY KEY,
     "email" varchar(255) UNIQUE NOT NULL,
     "phone" varchar(25) UNIQUE NOT NULL,
     "password" varchar(255) NOT NULL,
     "address" varchar(255),
     "fullname" varchar(150),
     "avatar" text,
     "role" varchar(25),
     "created_at" timestamp NOT NULL,
     "updated_at" timestamp NOT NULL
);

-- +migrate Down
DROP TABLE users;