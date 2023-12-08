# Generate sample data in SQL using LLM

1. Follow steps from [day before](../05-sample-sql-data/README.md) to get `OPENAI_SECRET_KEY`.
2. Generate some table schema such as:
```
echo "CREATE TABLE IF NOT EXISTS users (
	id TEXT PRIMARY KEY,
	first_name TEXT,
	last_name TEXT,
  email TEXT,
  creation_date TEXT,
  last_login TEXT,
  status TEXT)" | sqlite3 "sample.db"
  
```
You may also add few records to it.

For example:
```
echo "INSERT INTO users (id, first_name, last_name, email, creation_date, last_login, status) VALUES
('1', 'Carlos', 'Hernandez', 'carlos.hernandez@example.com', '2023-01-01', '2023-01-02', 'active'),
('2', 'Maria', 'Gonzalez', 'maria.gonzalez@example.com', '2023-01-03', '2023-01-04', 'active'),
('3', 'Luis', 'Alvarez', 'luis.alvarez@example.com', '2023-01-05', '2023-01-06', 'inactive'),
('4', 'Sofia', 'Cruz', 'sofia.cruz@example.com', '2023-01-07', '2023-01-08', 'active');" | sqlite3 "sample.db"
```

3. Run it with:
```
go run generate.go
```
