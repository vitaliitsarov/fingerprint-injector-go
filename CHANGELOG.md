# Changelog

Все значимые изменения в проекте будут документированы в этом файле.

Формат основан на [Keep a Changelog](https://keepachangelog.com/en/1.0.0/).

## [1.0.0] - 2024-10-11

### Добавлено

- Начальный релиз библиотеки fingerprint-injector
- Поддержка основных параметров fingerprint:
  - User Agent
  - Platform
  - Vendor
  - Languages
  - Screen parameters
  - WebGL
  - Canvas fingerprinting protection
  - WebRTC control
  - Timezone
  - Battery API
  - Hardware concurrency
  - Device memory
- Готовые пресеты для разных платформ:
  - Chrome 119 Windows 11
  - Chrome 119 MacOS
  - Chrome 119 Linux
  - Chrome 119 Android
- Утилиты для генерации случайных fingerprints
- Скрытие признаков автоматизации (webdriver, chrome.runtime)
- Полная интеграция с chromedp
- Примеры использования:
  - Базовый пример
  - Кастомный fingerprint
  - Stealth режим
  - С прокси
  - Множественные сессии
  - Случайный fingerprint
- Юнит-тесты
- Документация (README, CONTRIBUTING)
- MIT License

### Особенности

- Простой и понятный API
- Легкая настройка под конкретные задачи
- Поддержка всех современных техник fingerprinting
- Максимальная защита от детекции
