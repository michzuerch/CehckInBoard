
-- +migrate Up
CREATE TABLE "persons" (
  "id" BIGSERIAL PRIMARY KEY,
  "firstname" VARCHAR(50) NOT NULL,
  "lastname" VARCHAR(50) NOT NULL,
  "birthdate" DATE, 
  "emailaddress" VARCHAR(100) UNIQUE
);

CREATE TABLE "stamps" (
  "id" BIGSERIAL PRIMARY KEY,
  "person_id" BIGINT,
  "checkin" BOOLEAN NOT NULL DEFAULT FALSE,
  "stamp" TIMESTAMP NOT NULL,
  FOREIGN KEY ("person_id") REFERENCES "persons" ("id")
);

-- ALTER TABLE "stamps" ADD FOREIGN KEY ("person_id") REFERENCES "persons" ("id");

CREATE INDEX ON "persons" ("emailaddress");

COMMENT ON TABLE "stamps" IS 'Timestamps for persons';

-- +migrate Down
DROP TABLE "stamps";
DROP TABLE "persons";
