package database

func TablesSetup() []string {
	
	Uuid := `CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`


	EchoPost := `CREATE TABLE IF NOT EXISTS echo_post(
		id uuid DEFAULT uuid_generate_v4(),
		content TEXT NULL,
		posted_at TIMESTAMP NULL,
		created_at TIMESTAMP DEFAULT NOW(),
		draft BOOLEAN DEFAULT TRUE,
		PRIMARY KEY (id)
	)`;

	Images := `CREATE TABLE IF NOT EXISTS Images(
		id SERIAL PRIMARY KEY,
		url TEXT NOT NULL,
		created_at TIMESTAMP DEFAULT NOW()
	)
	`

	SubmissionTokens := `CREATE TABLE IF NOT EXISTS SubmissionTokens(
		id uuid DEFAULT uuid_generate_v4(),
		submitted BOOLEAN DEFAULT FALSE
	)`

	User := `CREATE TABLE IF NOT EXISTS EchoUser(
		id SERIAL PRIMARY KEY,
		username TEXT NOT NULL,
		password TEXT NOT NULL,
		post_passcode smallint not null check (post_passcode >= 1000 and post_passcode <= 9999),
		created_at TIMESTAMP DEFAULT NOW()
	)`

	



return []string{Uuid, EchoPost, Images, SubmissionTokens, User}

}