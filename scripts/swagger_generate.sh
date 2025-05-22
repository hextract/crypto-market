#!/bin/bash

cd ../

swagger generate server -f receiver/internal/presentation/web/api/swagger/receiver.yaml -t \
  receiver/internal/presentation/web --model-package dto --exclude-main

swagger generate client -f reporter/internal/infrastructure/adapters/delivery/web/api/swagger/delivery.yaml  \
  -t reporter/internal/infrastructure/adapters/delivery/web --model-package ../../../dto

swagger generate server -f fetcher/internal/presentation/web/api/swagger/fetcher.yaml -t \
  fetcher/internal/presentation/web --model-package dto --exclude-main

swagger generate client -f generator/internal/infrastructure/adapters/delivery/web/api/swagger/delivery.yaml  \
  -t generator/internal/infrastructure/adapters/delivery/web --model-package ../../../dto

echo "REGENERATED"

