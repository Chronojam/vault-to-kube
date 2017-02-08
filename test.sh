curl \
    -k \
    -H "X-Vault-Token: c33a75d9-810b-bd6b-53a0-d9b8355f504c" \
    -H "Content-Type: application/json" \
    -X POST \
    -d '{"value":"Secret1"}' \
    https://vault.chronojam.co.uk:8200/v1/secret/calum/secret1

curl -k \
    -H "X-Vault-Token: c33a75d9-810b-bd6b-53a0-d9b8355f504c" \
    -H "Content-Type: application/json" \
    -X POST \
    -d '{"value":"Secret2"}' \
    https://vault.chronojam.co.uk:8200/v1/secret/calum/secret2

curl -k \
    -H "X-Vault-Token: c33a75d9-810b-bd6b-53a0-d9b8355f504c" \
    -X GET \
    https://vault.chronojam.co.uk:8200/v1/secret/calum?list=true
