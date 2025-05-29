CREATE TABLE languages (
    code TEXT PRIMARY KEY
);

CREATE TABLE sections (
    id BIGSERIAL PRIMARY KEY,
    language_code TEXT REFERENCES languages(code) ON DELETE CASCADE,
    content_type TEXT NOT NULL,
    title TEXT NOT NULL,
    position INT NOT NULL
);

CREATE TABLE default_blocks (
    id BIGSERIAL PRIMARY KEY,
    section_id BIGINT REFERENCES sections(id) ON DELETE CASCADE,
    address TEXT NOT NULL,
    timetable TEXT NOT NULL,
    phone TEXT NOT NULL,
    description TEXT NOT NULL
);

-- TODO: тестовые данные, убрать --

-- Добавить языки
INSERT INTO languages (code)
VALUES
    ('ru'),
    ('en');

-- Добавить секции на русском
INSERT INTO sections (language_code, content_type, title, position)
VALUES
    ('ru', 'default_block', 'Секция 1', 1),
    ('ru', 'default_block', 'Секция 2', 2),
    ('ru', 'default_block', 'Секция 3', 3);

-- Добавить секции на английском
INSERT INTO sections (language_code, content_type, title, position)
VALUES
    ('en', 'default_block', 'Section 1', 1),
    ('en', 'default_block', 'Section 2', 2),
    ('en', 'default_block', 'Section 3', 3);

-- Добавить блоки (предполагается, что ID секций идут по порядку: 1–6)
INSERT INTO default_blocks (section_id, address, timetable, phone, description)
VALUES
    -- Russian
    (1, 'ул. Ленина, 1', 'Пн-Пт 9:00–18:00', '+7 900 000-00-01', 'Описание блока 1'),
    (2, 'ул. Ленина, 2', 'Пн-Пт 10:00–19:00', '+7 900 000-00-02', 'Описание блока 2'),
    (3, 'ул. Ленина, 3', 'Пн-Пт 11:00–20:00', '+7 900 000-00-03', 'Описание блока 3'),

    -- English
    (4, 'Lenina St, 1', 'Mon-Fri 9:00–18:00', '+1 900 000-00-01', 'Block description 1'),
    (5, 'Lenina St, 2', 'Mon-Fri 10:00–19:00', '+1 900 000-00-02', 'Block description 2'),
    (6, 'Lenina St, 3', 'Mon-Fri 11:00–20:00', '+1 900 000-00-03', 'Block description 3');
