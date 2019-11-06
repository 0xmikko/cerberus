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
*address _link* - ChainLink token smartcontract address<br/>
*address _alarmOracle* - ChainLink Alarm Oracle smartcontract address (more: https://docs.chain.link/docs/chainlink-alarm-clock)<br/>
*bytes32 _alarmJobId* - ChainLink Alarm Job Id<br/>
*uint256 _alarmPayment* - Amount in LINK for Alarm Job (usually 1 LINK)<br/>
*address _cerberusOracle* - Cerberus Oracle smartcontract address<br/>
*bytes32 _cerberusJobId* - Cerberus Job ID<br/>
*uint256 _cerberusPayment* - Amount in LINK for Cerberus Job (usually 1 LINK)<br/>

## User flow
The current version of the contract allows you to deposit money on it and transfer them to other accounts.

### Deposit money flow
Everything is simple here - the usual function for receiving funds.

### Sending money flow

1. To send funds, the user calls method *sendMoney(address _to, uint256 _amount)*

2. The sendMoney method, in turn, emits a *NewPaymentRegistered* event, which is used to monitor transactions.

3. *sendMoney* method also calls ChainLink Alarm Oracle and asks to call the *fulfillConfirmationRequest* method after the specified time (time to sign the transaction). The method also writes transaction parameters to mapping orders by the key that it received during the request (reqId)

![sc_flow_1](https://user-images.githubusercontent.com/26343374/68213193-259e3200-ffec-11e9-8731-a75fbda0e9a4.png)

4. After the specified time has passed, the fulfillConfirmationRequest method is called, which sends a request to Cerberus Oracle to find out if the user has signed the transaction.

![sc_flow2](https://user-images.githubusercontent.com/26343374/68214013-ad387080-ffed-11e9-8eb1-be20fb761db7.png)

5. The oracle calls the *fulfillPaymentRequest* method and passes the result of the confirmation request to the server as the *_data* parameter. If the answer is yes, then the funds are transferred to the specified address.

## Backlog
[] Add cold wallet account which has rights to transfer money without confirmation, but this transaction could be cancelled using Cerberus Wallet app
[] Cover tests
