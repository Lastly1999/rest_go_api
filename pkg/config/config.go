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
	Password string `yaml:"password"`
	Dbname   string `yaml:"dbname"`
}

type DatabaseConfig struct {
	MysqlConfig MysqlConfig `yaml:"mysql"`
}

type AppConfig struct {
	App      BasicAppConfig `yaml:"app"`
	Database DatabaseConfig `yaml:"database"`
}
