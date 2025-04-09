package config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // Драйвер PostgreSQL
)

// Параметры подключения к БД (замените на свои)
const (
	host     = "localhost" // Адрес сервера БД
	port     = 5432        // Порт PostgreSQL
	user     = "имя_пользователя"     // Имя пользователя БД
	password = "пароль_от_БД"   // Пароль пользователя
	dbname   = "Название_БД" // Название базы данных
)

// InitDB инициализирует подключение к PostgreSQL
func InitDB() (*sql.DB, error) {
	// Формируем строку подключения
	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	// Открываем соединение с БД
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, fmt.Errorf("ошибка подключения к БД: %v", err)
	}

	// Проверяем соединение
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("ошибка проверки соединения: %v", err)
	}

	fmt.Println("Успешное подключение к PostgreSQL!")
	return db, nil
}
