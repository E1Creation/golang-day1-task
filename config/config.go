package config

import (
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Port   string `mapstructure:"PORT"`
	DBConn string `mapstructure:"DB_Conn"`
}

// func Load() *Config {
// 	// Viper
// 	viper.AutomaticEnv()
// 	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

// 	if _, err := os.Stat(".env"); err == nil {
// 		viper.SetConfigFile(".env")
// 		if err := viper.ReadInConfig(); err != nil {
// 			log.Fatal(err)
// 		}
// 	}

// 	config := Config{
// 		Port:   viper.GetString("PORT"),
// 		DBConn: viper.GetString("DB_CONN"),
// 	}
// 	// Setup database
// 	db, err := InitDB(config.DBConn)
// 	if err != nil {
// 		log.Fatal("Failed to initialize database:", err)
// 	}
// 	defer db.Close()

//		if config.Port == "" {
//			log.Fatal("PORT belum diset")
//		}
//		return &config
//	}
func Load() *Config {
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if _, err := os.Stat(".env"); err == nil {
		viper.SetConfigFile(".env")
		_ = viper.ReadInConfig()
	}

	return &Config{
		Port:   viper.GetString("PORT"),
		DBConn: viper.GetString("DB_CONN"),
	}
}
