# Демо-приложение на Go (Golang) для Маркетплейса МоегоСклада

Демо-приложение предназначено для демонстрации взаимодействия приложений с Маркетплейсом МоегоСклада. 
На основе этого проекта можно создавать любое количество различных приложений, которые будут работать на одном сервере. 
Для этого необходимо создать папку `internal/appUid` с названием соответствующим appUid вашего приложения и описать в ней логику работы (главное чтобы приложение удовлетворяло интерфейсу `App`).
Затем прописать вызов функций приложения в соответствующих обработчиках (файл `internal/server/handlers.go`) и в описании сервера (`internal/server/server.go`).

В демо-приложении реализованы следующие функции:
* Активация (с получением токена доступа к JSON API 1.2) и деактивация по Vendor API
* Использование iframe для настройки приложения администратором аккаунта с обновлением статуса в Маркетплейсе
* Получение контекста пользователя для iframe (отображение информации по пользователю, проверка прав администратора)
* Получение информации из МоегоСклада по JSON API 1.2 с доступом по токену

ВНИМАНИЕ! Демо-приложение предназначено для ознакомления с функционалом Маркетплейса МоегоСклада и вопросы 
безопасности и стабильности его работы не рассматривались в рамках его разработки. 

## Настройка (конфигурирование) приложения

Перед использованием приложения нужно настроить следующие конфигурационные параметры:

* `appUid`           - appUid приложения в Маркетплейсе
* `secretKey`        - секретный ключ для подписи JWT

Для настройки конфигурационных параметров нужно переименовать папку internal/dummy-sloudel.sorochinsky соответственно желаемому appUid,
а также скопировать значение секретного ключа в файл internal/appUid/secret.key.

## Структура файлов приложения

### Основные файлы приложения

* `main.go`                                        - main package приложения
* `internal/app/app.go`                            - описание интерфейса приложения, а также базового приложения и его методов
* `internal/dummy-sloudel.sorochinsky/dummy.go`    - описание демо-приложения
* `internal/dummy-sloudel.sorochinsky/secret.key`  - секретный ключ приложения
* `internal/server/server.go`                      - описание сервера и логики его работы
* `internal/server/handlers.go`                    - обработчики REST-эндпоинтов и HTML
* `jsonapi/jsonapi.go`                             - логика работы с JSON API на стороне МоегоСклада
* `vendorapi/vendorapi.go`                         - логика работы с Vendor API на стороне МоегоСклада

### Логи

Логи пишутся в стандартный вывод, но по желанию можно писать их в отдельный файл. Поясняющие комментарии есть в файле `main.go`

### Данные

Состояние приложений хранится в базе данных (PostregSQL 12) в таблицах (см. `internal/migrations.sql`) baseapp и dummyapp.
