### Создание нового проекта

## Инициализация модуля

go mod init github.com/DimonVelikij/storage

## Теггирование модуля

git tag v0.0.1

git push --tags

go get github.com/DimonVelikij/storage@v0.0.1

## Публикация новой версии модуля

В файле go.mod меняем название модуля, например, module github.com/DimonVelikij/storage/v2

Заменяем старое название модуля на новое везде где он импортируется

Публикуем новый тег