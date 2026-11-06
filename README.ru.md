# The Battle on the Black Sea — Морской бой

[English](README.md)

Перенос на Go морской стратегии Владимира Башурова (BADY), выпущенной для DOS в 1992 году. Расставьте десять кораблей, сразитесь с компьютером, пройдите кампанию на Чёрном море и улучшите рекорд своего капитана.

Сохранены оригинальная графика 320×200 в 16 цветах, растровые шрифты, переходы между экранами, боевая анимация и эффекты PC Speaker. Восстановленные процедуры Pascal выполняются как скомпилированные функции Go: расстановка флотов, стратегия компьютера, подсчёт результатов и 16-битная арифметика. Ресурсы встроены в исполняемый файл; DOS и исходный EXE для запуска не требуются. Изображение выводится с пропорциями старого экрана 4:3.

## Управление

| Клавиша | Действие |
| --- | --- |
| Стрелки | Перемещение корабля или прицела |
| Enter | Установка корабля, выстрел, переход с информационного экрана |
| Insert / Page Down | Поворот корабля при расстановке |
| F2 | Дополнительный поворот для клавиатур без Insert |
| F1 | Включение и выключение звука |
| Esc | Выход во время расстановки или боя |
| F11 / Alt+Enter | Полноэкранный режим |
| Backspace | Исправление имени капитана |

Сохранён оригинальный ввод мышью. Имя капитана вводится латиницей. Карточки результатов сохраняются при выходе из игры и закрытии окна в пользовательском каталоге настроек системы, в папке `BattleOnTheBlackSea`.

## Сборка

Установите **Go 1.25 или новее**. **Ebitengine 2.10.5** и её зависимости зафиксированы в `go.mod` и `go.sum`. Команды выполняются из корня проекта. При первой сборке Go загрузит зависимости.

Все шесть настольных целей поддерживают **`CGO_ENABLED=0`**. Для кросс-компиляции достаточно Go; компилятор C, заголовки целевой системы и SDK Apple не нужны.

### Windows — amd64 и arm64

PowerShell:

```powershell
$env:CGO_ENABLED="0"
$env:GOOS="windows"
$env:GOARCH="amd64"
go build -buildvcs=false -trimpath -ldflags="-s -w -H=windowsgui" -o BattleOnTheBlackSea-windows-amd64.exe .
$env:GOARCH="arm64"
go build -buildvcs=false -trimpath -ldflags="-s -w -H=windowsgui" -o BattleOnTheBlackSea-windows-arm64.exe .
```

PNG-иконка встроена и используется окном игры; `icon.ico` приложен для оформления ресурсов исполняемого файла.

### Linux — amd64 и arm64

Из Linux или macOS:

```sh
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -buildvcs=false -trimpath -ldflags="-s -w" -o BattleOnTheBlackSea-linux-amd64 .
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -buildvcs=false -trimpath -ldflags="-s -w" -o BattleOnTheBlackSea-linux-arm64 .
```

Из **Windows CMD**, для Linux ARM64:

```cmd
set "CGO_ENABLED=0"
set "GOOS=linux"
set "GOARCH=arm64"
go build -buildvcs=false -trimpath -ldflags="-s -w" -o BattleOnTheBlackSea-linux-arm64 .
set "CGO_ENABLED="
set "GOOS="
set "GOARCH="
```

Полученный файл запускается на Linux ARM64. После переноса выполните `chmod +x BattleOnTheBlackSea-linux-arm64`, затем `./BattleOnTheBlackSea-linux-arm64`.

Для запуска в Linux нужны X11 либо XWayland и драйвер OpenGL. Звук использует PulseAudio/PipeWire либо резервный вывод ALSA. В Debian/Ubuntu основные графические библиотеки устанавливаются так:

```sh
sudo apt install libx11-6 libgl1 libglx-mesa0 libxcursor1 libxi6 libxinerama1 libxrandr2 libxrender1 libxext6
```

Для резервного вывода ALSA установите `libasound2` на Debian 12/Ubuntu 22.04 или `libasound2t64` на Debian 13/Ubuntu 24.04 и новее. Это библиотеки для запуска, не зависимости сборки.

### macOS — Intel и Apple Silicon

Из Unix-оболочки на Linux или macOS:

```sh
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -buildvcs=false -trimpath -ldflags="-s -w" -o BattleOnTheBlackSea-macos-amd64 .
CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -buildvcs=false -trimpath -ldflags="-s -w" -o BattleOnTheBlackSea-macos-arm64 .
```

В Windows CMD используйте такой же синтаксис `set`, как в примере Linux, задав `GOOS=darwin` и нужную `GOARCH`. `icon.icns` приложен для оформления необязательного пакета `.app`.

Для запуска из исходников используйте `go run .` в графическом сеансе. Собранный исполняемый файл содержит всю игру; отдельные иконки и скриншоты для запуска не нужны.

Проверка: `go test ./...` и `go vet ./...`. Сохранены регрессионные тесты исправления зависания при запуске звука в Windows.

## Авторы и границы проверки

Оригинальная игра, графика и персонажи: **Владимир Башуров (BADY), июнь 1992 года**. Права на оригинальные материалы принадлежат их авторам; в предоставленном DOS-архиве отсутствует лицензия на распространение. Ebitengine и его зависимости сохраняют собственные лицензии.

Логика и отрисовка восстановлены из предоставленных исполняемого файла и ресурсов. Сохранены структура данных и округление вещественной арифметики. Для анимации задана фиксированная временная база: скорость DOS-игры зависела от компьютера. Проведено попиксельное сравнение заставки и статистики, проверены арифметика и расстановка флота по оригинальным процедурам, пройдены полные бои до победы и поражения. Проверена сборка Windows и Linux для amd64/arm64; сборку macOS необходимо проверить на Mac. Идентичность темпа и акустического звучания на всех старых и современных компьютерах не подтверждена.
