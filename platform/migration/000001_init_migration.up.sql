CREATE TABLE employees (
   id INT primary key,
   first_name TEXT not null,
   last_name TEXT,
   created_at TIMESTAMP default now()
);

