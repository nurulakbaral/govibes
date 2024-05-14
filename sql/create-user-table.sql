CREATE TABLE users (
	id uuid DEFAULT gen_random_uuid(),
	name VARCHAR (255) NOT NULL, 
	username VARCHAR (50) UNIQUE NOT NULL,
	email VARCHAR(50) UNIQUE NOT NULL,
	password VARCHAR (255) NOT NULL,
  	created_at TIMESTAMP NOT NULL, 
  	deleted_at TIMESTAMP,
	PRIMARY KEY (id)
);
