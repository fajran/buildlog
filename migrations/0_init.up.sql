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

	type VARCHAR(50) NOT NULL,

	content_type VARCHAR(100) NOT NULL,
	content_type_parameter VARCHAR(200) NOT NULL,

	size BIGINT
);

CREATE INDEX logs_idx_type ON logs (type);
CREATE INDEX logs_idx_build_id ON logs (build_id);

