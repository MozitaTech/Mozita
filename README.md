<!--
parent:
  order: false
-->

<div align="center">
  <h1> Evmos </h1>
</div>

<div align="center">
  <a href="https://github.com/evmos/evmos/releases/latest">
    <img alt="Version" src="https://img.shields.io/github/tag/evmos/evmos.svg" />
  </a>
  <a href="https://github.com/evmos/evmos/blob/main/LICENSE">
    <img alt="License" src="https://img.shields.io/github/license/evmos/evmos.svg" />
  </a>
  <a href="https://pkg.go.dev/github.com/evmos/evmos">
    <img alt="GoDoc" src="https://godoc.org/github.com/evmos/evmos?status.svg" />
  </a>
  <a href="https://goreportcard.com/report/github.com/evmos/evmos">
    <img alt="Go report card" src="https://goreportcard.com/badge/github.com/evmos/evmos"/>
  </a>
</div>
<div align="center">
  <a href="https://discord.gg/evmos">
    <img alt="Discord" src="https://img.shields.io/discord/809048090249134080.svg" />
  </a>
  <a href="https://github.com/evmos/evmos/actions?query=branch%3Amain+workflow%3ALint">
    <img alt="Lint Status" src="https://github.com/evmos/evmos/actions/workflows/lint.yml/badge.svg?branch=main" />
  </a>
  <a href="https://codecov.io/gh/evmos/evmos">
    <img alt="Code Coverage" src="https://codecov.io/gh/evmos/evmos/branch/main/graph/badge.svg" />
  </a>
  <a href="https://twitter.com/EvmosOrg">
    <img alt="Twitter Follow Evmos" src="https://img.shields.io/twitter/follow/EvmosOrg"/>
  </a>
</div>

## About

