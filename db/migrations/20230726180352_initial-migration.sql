-- +goose Up
-- create "resource_providers" table
CREATE TABLE "resource_providers" ("id" character varying NOT NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "name" text NOT NULL, "description" text NULL, "owner_id" character varying NOT NULL, PRIMARY KEY ("id"));
-- create index "resourceprovider_created_at" to table: "resource_providers"
CREATE INDEX "resourceprovider_created_at" ON "resource_providers" ("created_at");
-- create index "resourceprovider_owner_id" to table: "resource_providers"
CREATE INDEX "resourceprovider_owner_id" ON "resource_providers" ("owner_id");
-- create index "resourceprovider_updated_at" to table: "resource_providers"
CREATE INDEX "resourceprovider_updated_at" ON "resource_providers" ("updated_at");

-- +goose Down
-- reverse: create index "resourceprovider_updated_at" to table: "resource_providers"
DROP INDEX "resourceprovider_updated_at";
-- reverse: create index "resourceprovider_owner_id" to table: "resource_providers"
DROP INDEX "resourceprovider_owner_id";
-- reverse: create index "resourceprovider_created_at" to table: "resource_providers"
DROP INDEX "resourceprovider_created_at";
-- reverse: create "resource_providers" table
DROP TABLE "resource_providers";
