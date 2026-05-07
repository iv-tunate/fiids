package handlers

import "github.com/iv-tunate/fiids/config"

type ConfigHandler struct {
	Config *config.ApiConfig
}

func New(cfg *config.ApiConfig) *ConfigHandler{
	return &ConfigHandler{Config: cfg}
}