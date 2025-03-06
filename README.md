# Биржа непрерывных во времени заявок
```
request_handler -> market_maker:

REST
JSON data format:
	Create:
		{
			"user_id": "int",
			"order_id": "int",
			"pair": ["string", "string"],
			"buy_sell_indicator": bool,
			"amount": float,
			"price_low": float,
			"price_high": float,
			"speed": float,
		}
	Cancel:
		{
			"user_id": "int"
			"order_id": "int"
		}


market_maker -> request_handler

Kafka connection
JSON data format:
	Every N seconds:
		{
			"order_id": "int",
			"action": "string ('Partial', 'Full')",
			"amount": "int",
			"average_price": "optional:int"
		}
	

price_calculator -> frontend

WebSocket connect:
JSON data format:
	Every N ms:
		{
			"pair": ["string", "string"],
			"price": float
		}
# Свеча - пока trading-view
```
