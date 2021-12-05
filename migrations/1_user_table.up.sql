CREATE TABLE users  (
    id UUID DEFAULT  uuid_generate_v4(),
    first_name varchar(40) NOT NULL,
    last_name varchar(40) NOT NULL,
    email varchar(40) PRIMARY KEY
)