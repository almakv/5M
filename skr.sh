#!/bin/bash

# Создать директорию для кэша VCS
mkdir -p /home/almakvt/go/pkg/mod/cache/vcs

# Локировать директорию с конкретным хэшом
LOCK_DIR="/home/almakvt/go/pkg/mod/cache/vcs/0901dc1ef67fcce1c9b3ae51078740de4a0e2dc673e720584ac302973af82f36"
LOCK_FILE="$LOCK_DIR.lock"

# Проверить, существует ли файл блокировки, если нет, создать его
if [ ! -f "$LOCK_FILE" ]; then
    touch "$LOCK_FILE"
    echo "Lock file created: $LOCK_FILE"
else
    echo "Lock file already exists: $LOCK_FILE"
fi

# Перейти в директорию и показать теги
cd "$LOCK_DIR" || { echo "Directory not found: $LOCK_DIR"; exit 1; }
echo "Listing tags in $LOCK_DIR:"
git tag -l

# Показать доступные ревизии
echo "Listing remote references:"
git ls-remote -q origin

# Проверка доступности версии
echo "Checking for version v3.1.0 in gopkg.in/yaml.v3..."
GO_VERSION_CHECK=$(go list -m gopkg.in/yaml.v3@v3.1.0 2>&1)
echo "$GO_VERSION_CHECK"