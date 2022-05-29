package main

import (
	"github.com/spf13/viper"
)

func GetConfig(path string) *viper.Viper {

	// Instancia a struct Vyper
	config := viper.New()

	// Diz qual será o caminho do arquivo de configuração
	config.AddConfigPath(path)

	// Diz qual será a extensão do arquivo de configuração
	config.SetConfigType("yml")

	// Diz qual será o prefixo para não existir conflitos entre variáveis de ambiente
	config.SetEnvPrefix("PREFIX")

	// Lê as configurações do arquivo
	config.ReadInConfig()

	// Olha primeiro para as variáveis de ambiente antes do arquivo de configuração
	config.AutomaticEnv()

	// Checa qualquer alteração no arquivo de configuração para refletir automaticamente
	config.WatchConfig()

	return config
}

func main() {
	// Só usar :D
	config := GetConfig(".")
	print(config.GetString("domain"))
}
