package cfg

import (
	"os"
	"runtime"
	"testing"
)

func TestConfigurate(t *testing.T) {
	// Определяем тестовые переменные окружения
	os.Setenv("RUN_ADDRESS", "localhost:8080")
	os.Setenv("SMTP_PORT", "587")
	os.Setenv("SMTP_SERVER", "smtp.example.com")
	os.Setenv("SMTP_LOGIN", "user@example.com")
	os.Setenv("SMTP_PASS", "password")
	os.Setenv("POSTGRES_CONNECTION", "host=localhost dbname=test sslmode=disable")
	os.Setenv("JWK_KEY", "supersecretkey")

	// Вызываем функцию конфигурации
	appCfg, messageCfg, postgresCfg, securityCfg := Configurate()

	// Проверяем наличие значений в конфигурации
	if appCfg.RunAddress != "localhost:8080" {
		t.Errorf("Expected RunAddress 'localhost:8080', got '%s'", appCfg.RunAddress)
	}

	if messageCfg.SMTPPort != 587 {
		t.Errorf("Expected SMTPPort 587, got %d", messageCfg.SMTPPort)
	}

	if messageCfg.SMTPServer != "smtp.example.com" {
		t.Errorf("Expected SMTPServer 'smtp.example.com', got '%s'", messageCfg.SMTPServer)
	}

	if messageCfg.SMTPLogin != "user@example.com" {
		t.Errorf("Expected SMTPLogin 'user@example.com', got '%s'", messageCfg.SMTPLogin)
	}

	if postgresCfg.ConnectionString != "host=localhost dbname=test sslmode=disable" {
		t.Errorf("Expected ConnectionString 'host=localhost dbname=test sslmode=disable', got '%s'", postgresCfg.ConnectionString)
	}

	if securityCfg.JWTKey != "supersecretkey" {
		t.Errorf("Expected JWTKey 'supersecretkey', got '%s'", securityCfg.JWTKey)
	}

	// Проверяем значения по умолчанию
	if messageCfg.FillWorkerCount != runtime.NumCPU()/2 {
		t.Errorf("Expected FillWorkerCount %d, got %d", runtime.NumCPU()/2, messageCfg.FillWorkerCount)
	}

	if messageCfg.SendWorkerCount != runtime.NumCPU()/2 {
		t.Errorf("Expected SendWorkerCount %d, got %d", runtime.NumCPU()/2, messageCfg.SendWorkerCount)
	}

	// Очистка переменных окружения
	os.Unsetenv("RUN_ADDRESS")
	os.Unsetenv("SMTP_PORT")
	os.Unsetenv("SMTP_SERVER")
	os.Unsetenv("SMTP_LOGIN")
	os.Unsetenv("SMTP_PASS")
	os.Unsetenv("POSTGRES_CONNECTION")
	os.Unsetenv("JWK_KEY")
}
