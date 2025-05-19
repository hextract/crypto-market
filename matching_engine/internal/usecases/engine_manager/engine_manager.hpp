#pragma once

#include <memory>

#include <domain/config.hpp>

#include <usecases/order_queue/order_queue.hpp>
#include <usecases/user_order_book/user_order_book.hpp>
#include <usecases/order_book/order_book.hpp>

#include <contracts/order_book_usecase.hpp>
#include <contracts/user_order_book_usecase.hpp>
#include <contracts/price_calc.hpp>
#include <contracts/order_queue_usecase.hpp>


class EngineManager {
 public:
  EngineManager(const std::shared_ptr<IOrderQueue> &order_queue,
                const std::shared_ptr<IUserOrderBook> &user_order_book,
                const std::shared_ptr<IOrderBook> &buy_order_book,
                const std::shared_ptr<IOrderBook> &sell_order_book,
                const std::shared_ptr<IPriceCalculator>& price_calculator)
      : order_queue_(order_queue), user_order_book_(user_order_book),
        buy_order_book_(buy_order_book), sell_order_book_(sell_order_book),
        price_calculator_(price_calculator) {time_delta_ = Config::GetDeltaTime(); }

  void Step();

  void Run();

  void SetBackendClient(const std::shared_ptr<IBackendClient>& client) {
    user_order_book_->SetBackendClient(client);
  }

 private:
  size_t time_delta_{0};
  std::shared_ptr<IOrderQueue> order_queue_{nullptr};
  std::shared_ptr<IUserOrderBook> user_order_book_{nullptr};
  std::shared_ptr<IOrderBook> buy_order_book_{nullptr};
  std::shared_ptr<IOrderBook> sell_order_book_{nullptr};
  std::shared_ptr<IPriceCalculator> price_calculator_{nullptr};
};
