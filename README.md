# 🍃 Tannery App API

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)](https://golang.org)
[![Gin Framework](https://img.shields.io/badge/Gin-1.9+-00ADD8?style=flat&logo=gin)](https://gin-gonic.com)
[![MongoDB](https://img.shields.io/badge/MongoDB-6.0+-47A248?style=flat&logo=mongodb)](https://mongodb.com)
[![Swagger](https://img.shields.io/badge/Swagger-2.0-85EA2D?style=flat&logo=swagger)](https://swagger.io)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)

## 📋 О проекте

Tannery App - это высокопроизводительный мессенджер с микросервисной архитектурой, построенный на Go с использованием Gin фреймворка и MongoDB.

### Ключевые особенности

- 🚀 **Высокая производительность** - асинхронная обработка запросов
- 🔐 **Безопасность** - JWT аутентификация и авторизация
- 📝 **API документация** - Автоматическая Swagger документация
- 🗄️ **Масштабируемость** - Горизонтальное масштабирование MongoDB
- 🐳 **Docker готов** - Полная контейнеризация
- 🔄 **CI/CD** - Готовые пайплайны для деплоя

### Технологический стек

| Компонент | Технология | Версия |
|-----------|------------|---------|
| Язык | Go | 1.21+ |
| Web фреймворк | Gin | 1.9+ |
| База данных | MongoDB | 6.0+ |
| Кэширование | Redis | 7.0+ |
| Документация | Swaggo/Swagger | 2.0 |
| Конфигурация | Viper | 1.18+ |

## 🚀 Быстрый старт

### Требования

- Go 1.21 или выше
- MongoDB 6.0 или выше
- Redis 7.0 или выше (опционально)
- Make (опционально)

### Установка

1. **Клонирование репозитория**

```bash
git clone https://github.com/desmeach/tannery.git
cd tannery
