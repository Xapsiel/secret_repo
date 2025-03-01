

CREATE TABLE IF NOT EXISTS holidays(
    id SERIAL PRIMARY KEY ,
    text text NOT NULL ,
    author varchar(100) NOT NULL DEFAULT 'Anonymus' ,
    telegram_name  varchar(150) unique not null
);