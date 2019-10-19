CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE IF NOT EXISTS books (id uuid DEFAULT (uuid_generate_v4()), name VARCHAR (100) NOT NULL, author VARCHAR (100))