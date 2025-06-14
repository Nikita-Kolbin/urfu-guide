CREATE TABLE languages (
    code TEXT PRIMARY KEY
);

CREATE TABLE sections (
    id BIGSERIAL PRIMARY KEY,
    language_code TEXT REFERENCES languages(code) ON DELETE CASCADE,
    content_type TEXT NOT NULL,
    title TEXT NOT NULL,
    icon TEXT NOT NULL,
    position INT NOT NULL
);

CREATE TABLE default_blocks (
    id BIGSERIAL PRIMARY KEY,
    section_id BIGINT REFERENCES sections(id) ON DELETE CASCADE,
    address TEXT NOT NULL,
    timetable TEXT NOT NULL,
    phone TEXT NOT NULL,
    description TEXT NOT NULL,
    link TEXT NOT NULL,
    image TEXT NOT NULL,
    position INT NOT NULL -- для типа default_block_list
);

CREATE TABLE drop_down_blocks (
    id BIGSERIAL PRIMARY KEY,
    section_id BIGINT REFERENCES sections(id) ON DELETE CASCADE,
    title TEXT NOT NULL,
    sub_title TEXT NOT NULL,
    icon TEXT NOT NULL,
    position INT NOT NULL -- для типа drop_down_list_list
);

CREATE TABLE drop_down_elements (
    id BIGSERIAL PRIMARY KEY,
    block_id BIGINT REFERENCES drop_down_blocks(id) ON DELETE CASCADE,
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    link TEXT NOT NULL,
    position INT NOT NULL
);




-- TODO: тестовые данные, убрать --

-- Добавить языки
INSERT INTO languages (code)
VALUES
    ('ru'),
    ('en');



