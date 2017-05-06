CREATE TABLE builds (
  id SERIAL,
  created TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),

  key VARCHAR(100) NOT NULL
);

CREATE INDEX builds_idx_key ON builds (key);


CREATE TABLE logs (
	id SERIAL,
	build_id INT NOT NULL,
	created TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),

	type VARCHAR(50),
	content_type VARCHAR(100),
	identifier VARCHAR(200),
	size BIGINT
);

CREATE INDEX logs_idx_type ON logs (type);

