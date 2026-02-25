-- Заполнение тестовыми данными для демонстрации кеширования
INSERT INTO users (name, email, age, created_at, updated_at) VALUES
('Александр Иванов', 'alexander.ivanov.1@example.com', 25, NOW(), NOW()),
('Алексей Петров', 'alexey.petrov.2@example.com', 30, NOW(), NOW()),
('Андрей Сидоров', 'andrey.sidorov.3@example.com', 28, NOW(), NOW()),
('Антон Смирнов', 'anton.smirnov.4@example.com', 35, NOW(), NOW()),
('Артем Кузнецов', 'artem.kuznetsov.5@example.com', 22, NOW(), NOW()),
('Борис Попов', 'boris.popov.6@example.com', 40, NOW(), NOW()),
('Вадим Васильев', 'vadim.vasiliev.7@example.com', 27, NOW(), NOW()),
('Валентин Соколов', 'valentin.sokolov.8@example.com', 33, NOW(), NOW()),
('Валерий Михайлов', 'valery.mikhailov.9@example.com', 29, NOW(), NOW()),
('Виктор Новиков', 'viktor.novikov.10@example.com', 31, NOW(), NOW()),
('Владимир Федоров', 'vladimir.fedorov.11@example.com', 26, NOW(), NOW()),
('Вячеслав Морозов', 'vyacheslav.morozov.12@example.com', 38, NOW(), NOW()),
('Геннадий Волков', 'gennady.volkov.13@example.com', 24, NOW(), NOW()),
('Георгий Алексеев', 'georgy.alekseev.14@example.com', 36, NOW(), NOW()),
('Григорий Лебедев', 'grigory.lebedev.15@example.com', 32, NOW(), NOW()),
('Денис Семенов', 'denis.semenov.16@example.com', 23, NOW(), NOW()),
('Дмитрий Егоров', 'dmitry.egorov.17@example.com', 37, NOW(), NOW()),
('Евгений Павлов', 'evgeny.pavlov.18@example.com', 28, NOW(), NOW()),
('Егор Козлов', 'egor.kozlov.19@example.com', 34, NOW(), NOW()),
('Иван Степанов', 'ivan.stepanov.20@example.com', 25, NOW(), NOW()),
('Игорь Николаев', 'igor.nikolaev.21@example.com', 30, NOW(), NOW()),
('Илья Орлов', 'ilya.orlov.22@example.com', 27, NOW(), NOW()),
('Кирилл Андреев', 'kirill.andreev.23@example.com', 29, NOW(), NOW()),
('Константин Макаров', 'konstantin.makarov.24@example.com', 31, NOW(), NOW()),
('Максим Никитин', 'maxim.nikitin.25@example.com', 26, NOW(), NOW()),
('Михаил Захаров', 'mikhail.zakharov.26@example.com', 35, NOW(), NOW()),
('Николай Зайцев', 'nikolay.zaitsev.27@example.com', 33, NOW(), NOW()),
('Олег Соловьев', 'oleg.soloviev.28@example.com', 28, NOW(), NOW()),
('Павел Борисов', 'pavel.borisov.29@example.com', 32, NOW(), NOW()),
('Петр Яковлев', 'petr.yakovlev.30@example.com', 24, NOW(), NOW()),
('Роман Григорьев', 'roman.grigoriev.31@example.com', 37, NOW(), NOW()),
('Сергей Романов', 'sergey.romanov.32@example.com', 29, NOW(), NOW()),
('Станислав Воробьев', 'stanislav.vorobiev.33@example.com', 31, NOW(), NOW()),
('Степан Сорокин', 'stepan.sorokin.34@example.com', 27, NOW(), NOW()),
('Федор Соколов', 'fedor.sokolov.35@example.com', 34, NOW(), NOW()),
('Юрий Медведев', 'yury.medvedev.36@example.com', 30, NOW(), NOW()),
('Ярослав Козлов', 'yaroslav.kozlov.37@example.com', 25, NOW(), NOW()),
('Александр Новиков', 'alexander.novikov.38@example.com', 28, NOW(), NOW()),
('Алексей Федоров', 'alexey.fedorov.39@example.com', 33, NOW(), NOW()),
('Андрей Морозов', 'andrey.morozov.40@example.com', 26, NOW(), NOW()),
('Антон Волков', 'anton.volkov.41@example.com', 32, NOW(), NOW()),
('Артем Алексеев', 'artem.alekseev.42@example.com', 29, NOW(), NOW()),
('Борис Лебедев', 'boris.lebedev.43@example.com', 35, NOW(), NOW()),
('Вадим Семенов', 'vadim.semenov.44@example.com', 27, NOW(), NOW()),
('Валентин Егоров', 'valentin.egorov.45@example.com', 31, NOW(), NOW()),
('Валерий Павлов', 'valery.pavlov.46@example.com', 24, NOW(), NOW()),
('Виктор Козлов', 'viktor.kozlov.47@example.com', 36, NOW(), NOW()),
('Владимир Степанов', 'vladimir.stepanov.48@example.com', 28, NOW(), NOW()),
('Вячеслав Николаев', 'vyacheslav.nikolaev.49@example.com', 30, NOW(), NOW()),
('Геннадий Орлов', 'gennady.orlov.50@example.com', 25, NOW(), NOW());

DO $$
DECLARE
    i INTEGER;
    first_names TEXT[] := ARRAY['Александр', 'Алексей', 'Андрей', 'Антон', 'Артем', 'Борис', 'Вадим', 'Валентин', 'Валерий', 'Виктор', 'Владимир', 'Вячеслав', 'Геннадий', 'Георгий', 'Григорий', 'Денис', 'Дмитрий', 'Евгений', 'Егор', 'Иван', 'Игорь', 'Илья', 'Кирилл', 'Константин', 'Максим', 'Михаил', 'Николай', 'Олег', 'Павел', 'Петр', 'Роман', 'Сергей', 'Станислав', 'Степан', 'Федор', 'Юрий', 'Ярослав'];
    last_names TEXT[] := ARRAY['Иванов', 'Петров', 'Сидоров', 'Смирнов', 'Кузнецов', 'Попов', 'Васильев', 'Соколов', 'Михайлов', 'Новиков', 'Федоров', 'Морозов', 'Волков', 'Алексеев', 'Лебедев', 'Семенов', 'Егоров', 'Павлов', 'Козлов', 'Степанов', 'Николаев', 'Орлов', 'Андреев', 'Макаров', 'Никитин', 'Захаров', 'Зайцев', 'Соловьев', 'Борисов', 'Яковлев', 'Григорьев', 'Романов', 'Воробьев', 'Сорокин', 'Медведев', 'Козлов', 'Новиков'];
    domains TEXT[] := ARRAY['gmail.com', 'yandex.ru', 'mail.ru', 'outlook.com', 'yahoo.com', 'rambler.ru', 'bk.ru', 'list.ru'];
    first_name TEXT;
    last_name TEXT;
    domain TEXT;
    email TEXT;
    age INTEGER;
BEGIN
    FOR i IN 51..1000 LOOP
        first_name := first_names[1 + (i % array_length(first_names, 1))];
        last_name := last_names[1 + (i % array_length(last_names, 1))];
        domain := domains[1 + (i % array_length(domains, 1))];
        email := LOWER(first_name) || '.' || LOWER(last_name) || '.' || i || '@' || domain;
        age := 18 + (i % 50);
        INSERT INTO users (name, email, age, created_at, updated_at)
        VALUES (first_name || ' ' || last_name, email, age, NOW(), NOW());
    END LOOP;
END $$;
