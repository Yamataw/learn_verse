CREATE INDEX IF NOT EXISTS res_collection_idx ON resources(collection_id);
CREATE INDEX  IF NOT EXISTS res_collection_created_at_idx ON resource_collections(created_at);
CREATE INDEX  IF NOT EXISTS res_created_at_idx ON resources(created_at)