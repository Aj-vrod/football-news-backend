CREATE TABLE IF NOT EXISTS News(
	news_id serial PRIMARY KEY,
	title VARCHAR(50) NOT NULL,
	link VARCHAR(50) NOT NULL,
	source VARCHAR(50) NOT NULL,
	creation_date VARCHAR(20) NOT NULL,
	date_id INT NOT NULL,
	FOREIGN KEY (date_id)
  	REFERENCES Dates(date_id)
);
