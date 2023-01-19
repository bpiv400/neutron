#!/bin/bash

BINARY=${BINARY:-neutrond}
CHAIN_DIR=./data
CHAINID=${CHAINID:-test-1}
STAKEDENOM=${STAKEDENOM:-stake}

ADMIN_ADDRESS=neutron1m9l358xunhhwds0568za49mzhvuxx9ux8xafx2
DAO_CONTRACT=/opt/neutron/contracts/dao/cwd_core.wasm
PRE_PROPOSAL_CONTRACT=/opt/neutron/contracts/dao/cwd_pre_propose_single.wasm
PROPOSAL_CONTRACT=/opt/neutron/contracts/dao/cwd_proposal_single.wasm
VOTING_REGISTRY_CONTRACT=/opt/neutron/contracts/dao/neutron_voting_registry.wasm
VAULT_CONTRACT=/opt/neutron/contracts/dao/neutron_vault.wasm
PROPOSAL_MULTIPLE_CONTRACT=/opt/neutron/contracts/dao/cwd_proposal_multiple.wasm
PRE_PROPOSAL_MULTIPLE_CONTRACT=/opt/neutron/contracts/dao/cwd_pre_propose_multiple.wasm
TREASURY_CONTRACT=/opt/neutron/contracts/dao/neutron_treasury.wasm

echo "Add consumer section..."
$BINARY add-consumer-section --home $CHAIN_DIR/$CHAINID

echo "Initializing dao contract in genesis..."

function store_binary() {
  CONTRACT_BINARY_PATH=$1
  $NEUTROND_BINARY add-wasm-message store "$CONTRACT_BINARY_PATH" --output json --run-as ${ADMIN_ADDRESS} --keyring-backend=test --home $CHAIN_DIR/$CHAINID_1
  echo $(jq -r "[.app_state.wasm.gen_msgs[] | select(.store_code != null)] | length" $CHAIN_DIR/$CHAINID_1/config/genesis.json)
}

# Upload the dao contracts

VAULT_CONTRACT_BINARY_ID=$(store_binary ${VAULT_CONTRACT})
DAO_CONTRACT_BINARY_ID=$(store_binary ${DAO_CONTRACT})
PROPOSAL_CONTRACT_BINARY_ID=$(store_binary ${PROPOSAL_CONTRACT})
VOTING_REGISTRY_CONTRACT_BINARY_ID=$(store_binary ${VOTING_REGISTRY_CONTRACT})
PRE_PROPOSAL_CONTRACT_BINARY_ID=$(store_binary ${PRE_PROPOSAL_CONTRACT})
PROPOSAL_MULTIPLE_CONTRACT_BINARY_ID=$(store_binary ${PROPOSAL_MULTIPLE_CONTRACT})
PRE_PROPOSAL_MULTIPLE_CONTRACT_BINARY_ID=$(store_binary ${PRE_PROPOSAL_MULTIPLE_CONTRACT})
TREASURY_CONTRACT_BINARY_ID=$(store_binary ${TREASURY_CONTRACT})
DISTRIBUTION_CONTRACT_BINARY_ID=$(store_binary ${DISTRIBUTION_CONTRACT})

# PRE_PROPOSE_INIT_MSG will be put into the PROPOSAL_SINGLE_INIT_MSG and PROPOSAL_MULTIPLE_INIT_MSG
PRE_PROPOSE_INIT_MSG='{
   "deposit_info":{
      "denom":{
         "token":{
            "denom":{
               "native":"stake"
            }
         }
      },
     "amount": "1000",
     "refund_policy":"always"
   },
   "open_proposal_submission":false
}'
PRE_PROPOSE_INIT_MSG_BASE64=$(echo ${PRE_PROPOSE_INIT_MSG} | base64 | tr -d "\n")

# -------------------- PROPOSE-SINGLE { PRE-PROPOSE } --------------------

PROPOSAL_SINGLE_INIT_MSG='{
   "allow_revoting":false,
   "pre_propose_info":{
      "ModuleMayPropose":{
         "info":{
            "code_id": '"${PRE_PROPOSAL_CONTRACT_BINARY_ID}"',
            "msg": "'"${PRE_PROPOSE_INIT_MSG_BASE64}"'",
            "label":"neutron"
         }
      }
   },
   "only_members_execute":false,
   "max_voting_period":{
      "time":604800
   },
   "close_proposal_on_execution_failure":false,
   "threshold":{
      "threshold_quorum":{
         "quorum":{
            "percent":"0.20"
         },
         "threshold":{
            "majority":{

            }
         }
      }
   }
}'
PROPOSAL_SINGLE_INIT_MSG_BASE64=$(echo ${PROPOSAL_SINGLE_INIT_MSG} | base64 | tr -d "\n")

# -------------------- PROPOSE-MULTIPLE { PRE-PROPOSE } --------------------

