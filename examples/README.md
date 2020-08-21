# examplesの動かし方

## 検証用を叩く場合

`$ API_PASSWORD=password go run orders/orders.go`

`password` の部分に検証用で設定したパスワードを入れてください


## 本番用を叩く場合

`$ API_PASSWORD=password IS_PROD=true go run orders/orders.go`

`password` の部分に本番用で設定したパスワードを入れてください
