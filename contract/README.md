# Cerberus Smart Contract

A smart contract for a wallet is developed as simple as possible, which simplifies its audit. The current version of the contract allows you to deposit money on it and transfer them to other accounts.

Implementation of a [Chainlink requesting contract](https://docs.chain.link/docs/create-a-chainlinked-project).

## Requirements

- NPM

## Installation

1. Clone mono repository `git clone https://github.com/MikaelLazarev/cerberus`

2. Go to contracts dir: `cd contract`

3. Install truffle box dependecies

```bash
npm install
```

Or

```bash
yarn install
```

## Test

```bash
npm test
```

## Deploy

If needed, edit the `truffle-config.js` config file to set the desired network to a different port. It assumes any network is running the RPC port on 8545.

```bash
npm run migrate:dev
```

For deploying to live networks, Truffle will use `truffle-hdwallet-provider` for your mnemonic and an RPC URL. Set your environment variables `$RPC_URL` and `$MNEMONIC` before running:

```bash
npm run migrate:live
```

### Deploy parameters
Deploy smartcontract
address _link - ChainLink token smartcontract address
address _alarmOracle - ChainLink Alarm Oracle smartcontract address (more: https://docs.chain.link/docs/chainlink-alarm-clock)
bytes32 _alarmJobId - ChainLink Alarm Job Id
uint256 _alarmPayment - Amount in LINK for Alarm Job (usually 1 LINK)
address _cerberusOracle - Cerberus Oracle smartcontract address
bytes32 _cerberusJobId - Cerberus Job ID
uint256 _cerberusPayment - Amount in LINK for Cerberus Job (usually 1 LINK)

## User flow
