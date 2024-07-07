package config

import "time"

type Config struct {
	Application Application
	Adapters    Adapters
	Services    Services
}

type Services struct {
	AuthJWT AuthJWT
}

type AuthJWT struct {
	JWTSecret string `config:"envVar"`
}

type Application struct {
	Name    string
	Version string
}

type Adapters struct {
	Primary   Primary
	Secondary Secondary
}

type Secondary struct {
	Databases Databases
}

type Primary struct {
	HttpAdapter HttpAdapter
}

type Database struct {
	Host     string `config:"envVar"`
	Port     string `config:"envVar"`
	Type     string
	Name     string
	User     string `config:"envVar"`
	Password string `config:"envVar"`
}

type HttpAdapter struct {
	Server Server
	Router Router
}

type Server struct {
	Port string
}

type Router struct {
	Shutdown            shutdown
	Timeout             timeout
	AuthorizationConfig string `config:"envVar"`
}

type shutdown struct {
	Duration time.Duration
}

type timeout struct {
	Duration time.Duration
}

type Databases struct {
	Todo Database
}
