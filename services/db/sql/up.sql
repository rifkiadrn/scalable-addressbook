DROP TABLE IF EXISTS addresslist;
CREATE TABLE addresslist (
  id SERIAL PRIMARY KEY,
  Name varchar(50) NOT NULL,
  Phone_number varchar(50) NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE default NOW(),
  CHECK (Name <> '' AND Phone_number <> '')
);

