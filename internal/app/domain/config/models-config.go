package config

import "time"

type Config struct {
	Application Application `yaml:"application"`
	Adapters    Adapters    `yaml:"adapters"`
	Services    Services    `yaml:"services"`
}

type Services struct {
	AuthJWT AuthJWT `yaml:"authJWT"`
}

type AuthJWT struct {
	JWTSecret string `yaml:"jwtSecret"`
}

type Application struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
}

type Adapters struct {
	Primary   Primary   `yaml:"primary"`
	Secondary Secondary `yaml:"secondary"`
}

type Secondary struct {
	Databases Databases `yaml:"databases"`
}

type Primary struct {
	HttpAdapter HttpAdapter `yaml:"httpAdapter"`
}

type Database struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Type     string `yaml:"type"`
	Name     string `yaml:"name"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

type HttpAdapter struct {
	Server Server
	Router Router
}

type Server struct {
	Port string `yaml:"port"`
}

type Router struct {
	Shutdown            shutdown `yaml:"shutdown"`
	Timeout             timeout  `yaml:"timeout"`
	AuthorizationConfig string   `yaml:"authorizationConfig"`
}

type shutdown struct {
	Duration time.Duration `yaml:"duration"`
}

type timeout struct {
	Duration time.Duration `yaml:"duration"`
}

type Databases struct {
	Todo Database `yaml:"todo"`
}
