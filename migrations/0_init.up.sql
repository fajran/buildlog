CREATE TABLE buildlog (
  id SERIAL,
  created TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),

  key VARCHAR(100),
  name VARCHAR(100),
  status VARCHAR(20),
  started TIMESTAMP WITH TIME ZONE,
  finished TIMESTAMP WITH TIME ZONE
);

CREATE INDEX buildlog_idx_key ON buildlog (key);