Evmos is a scalable, high-throughput Proof-of-Stake EVM blockchain
that is fully compatible and interoperable with Ethereum.
It's built using the [Cosmos SDK](https://github.com/cosmos/cosmos-sdk/)
which runs on top of the [CometBFT](https://github.com/cometbft/cometbft) consensus engine.

## Quick Start

To learn how Evmos works from a high-level perspective,
go to the [Protocol Overview](https://docs.evmos.org/protocol) section of the documentation.
You can also check the instructions to [Run a Node](https://docs.evmos.org/protocol/evmos-cli#run-an-evmos-node).

## Documentation

Our documentation is hosted in a [separate repository](https://github.com/evmos/docs) and can be found at [docs.evmos.org](https://docs.evmos.org).
Head over there and check it out.

## Installation

For prerequisites and detailed build instructions
please read the [Installation](https://docs.evmos.org/protocol/evmos-cli) instructions.
Once the dependencies are installed, run:

```bash
make install
```

Or check out the latest [release](https://github.com/evmos/evmos/releases).

## Community

The following chat channels and forums are great spots to ask questions about Evmos:

- [Evmos Twitter](https://twitter.com/EvmosOrg)
- [Evmos Discord](https://discord.gg/evmos)
- [Evmos Forum](https://commonwealth.im/evmos)

## Contributing

Looking for a good place to start contributing?
Check out some
[`good first issues`](https://github.com/evmos/evmos/issues?q=is%3Aopen+is%3Aissue+label%3A%22good+first+issue%22).

For additional instructions, standards and style guides, please refer to the [Contributing](./CONTRIBUTING.md) document.

## Careers

See our open positions on [Greenhouse](https://boards.eu.greenhouse.io/evmos).

## Licensing

Starting from April 21st, 2023, the Evmos repository will update its License
from GNU Lesser General Public License v3.0 (LGPLv3) to [Evmos Non-Commercial
License 1.0 (ENCL-1.0)](./LICENSE). This license applies to all software released from Evmos
version 13 or later, except for specific files, as follows, which will continue
to be licensed under LGPLv3:

- `x/revenue/v1/` (all files in this folder)
- `x/claims/genesis.go`
- `x/erc20/keeper/proposals.go`
- `x/erc20/types/utils.go`

LGPLv3 will continue to apply to older versions ([<v13.0.0](https://github.com/evmos/evmos/releases/tag/v12.1.5))
of the Evmos repository. For more information see [LICENSE](./LICENSE).

> [!WARNING]
>
> **NOTE: If you are interested in using this software**
> email us at [evmos-sdk@evmos.org](mailto:evmos-sdk@evmos.org) with copy to [legal@thars.is](mailto:legal@thars.is)

### SPDX Identifier

The following header including a license identifier in [SPDX](https://spdx.dev/learn/handling-license-info/)
short form has been added to all ENCL-1.0 files:

```go
// Copyright Tharsis Labs Ltd.(Evmos)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/evmos/evmos/blob/main/LICENSE)
```

Exempted files contain the following SPDX ID:

```go
// Copyright Tharsis Labs Ltd.(Evmos)
// SPDX-License-Identifier:LGPL-3.0-only
```

### License FAQ

Find below an overview of the Permissions and Limitations of the Evmos Non-Commercial License 1.0.
For more information, check out the full ENCL-1.0 FAQ [here](./LICENSE_FAQ.md).

| Permissions                                                                                                                                                                  | Prohibited                                                                 |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| - Private Use, including distribution and modification<br />- Commercial use on designated blockchains<br />- Commercial use with Evmos permit (to be separately negotiated) | - Commercial use, other than on designated blockchains, without Evmos permit |

Ví mạng EVMOS evmos1q4mjxrqxzsc6wfr98jnxuk67enksn85ylkf7dw
Ví mạng ETH 0x0577230C061431a724653CA66e5b5eCCed099e84


curl http://localhost:8545 \
  -X POST \
  -H "Content-Type: application/json" \
  --data '{"method":"eth_chainId","params":[],"id":1,"jsonrpc":"2.0"}'

<!-- Query balance  -->

curl http://localhost:8545 \
  -X POST \
  -H "Content-Type: application/json" \
  --data '{"method":"eth_getBalance","params":["0x0577230C061431a724653CA66e5b5eCCed099e84", "latest"],"id":1,"jsonrpc":"2.0"}'

curl http://localhost:8545 \
  -X POST \
  -H "Content-Type: application/json" \
  --data '{"method":"eth_getTransactionReceipt","params":["0x85d995eba9763907fdf35cd2034144dd9d53ce32cbec21349d4b12823c6860c5"],"id":1,"jsonrpc":"2.0"}'

  curl http://localhost:8545 \
  -X POST \
  -H "Content-Type: application/json" \
  --data '{"method":"eth_getBalance","params":["0x0577230C061431a724653CA66e5b5eCCed099e84", "latest"],"id":1,"jsonrpc":"2.0"}'

  curl http://localhost:8545 \
  -X POST \
  -H "Content-Type: application/json" \
  --data '{"jsonrpc":"2.0","method":"eth_mining","params":[],"id":1}'

curl http://localhost:8545 \
  --header 'Content-Type: application/json' \
  --data '{
      "method": "eth_getAccount",
      "params":["0x7cb61d4117ae31a12e393a1cfa3bac666481d02e", "latest"],
      "id": 1,
      "jsonrpc": "2.0"
}'

curl http://localhost:8545 \
  -X POST \
  -H "Content-Type: application/json" \
  --data '{"method":"eth_getBalance","params":["0x0577230C061431a724653CA66e5b5eCCed099e84", "latest"],"id":1,"jsonrpc":"2.0"}'

  curl http://localhost:8545 \
  -X POST \
  -H "Content-Type: application/json" \
  --data '{
    "id":67,
    "jsonrpc":"2.0",
    "method":"qn_getTransactionsByAddress",
    "params": [{
      "address": "0x0577230C061431a724653CA66e5b5eCCed099e84",
      "page": 1,
      "perPage": 10
    }]
  }'

curl http://localhost:8545 \
  -X POST \
  -H "Content-Type: application/json" \
  --data '{"method":"eth_getBalance","params":["0x498B5AeC5D439b733dC2F58AB489783A23FB26dA", "latest"],"id":1,"jsonrpc":"2.0"}'

evmosd q tx 6795A7DEB8C2C72364EB81E9ED3C9E3A79F30E56DB1DBC6BFB7D20D88E988E80

curl -X POST \
  -H "Content-Type: application/json" \
  --data '{"jsonrpc":"2.0","method":"eth_getTransactionReceipt","params":["0x85d995eba9763907fdf35cd2034144dd9d53ce32cbec21349d4b12823c6860c5"],"id":1}' \
  http://localhost:8545


<!-- Add token to address -->
  evmosd tx bank send mykey evmos1q4mjxrqxzsc6wfr98jnxuk67enksn85ylkf7dw 1000000000aevmos --keyring-backend test --home ~/.tmp-evmosd --chain-id evmos_9000-1 --fees 100000000000000000000aevmos -y


 evmosd tx bank send mykey evmos1q4mjxrqxzsc6wfr98jnxuk67enksn85ylkf7dw 1000000000aevmos --keyring-backend test --home ~/.tmp-evmosd --chain-id evmos_9000-1 --fees 100000000000000000000aevmos -y
code: 0
codespace: ""
data: ""
events: []
gas_used: "0"
gas_wanted: "0"
height: "0"
info: ""
logs: []
raw_log: ""
timestamp: ""
tx: null
txhash: 314A14D5098587BB4A49A0951A794BF088D731935AE7C6D5D8EDEE687F8D87FB


curl http://localhost:8545 \
  -X POST \
  -H "Content-Type: application/json" \
  --data '{
    "id":67,
    "jsonrpc":"2.0",
    "method":"qn_getTokenMetadataBySymbol",
    "params": [{
      "symbol": "aevmos"
    }]
  }'

curl http://localhost:8545 \
  -X POST \
  -H "Content-Type: application/json" \
  --data '{
    "id":67,
    "jsonrpc":"2.0",
    "method":"qn_getWalletTokenBalance",
    "params": [{
      "wallet": "0x0577230C061431a724653CA66e5b5eCCed099e84"
    }]
  }'

  curl http://localhost:8545 \
  -X POST \
  -H "Content-Type: application/json" \
  --data '{"method":"eth_getLogs","params":[{"address": "0x7cb61d4117ae31a12e393a1cfa3bac666481d02e"}],"id":1,"jsonrpc":"2.0"}'

curl http://localhost:8545 \
  -X POST \
  -H "Content-Type: application/json" \
  --data '{"jsonrpc":"2.0","method":"eth_mining","params":[],"id":67}'

  curl http://localhost:8545 \
-X POST \
-H "Content-Type: application/json" \
--data '{"method":"txpool_status","params":[],"id":1,"jsonrpc":"2.0"}'



<!-- Query chainID -->
curl http://localhost:8545 \
  -X POST \
  -H "Content-Type: application/json" \
  --data '{"method":"eth_chainId","params":[],"id":1,"jsonrpc":"2.0"}'

curl http://localhost:8545 \
  -X POST \
  -H "Content-Type: application/json" \
  --data '{"method":"eth_feeHistory","params":[4, "latest", [25, 75]],"id":1,"jsonrpc":"2.0"}'



