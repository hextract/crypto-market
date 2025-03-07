#!/bin/bash

swagger generate server -f stack_connector/api/swagger/stack_connector.yaml -t stack_connector/internal # --exclude-main --principal models.User

swagger generate server -f auth/api/swagger/auth.yaml -t auth/internal # --exclude-main

echo "REGENERATED. NOW TIDYING"

cd stack_connector || exit
go mod tidy

cd ../auth || exit
go mod tidy

