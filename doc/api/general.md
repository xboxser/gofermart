# Общая схема взаимодействия АPI


```mermaid

sequenceDiagram
  participant P3 as Пользователь
  participant P1 as Клиент
  participant P2 as Гофермарт
  participant P4 as Accrual
  P3 ->> P1: Регистрация пользователя
  P1 ->> P2: POST /api/user/register
  P2 ->> P1: HTTP-заголовок Authorization
  P3 ->> P1: Аутентификация пользователя
  P1 ->> P2: POST /api/user/login
  P2 ->> P1: HTTP-заголовок Authorization
  P3 ->> P1: Загрузка номера заказа
  P1 ->> P2: POST /api/user/orders
  P2 ->> P1: Признак загрузки заказа
  P3 ->> P1: Получение списка загруженных номеров заказов
  P1 ->> P2: GET /api/user/orders
  P2 ->> P1: Список заказов
  P3 ->> P1: Получение текущего баланса пользователя
  P1 ->> P2: GET /api/user/balance
  P2 ->> P1: Информация по счету баллов лояльности
  P3 ->> P1: Запрос на списание средств
  P1 ->> P2: POST /api/user/balance/withdraw
  P2 ->> P1: Информация по списанию средств
  P3 ->> P1: Получение информации о выводе средств
  P1 ->> P2: GET /api/user/withdrawals
  P2 ->> P1: Список операций
  P2 ->> P4: Получение информации по заказам
  P2 ->> P4: GET /api/orders/{number}
  P4 ->> P2: Информация о расчёте начислений баллов лояльности.

```
