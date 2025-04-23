DROP TABLE IF EXISTS resources;
DROP TABLE IF EXISTS resource_collections;
DROP TYPE IF EXISTS resource_type;
CREATE TYPE resource_type AS ENUM ('note','flashcard','quiz','file');

CREATE EXTENSION IF NOT EXISTS ulid;

CREATE TABLE resource_collections (
                                      id          ULID PRIMARY KEY DEFAULT gen_ulid(),
                                      name        TEXT NOT NULL,
                                      description TEXT,
                                      created_at  TIMESTAMPTZ DEFAULT now(),
                                      updated_at    TIMESTAMPTZ DEFAULT now(),
                                      deleted_at    TIMESTAMPTZ DEFAULT null
);
CREATE TABLE resources (
                           id            ULID PRIMARY KEY DEFAULT gen_ulid(),
                           collection_id ULID REFERENCES resource_collections(id) ON DELETE SET NULL,
                           type          resource_type NOT NULL,
                           title         TEXT NOT NULL,
                           content       JSONB NOT NULL,
                           metadata      JSONB,
                           created_at    TIMESTAMPTZ DEFAULT now(),
                           updated_at    TIMESTAMPTZ DEFAULT now(),
                           deleted_at    TIMESTAMPTZ DEFAULT null

);

CREATE INDEX IF NOT EXISTS res_collection_idx ON resources(collection_id);
CREATE INDEX  IF NOT EXISTS res_collection_created_at_idx ON resource_collections(created_at);
CREATE INDEX  IF NOT EXISTS res_created_at_idx ON resources(created_at)

