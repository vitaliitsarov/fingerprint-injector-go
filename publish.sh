#!/bin/bash

# Скрипт для публикации новой версии модуля

echo "🚀 Fingerprint Injector - Publishing Script"
echo "=========================================="
echo ""

# Проверка что мы в git репозитории
if [ ! -d .git ]; then
    echo "❌ Ошибка: не найден .git директория"
    echo "Выполните: git init"
    exit 1
fi

# Запрос версии
read -p "Введите версию (например, v1.0.0): " VERSION

if [ -z "$VERSION" ]; then
    echo "❌ Версия не может быть пустой"
    exit 1
fi

# Проверка формата версии
if [[ ! $VERSION =~ ^v[0-9]+\.[0-9]+\.[0-9]+$ ]]; then
    echo "⚠️  Внимание: версия должна быть в формате v1.0.0"
    read -p "Продолжить? (y/n): " CONTINUE
    if [ "$CONTINUE" != "y" ]; then
        exit 1
    fi
fi

echo ""
echo "📝 Проверка статуса..."

# Проверка что нет незакоммиченных изменений
if [[ -n $(git status -s) ]]; then
    echo "⚠️  Есть незакоммиченные изменения:"
    git status -s
    echo ""
    read -p "Сделать коммит? (y/n): " DO_COMMIT
    
    if [ "$DO_COMMIT" = "y" ]; then
        read -p "Сообщение коммита: " COMMIT_MSG
        git add .
        git commit -m "$COMMIT_MSG"
    else
        echo "❌ Отменено"
        exit 1
    fi
fi

echo ""
echo "🧪 Запуск тестов..."
go test ./...

if [ $? -ne 0 ]; then
    echo "❌ Тесты не прошли!"
    read -p "Продолжить несмотря на ошибки? (y/n): " CONTINUE
    if [ "$CONTINUE" != "y" ]; then
        exit 1
    fi
fi

echo ""
echo "📦 Создание тега $VERSION..."
git tag $VERSION

if [ $? -ne 0 ]; then
    echo "❌ Ошибка создания тега"
    exit 1
fi

echo ""
echo "⬆️  Загрузка на GitHub..."

# Push код
git push
if [ $? -ne 0 ]; then
    echo "❌ Ошибка push кода"
    git tag -d $VERSION
    exit 1
fi

# Push тег
git push origin $VERSION
if [ $? -ne 0 ]; then
    echo "❌ Ошибка push тега"
    git tag -d $VERSION
    exit 1
fi

echo ""
echo "✅ Успешно опубликовано!"
echo ""
echo "📋 Следующие шаги:"
echo "  1. Создайте Release на GitHub (опционально)"
echo "  2. Обновите CHANGELOG.md"
echo "  3. В других проектах используйте:"
echo "     go get github.com/YOUR_USERNAME/fingerprint-injector@$VERSION"
echo ""

