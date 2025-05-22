#pragma once

#include <usecases/engine_manager/engine_manager.hpp>

#include "usecases/order_book_configuration/order_book_configuration.hpp"
#include "contracts/order_book_configuration_usecase.hpp"

#include "pkg/transport/api/rest/server.hpp"

class MatchingEngineService {
 public:
  MatchingEngineService();

  void Run();

 private:
  std::shared_ptr<IOrderBookConfiguration> order_book_config_ = nullptr;
  std::shared_ptr<HTTPServer> http_server_{nullptr};
  std::shared_ptr<EngineManager> engine_manager_{nullptr};
};