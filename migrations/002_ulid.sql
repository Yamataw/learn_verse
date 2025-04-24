DROP TABLE IF EXISTS resources;
DROP TABLE IF EXISTS resource_collections;
DROP TYPE IF EXISTS resource_type;

CREATE EXTENSION IF NOT EXISTS ulid;

CREATE TABLE resource_collections (
                                      id          ULID PRIMARY KEY DEFAULT gen_ulid(),
                                      name        TEXT NOT NULL,
                                      description TEXT,
                                      created_at  TIMESTAMPTZ DEFAULT now(),
                                      updated_at    TIMESTAMPTZ DEFAULT now(),
                                      deleted_at    TIMESTAMPTZ DEFAULT null
);

CREATE TABLE note (
                           id            ULID PRIMARY KEY DEFAULT gen_ulid(),
                           collection_id ULID REFERENCES resource_collections(id) ON DELETE SET NULL,
                           title         TEXT NOT NULL,
                           theme         TEXT NOT NULL,
                           content       JSONB NOT NULL,
                           metadata      JSONB,
                           created_at    TIMESTAMPTZ DEFAULT now(),
                           updated_at    TIMESTAMPTZ DEFAULT now(),
                           deleted_at    TIMESTAMPTZ DEFAULT null

);

CREATE TABLE flashcard (
                      id            ULID PRIMARY KEY DEFAULT gen_ulid(),
                      collection_id ULID REFERENCES resource_collections(id) ON DELETE SET NULL,
                      title         TEXT NOT NULL,
                      theme         TEXT NOT NULL,
                      content       JSONB NOT NULL,
                      metadata      JSONB,
                      created_at    TIMESTAMPTZ DEFAULT now(),
                      updated_at    TIMESTAMPTZ DEFAULT now(),
                      deleted_at    TIMESTAMPTZ DEFAULT null

);

CREATE TABLE quiz (
                      id            ULID PRIMARY KEY DEFAULT gen_ulid(),
                      collection_id ULID REFERENCES resource_collections(id) ON DELETE SET NULL,
                      title         TEXT NOT NULL,
                      theme         TEXT NOT NULL,
                      content       JSONB NOT NULL,
                      num_questions INTEGER,
                      metadata      JSONB,
                      created_at    TIMESTAMPTZ DEFAULT now(),
                      updated_at    TIMESTAMPTZ DEFAULT now(),
                      deleted_at    TIMESTAMPTZ DEFAULT null

);

CREATE TABLE file (
                      id            ULID PRIMARY KEY DEFAULT gen_ulid(),
                      collection_id ULID REFERENCES resource_collections(id) ON DELETE SET NULL,
                      title         TEXT NOT NULL,
                      theme         TEXT NOT NULL,
                      file_type     TEXT NOT NULL,
                      -- file_type     TEXT NOT NULL CHECK (file_type IN ('pdf', 'docx', 'png', 'jpg', 'txt')),
                      content       JSONB NOT NULL,
                      metadata      JSONB,
                      created_at    TIMESTAMPTZ DEFAULT now(),
                      updated_at    TIMESTAMPTZ DEFAULT now(),
                      deleted_at    TIMESTAMPTZ DEFAULT null

);

CREATE INDEX IF NOT EXISTS res_collection_created_at_idx ON resource_collections(created_at);
CREATE INDEX IF NOT EXISTS res_collection_active_at_idx ON resource_collections(id) WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS note_collection_idx ON note(collection_id);
CREATE INDEX IF NOT EXISTS note_created_at_idx ON note(created_at);
CREATE INDEX IF NOT EXISTS note_active_idx ON note(id) WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS note_collection_created ON note (collection_id, deleted_at, created_at DESC);
CREATE INDEX IF NOT EXISTS flashcard_collection_idx ON flashcard(collection_id);
CREATE INDEX IF NOT EXISTS flashcard_created_at_idx ON flashcard(created_at);
CREATE INDEX IF NOT EXISTS flashcard_active_idx ON flashcard(id) WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS flashcard_collection_created ON flashcard (collection_id, deleted_at, created_at DESC);
CREATE INDEX IF NOT EXISTS quiz_collection_idx ON quiz(collection_id);
CREATE INDEX IF NOT EXISTS quiz_created_at_idx ON quiz(created_at);
CREATE INDEX IF NOT EXISTS quiz_active_idx ON quiz(id) WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS quiz_collection_created ON quiz (collection_id, deleted_at, created_at DESC);
CREATE INDEX IF NOT EXISTS file_collection_idx ON file(collection_id);
CREATE INDEX IF NOT EXISTS file_created_at_idx ON file(created_at);
CREATE INDEX IF NOT EXISTS file_active_idx ON file(id) WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS file_collection_created ON file (collection_id, deleted_at, created_at DESC);
