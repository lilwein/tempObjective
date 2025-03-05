package postgres

type Config struct {
	Host     string `yaml:"host" mapstructure:"host" json:"host"`
	Port     int    `yaml:"port" mapstructure:"port" json:"port"`
	User     string `yaml:"user" mapstructure:"user" json:"user"`
	Password string `yaml:"password" mapstructure:"password" json:"password"`
	Dbname   string `yaml:"dbname" mapstructure:"dbname" json:"dbname"`
}
