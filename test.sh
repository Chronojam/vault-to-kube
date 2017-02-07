curl \
    -H "X-Vault-Token: ab7c85c7-40bc-100b-d3ad-90575c6b192f" \
    -H "Content-Type: application/json" \
    -X POST \
    -d '{"value":"Secret1"}' \
    https://localhost:8200/v1/secret/calum/secret1

curl \
    -H "X-Vault-Token: ab7c85c7-40bc-100b-d3ad-90575c6b192f" \
    -H "Content-Type: application/json" \
    -X POST \
    -d '{"value":"Secret2"}' \
    https://localhost:8200/v1/secret/calum/secret2

curl \
    -H "X-Vault-Token: ab7c85c7-40bc-100b-d3ad-90575c6b192f" \
    -X GET \
    https://localhost:8200/v1/secret/calum?list=true