PROPOSAL_MULTIPLE_INIT_MSG='{
   "allow_revoting":false,
   "pre_propose_info":{
      "ModuleMayPropose":{
         "info":{
            "code_id": '"${PRE_PROPOSAL_MULTIPLE_CONTRACT_BINARY_ID}"',
            "msg": "'"${PRE_PROPOSE_INIT_MSG_BASE64}"'",
            "label":"neutron"
         }
      }
   },
   "only_members_execute":false,
   "max_voting_period":{
      "time":604800
   },
   "close_proposal_on_execution_failure":false,
   "voting_strategy":{
     "single_choice": {
        "quorum": {
          "majority": {
          }
        }
     }
   }
}'
PROPOSAL_MULTIPLE_INIT_MSG_BASE64=$(echo ${PROPOSAL_MULTIPLE_INIT_MSG} | base64 | tr -d "\n")

VOTING_REGISTRY_INIT_MSG='{
  "manager": null,
  "owner": null,
  "voting_vault": "neutron14hj2tavq8fpesdwxxcu44rty3hh90vhujrvcmstl4zr3txmfvw9s5c2epq"
}'
VOTING_REGISTRY_INIT_MSG_BASE64=$(echo ${VOTING_REGISTRY_INIT_MSG} | base64 | tr -d "\n")

INIT='{
  "denom":"stake",
  "description": "based neutron vault"
}'
DAO_INIT='{
  "description": "basic neutron dao",
  "name": "Neutron",
  "initial_items": null,
  "proposal_modules_instantiate_info": [
    {
      "code_id": '"${PROPOSAL_CONTRACT_BINARY_ID}"',
      "label": "DAO_Neutron_cw-proposal-single",
      "msg": "'"${PROPOSAL_SINGLE_INIT_MSG_BASE64}"'"
    },
    {
      "code_id": '"${PROPOSAL_MULTIPLE_CONTRACT_BINARY_ID}"',
      "label": "DAO_Neutron_cw-proposal-multiple",
      "msg": "'"${PROPOSAL_MULTIPLE_INIT_MSG_BASE64}"'"
    }
  ],
  "voting_registry_module_instantiate_info": {
    "code_id": '"${VOTING_REGISTRY_CONTRACT_BINARY_ID}"',
    "label": "DAO_Neutron_voting_registry",
    "msg": "'"${VOTING_REGISTRY_INIT_MSG_BASE64}"'"
  }
}'

# TODO: properly initialize treasury
DISTRIBUTION_CONTRACT_ADDRESS="neutron1vhndln95yd7rngslzvf6sax6axcshkxqpmpr886ntelh28p9ghuq56mwja"
TREASURY_INIT="$(printf '{
  "main_dao_address": "%s",
  "security_dao_address": "%s",
  "denom": "stake",
  "distribution_rate": "0",
  "min_period": 10,
  "distribution_contract": "%s",
  "reserve_contract": "%s",
  "vesting_denominator": "1"
}' "$ADMIN_ADDRESS" "$ADMIN_ADDRESS" "$DISTRIBUTION_CONTRACT_ADDRESS" "$ADMIN_ADDRESS")"

DISTRIBUTION_INIT="$(printf '{
                           "main_dao_address": "%s",
                           "security_dao_address": "%s",
                           "denom": "stake"
}' "$ADMIN_ADDRESS" "$ADMIN_ADDRESS")"

echo "Instantiate contracts"
$NEUTROND_BINARY add-wasm-message instantiate-contract "$VAULT_CONTRACT_BINARY_ID" "$INIT" --run-as ${ADMIN_ADDRESS} --admin ${ADMIN_ADDRESS}  --label "DAO_Neutron_voting_vault"  --home $CHAIN_DIR/$CHAINID_1
$NEUTROND_BINARY add-wasm-message instantiate-contract "$DAO_CONTRACT_BINARY_ID" "$DAO_INIT" --run-as ${ADMIN_ADDRESS} --admin ${ADMIN_ADDRESS}  --label "DAO"  --home $CHAIN_DIR/$CHAINID_1
$NEUTROND_BINARY add-wasm-message instantiate-contract "$TREASURY_CONTRACT_BINARY_ID" "$TREASURY_INIT" --run-as ${ADMIN_ADDRESS} --admin ${ADMIN_ADDRESS} --label "Treasury" --home $CHAIN_DIR/$CHAINID_1
$NEUTROND_BINARY add-wasm-message instantiate-contract "$DISTRIBUTION_CONTRACT_BINARY_ID" "$DISTRIBUTION_INIT" --run-as ${ADMIN_ADDRESS} --admin ${ADMIN_ADDRESS} --label "Distribution" --home $CHAIN_DIR/$CHAINID_1

sed -i -e 's/\"admins\":.*/\"admins\": [\"neutron1nc5tatafv6eyq7llkr2gv50ff9e22mnf70qgjlv737ktmt4eswrqcd0mrx\"]/g' $CHAIN_DIR/$CHAINID/config/genesis.json