-- Добавить дефолтую секцию
INSERT INTO sections (language_code, content_type, title, icon, position)
VALUES
    ('ru', 'default_block', 'Секция 1', 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcThg82X5FYawxzhUZBeFRy2xo4TQB_8FjR7Wg&s', 1),
    ('en', 'default_block', 'Section 1', 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcThg82X5FYawxzhUZBeFRy2xo4TQB_8FjR7Wg&s', 1);

-- Добавить блоки (предполагается, что ID секций идут по порядку: 1–6)
INSERT INTO default_blocks (section_id, address, timetable, phone, description, link, image, position)
VALUES
    -- Russian
    (1, 'ул. Ленина, 1', 'Пн-Пт 9:00–18:00', '+7 900 000-00-01', 'Описание блока 1', 'https://urfu.ru/ru/', 'https://static.wikia.nocookie.net/silly-cat/images/3/35/CuhBG.png/revision/latest?cb=20231025181331', 1),
    -- English
    (2, 'Lenina St, 1', 'Mon-Fri 9:00–18:00', '+1 900 000-00-01', 'Block description 1', 'https://urfu.ru/en/', 'https://static.wikia.nocookie.net/silly-cat/images/3/35/CuhBG.png/revision/latest?cb=20231025181331', 1);



-- Секция список дроп даунов
INSERT INTO sections (language_code, content_type, title, icon, position)
VALUES
    ('ru', 'drop_down_list', 'Секция 2', 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcThg82X5FYawxzhUZBeFRy2xo4TQB_8FjR7Wg&s',2),
    ('en', 'drop_down_list', 'Section 2', 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcThg82X5FYawxzhUZBeFRy2xo4TQB_8FjR7Wg&s',2);

-- Добавить drop_down_blocks
INSERT INTO drop_down_blocks (section_id, title, sub_title, icon, position)
VALUES
    -- Russian block
    (3, 'Основные программы (бюджет)', 'Для студентов, обучающихся на основных программах на бюджетной основе', 'document', 1),
    -- English block
    (4, 'Main programs (budget)', 'For students enrolled in main programs on a budget basis', 'document', 1);

-- Добавить элементы в русский блок (block_id = 1)
INSERT INTO drop_down_elements (block_id, title, description, link, position)
VALUES
    (1, 'Прибытие', 'Для того, чтобы наши волонтёры встретили тебя в аэропорту, заполни анкету прибытия.', 'https://istudent.urfu.ru/', 1),
    (1, 'Сдать документы для зачисления', 'Чтобы начать процедуру зачисления, сдай необходимые документы в Центр приёма (ГУК-109).', '', 2),
    (1, 'Заселение в общежитие', 'Адрес твоего общежития указан в договоре', 'https://istudent.urfu.ru/s/modeus', 3);

-- Добавить элементы в английский блок (block_id = 2)
INSERT INTO drop_down_elements (block_id, title, description, link, position)
VALUES
    (2, 'Arrival', 'To be met by our volunteers at the airport, fill out the arrival form.', 'https://istudent.urfu.ru/', 1),
    (2, 'Submit documents for enrollment', 'To start the enrollment procedure, submit the required documents to the Admission Center (GUK-109).', '', 2),
    (2, 'Dormitory check-in', 'Your dormitory address is stated in the contract.', 'https://istudent.urfu.ru/s/modeus', 3);



-- Добавить секцию список дефолтных секций
INSERT INTO sections (language_code, content_type, title, icon, position)
VALUES
    ('ru', 'default_block_list', 'Секция 3', 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcThg82X5FYawxzhUZBeFRy2xo4TQB_8FjR7Wg&s', 3),
    ('en', 'default_block_list', 'Section 3', 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcThg82X5FYawxzhUZBeFRy2xo4TQB_8FjR7Wg&s', 3);

-- Добавить блоки (предполагается, что ID секций идут по порядку: 1–6)
INSERT INTO default_blocks (section_id, address, timetable, phone, description, link, image, position)
VALUES
    -- Russian
    (5, 'ул. Ленина, 1', 'Пн-Пт 9:00–18:00', '+7 900 000-00-01', 'Описание блока 1', 'https://urfu.ru/ru/', 'https://static.wikia.nocookie.net/silly-cat/images/3/35/CuhBG.png/revision/latest?cb=20231025181331', 1),
    (5, 'ул. Ленина, 2', 'Пн-Пт 9:00–18:00', '+7 900 000-00-01', 'Описание блока 2', 'https://urfu.ru/ru/', 'https://static.wikia.nocookie.net/silly-cat/images/3/35/CuhBG.png/revision/latest?cb=20231025181331', 2),
    (5, 'ул. Ленина, 3', 'Пн-Пт 9:00–18:00', '+7 900 000-00-01', 'Описание блока 3', 'https://urfu.ru/ru/', 'https://static.wikia.nocookie.net/silly-cat/images/3/35/CuhBG.png/revision/latest?cb=20231025181331', 3),
    -- English
    (6, 'Lenina St, 1', 'Mon-Fri 9:00–18:00', '+1 900 000-00-01', 'Block description 1', 'https://urfu.ru/en/', 'https://static.wikia.nocookie.net/silly-cat/images/3/35/CuhBG.png/revision/latest?cb=20231025181331', 1),
    (6, 'Lenina St, 2', 'Mon-Fri 9:00–18:00', '+1 900 000-00-01', 'Block description 2', 'https://urfu.ru/en/', 'https://static.wikia.nocookie.net/silly-cat/images/3/35/CuhBG.png/revision/latest?cb=20231025181331', 2),
    (6, 'Lenina St, 3', 'Mon-Fri 9:00–18:00', '+1 900 000-00-01', 'Block description 3', 'https://urfu.ru/en/', 'https://static.wikia.nocookie.net/silly-cat/images/3/35/CuhBG.png/revision/latest?cb=20231025181331', 3);



-- Секция секцию список списков дроп даунов
INSERT INTO sections (language_code, content_type, title, icon, position)
VALUES
    ('ru', 'drop_down_list_list', 'Секция 4', 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcThg82X5FYawxzhUZBeFRy2xo4TQB_8FjR7Wg&s',4),
    ('en', 'drop_down_list_list', 'Section 4', 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcThg82X5FYawxzhUZBeFRy2xo4TQB_8FjR7Wg&s',4);

-- Добавить drop_down_blocks
INSERT INTO drop_down_blocks (section_id, title, sub_title, icon, position)
VALUES
    -- Russian block
    (7, 'Основные программы (бюджет)', 'Для студентов, обучающихся на основных программах на бюджетной основе', 'document', 1),
    (7, 'Основные программы (бюджет)', 'Для студентов, обучающихся на основных программах на бюджетной основе', 'document', 2),
    -- English block
    (8, 'Main programs (budget)', 'For students enrolled in main programs on a budget basis', 'document', 1),
    (8, 'Main programs (budget)', 'For students enrolled in main programs on a budget basis', 'document', 2);

-- Добавить элементы в русский блок (block_id = 1)
INSERT INTO drop_down_elements (block_id, title, description, link, position)
VALUES
    (3, 'Прибытие', 'Для того, чтобы наши волонтёры встретили тебя в аэропорту, заполни анкету прибытия.', 'https://istudent.urfu.ru/', 1),
    (3, 'Сдать документы для зачисления', 'Чтобы начать процедуру зачисления, сдай необходимые документы в Центр приёма (ГУК-109).', '', 2),
    (4, 'Заселение в общежитие', 'Адрес твоего общежития указан в договоре', 'https://istudent.urfu.ru/s/modeus', 1),
    (4, 'Заселение в общежитие', 'Адрес твоего общежития указан в договоре', 'https://istudent.urfu.ru/s/modeus', 2);

-- Добавить элементы в английский блок (block_id = 2)
INSERT INTO drop_down_elements (block_id, title, description, link, position)
VALUES
    (5, 'Arrival', 'To be met by our volunteers at the airport, fill out the arrival form.', 'https://istudent.urfu.ru/', 1),
    (5, 'Submit documents for enrollment', 'To start the enrollment procedure, submit the required documents to the Admission Center (GUK-109).', '', 2),
    (6, 'Dormitory check-in', 'Your dormitory address is stated in the contract.', 'https://istudent.urfu.ru/s/modeus', 1),
    (6, 'Dormitory check-in', 'Your dormitory address is stated in the contract.', 'https://istudent.urfu.ru/s/modeus', 2);


