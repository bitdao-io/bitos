#!/bin/bash

KEY="mykey"
CHAINID="bitos_9000-1"
MONIKER="mymoniker"
DATA_DIR=$(mktemp -d -t bitos-datadir.XXXXX)

echo "create and add new keys"
./bitosd keys add $KEY --home $DATA_DIR --no-backup --chain-id $CHAINID --algo "eth_secp256k1" --keyring-backend test
echo "init Bitos with moniker=$MONIKER and chain-id=$CHAINID"
./bitosd init $MONIKER --chain-id $CHAINID --home $DATA_DIR
echo "prepare genesis: Allocate genesis accounts"
./bitosd add-genesis-account \
"$(./bitosd keys show $KEY -a --home $DATA_DIR --keyring-backend test)" 1000000000000000000abit,1000000000000000000stake \
--home $DATA_DIR --keyring-backend test
echo "prepare genesis: Sign genesis transaction"
./bitosd gentx $KEY 1000000000000000000stake --keyring-backend test --home $DATA_DIR --keyring-backend test --chain-id $CHAINID
echo "prepare genesis: Collect genesis tx"
./bitosd collect-gentxs --home $DATA_DIR
echo "prepare genesis: Run validate-genesis to ensure everything worked and that the genesis file is setup correctly"
./bitosd validate-genesis --home $DATA_DIR

echo "starting bitos node $i in background ..."
./bitosd start --pruning=nothing --rpc.unsafe \
--keyring-backend test --home $DATA_DIR \
>$DATA_DIR/node.log 2>&1 & disown

echo "started bitos node"
tail -f /dev/null