#!/bin/bash

swagger generate server -f main/api/swagger/main.yaml -t main/internal --exclude-main # --principal models.User

#swagger generate server -f payment/api/swagger/payment.yaml -t payment/internal --exclude-main
#
swagger generate server -f auth/api/swagger/auth.yaml -t auth/internal --exclude-main

echo "REGENERATED. NOW TIDYING"

cd main || exit
go mod tidy

cd ../auth || exit
go mod tidy

