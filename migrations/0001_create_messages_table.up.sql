CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS messages (
  ID uuid,
  Content text,
  ToPhoneNumber varchar
);
