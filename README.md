# widlberries

Тестовое задание WBTech

В сервисе:
<ol>
  <li>Подключение и подписка на канал в nats-streaming</li>
  <li>Полученные данные записываются в Postgres</li>
  <li>Полученные данные сохранить in memory в сервисе (Кеш)</li>
  <li>В случае падения сервиса восстанавливается Кеш из Postgres</li>
  <li>http сервер, выдающий данные по id из кеша</li>
</ol>

Сборка:
<ol>
  <li>make up - поднять контейнеры
  <li>make migrate - создать таблицу в Postgres
  <li>make server - поднять сервер
  <li>make publisher - собрать publisher
  <li>./publisher - запустить publisher
</ol>
publisher ожидает путь до json, который мы хотим добавить
