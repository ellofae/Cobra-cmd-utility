# Cobra-cmd-utility
___________________________________

## Утилита командной строки, позволяющая работать с базой данных через командную строку

Утилита состоит из трех частей: **sysInfo**, **dbSetting**, **dbUpdate**

### sysInfo
**sysInfo** позволяет вывести информацию об операционной системе пользователя при помощи комманды: --spec **[OPTION]**
* --spec all (default) - показывает информацию о хосте, ОС и памяти
* --spec host (или --spec) - показывает информацию о хосте и ОС
* --spec mem (или --spec memory) - показывает информацию о памяти.

## Пример использования sysinfo:
![result1](https://github.com/ellofae/Cobra-cmd-utility/blob/main/imgs/Screenshot%20from%202023-03-15%2022-05-44.png?raw=true)

### dbUpdate
**dbUpdate** позволяет считать данные из файла **{database.txt}** каталога **(./file-storage/)** и записать данные в таблицу **data** базы данных **{database.db}**.
Функционал осуществляется при помощи команды:  --df **[OPTION]**
* --df **{file.txt}** - считывает данные из файла **{file.txt}** каталога  **(./file-storage/})** и записываеь данные в таблицу **data** базы данных **{file.db}**

## Пример использования dbUpdate:
![result1](https://github.com/ellofae/Cobra-cmd-utility/blob/main/imgs/Screenshot%20from%202023-03-15%2021-47-50.png?raw=true)

### dbSetting
**dbSetting** позволяет вывести информацию о доступных пользователю базах данных каталога **(./db-storage/...)** командой: --ops **[OPTION]**
* --ops list-db (default) - показывает все доступные пользователю базы данных из каталога **(./db-storage/...)**
* --ops show --dbName **{name.db}** - выводит данные из таблицы **data** базы данных **{name.db}**

## Пример использования dbSetting:
![result1](https://github.com/ellofae/Cobra-cmd-utility/blob/main/imgs/Screenshot%20from%202023-03-15%2021-54-39.png?raw=true)

![result1](https://github.com/ellofae/Cobra-cmd-utility/blob/main/imgs/Screenshot%20from%202023-03-15%2021-43-34.png?raw=true)
