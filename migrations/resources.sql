DROP TABLE IF EXISTS resources;
DROP TABLE IF EXISTS resource_collections;
DROP TYPE IF EXISTS resource_type;
CREATE TYPE resource_type AS ENUM ('note','flashcard','quiz','file');

CREATE EXTENSION IF NOT EXISTS pgcrypto;      -- gen_random_uuid()

CREATE TABLE resource_collections (
                                      id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                                      name        TEXT NOT NULL,
                                      description TEXT,
                                      created_at  TIMESTAMPTZ DEFAULT now(),
                                      updated_at    TIMESTAMPTZ DEFAULT now(),
                                      deleted_at    TIMESTAMPTZ DEFAULT null    
);
CREATE TABLE resources (
                           id            UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                           collection_id UUID REFERENCES resource_collections(id) ON DELETE SET NULL,
                           type          resource_type NOT NULL,
                           title         TEXT NOT NULL,
                           content       JSONB NOT NULL,
                           metadata      JSONB,
                           created_at    TIMESTAMPTZ DEFAULT now(),
                           updated_at    TIMESTAMPTZ DEFAULT now(),
                           deleted_at    TIMESTAMPTZ DEFAULT null

);

