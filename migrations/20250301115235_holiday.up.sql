CREATE TABLE IF NOT EXISTS holidays(
                                       id SERIAL PRIMARY KEY,
                                       text TEXT NOT NULL,
                                       author VARCHAR(100) NOT NULL DEFAULT 'Anonymus',
                                       telegram_name VARCHAR(150) NOT NULL UNIQUE CHECK (telegram_name IN (
                                                                                                           'banan4ik_kmy','glavadetdoma', 'heriofOmen', 'saniysotkuverni', 'xapsiel', 'zavsegdat1y',
                                                                                                           'r7821699990', 'metelykx', 'whysirnik', 'squerizz', 'f1rer0se',
                                                                                                           'timedust11', 'dsedulya', 'ssvivss', 'beo_blod', 'quinsone425',
                                                                                                           'why_n00tt', 'bruhrayz', 'letai_kak_ya', 'selimqwe', 'belyasha82', 'myatiy_lox'
                                           ))
);
