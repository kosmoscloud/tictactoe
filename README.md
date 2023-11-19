# tictactoe
## projekt pilotażowy
założenia projektu (na dzień 12/11/23):
- w pełni funkcjonalny backend aplikacji pozwalający tworzyć użytkowników, tworzyć pokoje do gry w kółko i krzyżyk oraz dołączać do nich i wykonywać ruchy, które następnie zapisywane są w bazie danych mySQL
- cały system wywoływany jednym poleceniem docker compose
- zewnętrzny serwis autoryzacyjny
- (opcjonalne) wykorzystanie protocol buffers do komunikacji wewnętrznej
## Docker
Baza danych włączana jest za pomocą kontenera na porcie 3306. Komenda do zbudowania obrazu: \
```docker build -t tictactoe:database ./src/database/``` \
Komenda do uruchamiania aplikacji w kontenarze \
```docker run -p 3306:3306 -d tictactoe:database```
## Postman 
Aktualna kolekcja znajduje się pod [tym linkiem](https://solar-capsule-457081.postman.co/workspace/3d5353de-a10b-461d-95d5-9f19dc4ff8f9).