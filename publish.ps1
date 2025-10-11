# PowerShell скрипт для публикации новой версии модуля

Write-Host "🚀 Fingerprint Injector - Publishing Script" -ForegroundColor Cyan
Write-Host "==========================================" -ForegroundColor Cyan
Write-Host ""

# Проверка что мы в git репозитории
if (-not (Test-Path .git)) {
    Write-Host "❌ Ошибка: не найден .git директория" -ForegroundColor Red
    Write-Host "Выполните: git init"
    exit 1
}

# Запрос версии
$VERSION = Read-Host "Введите версию (например, v1.0.0)"

if ([string]::IsNullOrWhiteSpace($VERSION)) {
    Write-Host "❌ Версия не может быть пустой" -ForegroundColor Red
    exit 1
}

# Проверка формата версии
if ($VERSION -notmatch "^v\d+\.\d+\.\d+$") {
    Write-Host "⚠️  Внимание: версия должна быть в формате v1.0.0" -ForegroundColor Yellow
    $continue = Read-Host "Продолжить? (y/n)"
    if ($continue -ne "y") {
        exit 1
    }
}

Write-Host ""
Write-Host "📝 Проверка статуса..." -ForegroundColor Yellow

# Проверка что нет незакоммиченных изменений
$status = git status -s
if ($status) {
    Write-Host "⚠️  Есть незакоммиченные изменения:" -ForegroundColor Yellow
    git status -s
    Write-Host ""
    $doCommit = Read-Host "Сделать коммит? (y/n)"
    
    if ($doCommit -eq "y") {
        $commitMsg = Read-Host "Сообщение коммита"
        git add .
        git commit -m "$commitMsg"
    } else {
        Write-Host "❌ Отменено" -ForegroundColor Red
        exit 1
    }
}

Write-Host ""
Write-Host "🧪 Запуск тестов..." -ForegroundColor Yellow
go test ./...

if ($LASTEXITCODE -ne 0) {
    Write-Host "❌ Тесты не прошли!" -ForegroundColor Red
    $continue = Read-Host "Продолжить несмотря на ошибки? (y/n)"
    if ($continue -ne "y") {
        exit 1
    }
}

Write-Host ""
Write-Host "📦 Создание тега $VERSION..." -ForegroundColor Yellow
git tag $VERSION

if ($LASTEXITCODE -ne 0) {
    Write-Host "❌ Ошибка создания тега" -ForegroundColor Red
    exit 1
}

Write-Host ""
Write-Host "⬆️  Загрузка на GitHub..." -ForegroundColor Yellow

# Push код
git push
if ($LASTEXITCODE -ne 0) {
    Write-Host "❌ Ошибка push кода" -ForegroundColor Red
    git tag -d $VERSION
    exit 1
}

# Push тег
git push origin $VERSION
if ($LASTEXITCODE -ne 0) {
    Write-Host "❌ Ошибка push тега" -ForegroundColor Red
    git tag -d $VERSION
    exit 1
}

Write-Host ""
Write-Host "✅ Успешно опубликовано!" -ForegroundColor Green
Write-Host ""
Write-Host "📋 Следующие шаги:" -ForegroundColor Cyan
Write-Host "  1. Создайте Release на GitHub (опционально)"
Write-Host "  2. Обновите CHANGELOG.md"
Write-Host "  3. В других проектах используйте:"
Write-Host "     go get github.com/YOUR_USERNAME/fingerprint-injector@$VERSION" -ForegroundColor Yellow
Write-Host ""

