CREATE TABLE IF NOT EXISTS holidays(
                                       id SERIAL PRIMARY KEY,
                                       text TEXT NOT NULL,
                                       author VARCHAR(100) NOT NULL DEFAULT 'Anonymus',
                                       telegram_name VARCHAR(150) NOT NULL UNIQUE CHECK (telegram_name IN (
                                                                                                           'Banan4ik_kmy', 'HeriofOmen', 'SaniySotkuVerni', 'Xapsiel', 'Zavsegdat1y',
                                                                                                           'R7821699990', 'Metelykx', 'Whysirnik', 'Squerizz', 'F1rer0se',
                                                                                                           'Timedust11', 'Dsedulya', 'Ssvivss', 'Beo_blod', 'Quinsone425',
                                                                                                           'Why_n00tt', 'Bruhrayz', 'Letai_kak_ya', 'Selimqwe', 'Belyasha82', 'Myatiy_lox'
                                           ))
);
