Запуск
"docker run -e "MEMORY=postgres" --rm -p 5000:5000 image_name" при запуске контейнера в env указывается место хранения данных, при остутсвии данной переменной автоматически выставляется "in-memory"

Данные о подключении к серверу и базе данных лежат в arg.env

Запросы 
POST http://127.0.0.1:5000/?url=https://stackoverflow.com/questions/49545146/how-to-exec-sql-file-with-commands-in-golang

GET http://127.0.0.1:5000/?url=w8FO6AlVj8

по ключу url передается полный url