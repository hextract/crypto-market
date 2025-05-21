#pragma once

#include "env_loader.hpp"
#include <cstdlib>
#include <string>
#include <iostream>

class Config {
  public:
  static void Init(const std::string &env_file = ".env") {
    LoadEnvFile(env_file);
  }

  static size_t GetDeltaTime() {
    const char *v = std::getenv("DELTA_TIME");
    return v ? std::stoul(v) : 7;
  }

  static size_t GetServerPort() {
    const char *v = std::getenv("SERVER_PORT");
    return v ? std::stoul(v) : 18888;
  }

  static std::string GetExchangeConnectorUrl() {
    const char *v = std::getenv("EXCHANGE_CONNECTOR_URL");
    return v ? v : "http://127.0.0.1:8000";
  }

  static std::string GetBackendUrl() {
    const char *v = std::getenv("BACKEND_URL");
    return v ? v : "http://127.0.0.1:8088";
  }
};
