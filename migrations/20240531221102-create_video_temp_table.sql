
-- +migrate Up
CREATE TABLE IF NOT EXISTS video_temp (
  id uuid PRIMARY KEY,
  content_type text not null,
  data BYTEA not null,
  data_size bigint not null, 
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT,
);
-- +migrate Down
