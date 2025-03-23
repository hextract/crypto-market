package config

import (
	"errors"
	"os"
)

const (
	KeycloakPort          = "KEYCLOAK_INNER_PORT"
	KeycloakClient        = "KEYCLOAK_CLIENT"
	KeycloakRealm         = "KEYCLOAK_REALM"
	KeycloakClientSecret  = "KEYCLOAK_CLIENT_SECRET"
	KeycloakAdmin         = "KEYCLOAK_ADMIN"
	KeycloakAdminPassword = "KEYCLOAK_ADMIN_PASSWORD"
	KeycloakMasterRealm   = "master"
)

type Config struct {
	Client        string
	Realm         string
	ClientSecret  string
	Port          string
	Admin         string
	AdminPassword string
	MasterRealm   string
}

func NewConfig() (*Config, error) {
	keycloakClient := os.Getenv(KeycloakClient)
	if keycloakClient == "" {
		return nil, errors.New("keycloak client not set")
	}
	keycloakRealm := os.Getenv(KeycloakRealm)
	if keycloakRealm == "" {
		return nil, errors.New("keycloak realm not set")
	}
	keycloakSecret := os.Getenv(KeycloakClientSecret)
	if keycloakSecret == "" {
		return nil, errors.New("keycloak secret not set")
	}
	keycloakPort := os.Getenv(KeycloakPort)
	if keycloakPort == "" {
		return nil, errors.New("keycloak port not set")
	}
	keycloakAdmin := os.Getenv(KeycloakAdmin)
	if keycloakAdmin == "" {
		return nil, errors.New("keycloak admin not set")
	}
	keycloakAdminPassword := os.Getenv(KeycloakAdminPassword)
	if keycloakAdminPassword == "" {
		return nil, errors.New("keycloak admin password not set")
	}
	return &Config{
		Client:        keycloakClient,
		Realm:         keycloakRealm,
		ClientSecret:  keycloakSecret,
		Port:          keycloakPort,
		Admin:         keycloakAdmin,
		AdminPassword: keycloakAdminPassword,
		MasterRealm:   KeycloakMasterRealm,
	}, nil
}
