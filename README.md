# ads-microservice

Решение тестового задания, текст ниже. ✅ - выполнено

👣 Тестовое задание на позицию Golang разработчика.

Задача
Необходимо создать сервис для хранения и подачи объявлений. Объявления должны храниться в базе данных. Сервис должен предоставлять API, работающее поверх HTTP в формате JSON.


Требования
✅  - Простая инструкция для запуска (в идеале — с возможностью запустить через docker-compose up, но это необязательно);
  Проект запускается через клонирование репозитория и make up
✅  - Валидация полей: не больше 3 ссылок на фото, описание не больше 1000 символов, название не больше 200 символов;

Детали
Метод получения списка объявлений

✅ - Пагинация: на одной странице должно присутствовать 10 объявлений;
✅ - Cортировки: по цене (возрастание/убывание) и по дате создания (возрастание/убывание);
✅ - Поля в ответе: название объявления, ссылка на главное фото (первое в списке), цена.

Метод получения конкретного объявления
- Обязательные поля в ответе: название объявления, цена, ссылка на главное фото;
- Опциональные поля (можно запросить, передав параметр fields): описание, ссылки на все фото.

Метод создания объявления:
- Принимает все вышеперечисленные поля: название, описание, несколько ссылок на фотографии (сами фото загружать никуда не требуется), цена;
- Возвращает ID созданного объявления и код результата (ошибка или успех).

Дополнительно
- Юнит тесты: постарайтесь достичь покрытия в 70% и больше;
✅ - Контейнеризация: есть возможность поднять проект с помощью команды docker-compose up;
- Архитектура сервиса описана в виде текста и/или диаграмм

✅ - Документация: есть структурированное описание методов сервиса.
