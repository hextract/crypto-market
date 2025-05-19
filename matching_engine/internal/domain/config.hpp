#pragma once

#include <string>

class Config {
 public:
  static size_t GetDeltaTime() { return 7; }

  static size_t GetServerPort() { return 18888; }

  static std::string GetExchangeConnectorUrl() { return "http://127.0.0.1:8000"; }
};