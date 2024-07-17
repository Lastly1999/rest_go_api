package config

type BasicAppConfig struct {
	Mode    string `yaml:"mode"`
	Appname string `yaml:"appname"`
	Port    int    `yaml:"port"`
}

type MysqlConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"cypher"`
	Dbname   string `yaml:"dbname"`
}

type RedisConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Password string `yaml:"cypher"`
	DB       int    `yaml:"db"`
}

type JwtConfig struct {
	SecretKey string `yaml:"secretkey"`
	ExpiresIn int64  `yaml:"expiresin"`
	Isuser    string `yaml:"isuser"`
	HeaderKey string `yaml:"headerkey"`
}

type DatabaseConfig struct {
	Mysql MysqlConfig `yaml:"mysql"`
	Redis RedisConfig `yaml:"redis"`
}

type AppConfig struct {
	App      BasicAppConfig `yaml:"app"`
	Database DatabaseConfig `yaml:"database"`
	Jwt      JwtConfig      `yaml:"jwt"`
}
