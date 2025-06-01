package storage


import (  
    "database/sql"  
    "fmt"  
	_"log/slog"  
    "time"  
)  

type Config struct {  
    Host     string `yaml:"host" env:"DB_HOST"`  
    Port     int    `yaml:"port" env:"DB_PORT"`  
    User     string `yaml:"user" env:"DB_USER"`  
    Password string `yaml:"password" env:"DB_PASSWORD"`  
    DBName   string `yaml:"dbname" env:"DB_NAME"`  
    SSLMode  string `yaml:"sslmode" env:"DB_SSLMODE"`  
    
    // Настройки пула соединений  
    MaxOpenConns    int           `yaml:"max_open_conns"`  
    MaxIdleConns    int           `yaml:"max_idle_conns"`  
    ConnMaxLifetime time.Duration `yaml:"conn_max_lifetime"`  
    ConnMaxIdleTime time.Duration `yaml:"conn_max_idle_time"`  
}  

func (c *Config) DSN() string {  
    return fmt.Sprintf(  
        "host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",  
        c.Host, c.Port, c.User, c.Password, c.DBName, c.SSLMode,  
    )  
}  

func NewConnection(cfg *Config) (*sql.DB, error) {  
    db, err := sql.Open("postgres", cfg.DSN())  
    if err != nil {  
        return nil, fmt.Errorf("failed to open database: %w", err)  
    }  
    
    // Настройка пула соединений  
    db.SetMaxOpenConns(cfg.MaxOpenConns)  
    db.SetMaxIdleConns(cfg.MaxIdleConns)  
    db.SetConnMaxLifetime(cfg.ConnMaxLifetime)  
    db.SetConnMaxIdleTime(cfg.ConnMaxIdleTime)  
    
    // Проверка подключения  
    if err := db.Ping(); err != nil {  
        return nil, fmt.Errorf("failed to ping database: %w", err)  
    }  
    
    return db, nil  
}  