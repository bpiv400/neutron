export DEMOWALLET_1=$(gaia-wasm-zoned keys show demowallet1 -a --keyring-backend test --home ./data/test-1) && echo $DEMOWALLET_1;
export DEMOWALLET_2=$(gaia-wasm-zoned keys show demowallet2 -a --keyring-backend test --home ./data/test-2) && echo $DEMOWALLET_2;

gaia-wasm-zoned tx interchaintxs register-interchain-account connection-0 cosmos1mjk79fjjgpplak5wq838w0yd982gzkyfrk07am \
    --from $DEMOWALLET_1 --chain-id test-1 --home ./data/test-1 --node tcp://localhost:16657 --keyring-backend test -y