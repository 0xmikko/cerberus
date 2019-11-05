# Cerberus Wallet
### Two factor authorisation for crypto wallets

Cerberus Wallet is open source project which adds two factor authentification for your crypto wallet. Even if your private key would be stolen, hacker would not be able to transfer your assets without your confirmation.

This application was designed from scratch especially for <a href="https://coinlist.co/build/chainlink/">Chainlink + CoinList Hackathon</a>.

Official site: https://cerberus.ledger-labs.com/

Video presentation: https://www.youtube.com/watch?v=4S8OyUf7cIA

<img src="https://user-images.githubusercontent.com/26343374/68207919-18c81100-ffe1-11e9-8623-23be911b6718.png" width="23%" />&nbsp;&nbsp;
<img src="https://user-images.githubusercontent.com/26343374/68207920-1960a780-ffe1-11e9-87ac-01122a3ef0e4.png" width="23%" />&nbsp;&nbsp;
<img src="https://user-images.githubusercontent.com/26343374/68207923-1960a780-ffe1-11e9-9781-61a1af9c54c3.png" width="23%" />&nbsp;&nbsp;
<img src="https://user-images.githubusercontent.com/26343374/68207924-1960a780-ffe1-11e9-8863-2cd482cb9059.png" width="23%" />

### Problem

Cryptocurrency security is based on the idea that only you have access to the private key. And indeed, today it is impossible to pick up a private key, knowing the public one.

Despite this, for daily transactions, you must store the private key on your computer or smartphone. Often, hackers gain access to a private key using viruses or trojans.

To ensure safety you must follow a lot of rules. One small mistake (for example, accidentally opening a file sent by e-mail) can lead to a complete loss of all crypto assets.

In addition, you cannot verify that your private key is cracked. Some hackers can store it and control your account to get the most money.

### Solution

Two-factor authorization. In order to transfer funds, you need to confirm the transaction using the Cerberus Wallet mobile application within 3-5 minutes after it got into the blockchain. When the time is up, the smart contract checks the confirmation and sends the money only if the transaction has been confirmed.

This approach significantly increases the security of a crypto wallet, since two conditions must be met to transfer money:
Know private key
Confirm transaction using smartphone

In addition, if you receive a notification that someone signed a transaction using your private key, it means that your private key has been hacked. Unlike a regular crypto wallet, you can cancel this transaction and then transfer your money to a secure wallet or account.

### Components & stack

* SmartContract Wallet (Solidity, ChainLink)
* ChainLink external adapter (Golang)
* Backend server (Golang, MongoDB, go-ethereum, gin-gonic, APNS2)
* Mobile Application (React Native)
* Web wallet (ReactJS, Web3)

### SmartContract Wallet
Stack: Solidity, ChainLink

A smart contract for a wallet is developed as simple as possible, which simplifies its audit. The current version of the contract allows you to deposit money on it and transfer them to other accounts.

#### Deploy smartcontract
*address _link* - ChainLink token smartcontract address<br/>
*address _alarmOracle* - ChainLink Alarm Oracle smartcontract address (more: https://docs.chain.link/docs/chainlink-alarm-clock)<br/>
*bytes32 _alarmJobId* - ChainLink Alarm Job Id<br/>
*uint256 _alarmPayment* - Amount in LINK for Alarm Job (usually 1 LINK)<br/>
*address _cerberusOracle* - Cerberus Oracle smartcontract address<br/>
*bytes32 _cerberusJobId* - Cerberus Job ID<br/>
*uint256 _cerberusPayment* - Amount in LINK for Cerberus Job (usually 1 LINK)<br/>

#### Deposit money flow
Everything is simple here - the usual function for receiving funds.

#### Sending money flow

1. To send funds, the user calls method *sendMoney(address _to, uint256 _amount)*

2. The sendMoney method, in turn, emits a *NewPaymentRegistered* event, which is used to monitor transactions.

3. *sendMoney* method also calls ChainLink Alarm Oracle and asks to call the *fulfillConfirmationRequest* method after the specified time (time to sign the transaction). The method also writes transaction parameters to mapping orders by the key that it received during the request (reqId)

![sc_flow_1](https://user-images.githubusercontent.com/26343374/68213193-259e3200-ffec-11e9-8731-a75fbda0e9a4.png)

4. After the specified time has passed, the fulfillConfirmationRequest method is called, which sends a request to Cerberus Oracle to find out if the user has signed the transaction.

![sc_flow2](https://user-images.githubusercontent.com/26343374/68214013-ad387080-ffed-11e9-8eb1-be20fb761db7.png)

5. The oracle calls the *fulfillPaymentRequest* method and passes the result of the confirmation request to the server as the *_data* parameter. If the answer is yes, then the funds are transferred to the specified address.

### External Adapter
Stack: Golang, Bridges framework (https://github.com/linkpoolio/bridges)

The task of the external adapter is to make a GET request to the Cerberus Server and pass the received value back to the oracle. For this, the project uses the Bridges framework, which makes this task elementary.

### Backend
Stack: Golang, MongoDB, go-ethereum, gin-gonic

#### User authorisation / registration
The current version uses simple login / password authorisation using JWT tokens. Authorisation is required in order to transfer the list of user wallets and provide the ability to receive data. Currently, the binding of wallets to the token for push notifications also uses authorisation.

#### Transaction monitoring
The Cerberus Wallet server saves all user wallets and constantly monitors events at connected addresses. To monitor events for each wallet, a separate goroutine is used.

As soon as the NewPaymentRegistered event occurs, the server sends a push notification to all devices connected by the user. The structure of push notifications includes a code that must be entered to confirm the transaction (more simple logic is currently implemented).

As you know, the push notification delivery system is quite reliable and has proven itself in banking applications.

Also, in addition to sending push notifications, the system can send notifications to email clients, via SMS and instant messengers (a number of these functions are under development)

#### Transaction Confirmations

When the monitoring system detects a new transaction, a push notification is sent to the user, and the transaction itself falls into the list of transactions for confirmation.

If the user logged in to the application within the specified time, the system will mark this in the database and confirm it upon request from the Oracle Cerberus.

At the same time, the logic is arranged in such a way that if at least one device rejected the transaction, then it is already impossible to confirm it. This is done to provide greater security.

### Disclaimer

This application is provided "as is" and "with all faults." Me as developer makes no representations or warranties of any kind concerning the safety, suitability, lack of viruses, inaccuracies, typographical errors, or other harmful components of this software. There are inherent dangers in the use of any software, and you are solely responsible for determining whether this software product is compatible with your equipment and other software installed on your equipment. You are also solely responsible for the protection of your equipment and backup of your data, and THE PROVIDER will not be liable for any damages you may suffer in connection with using, modifying, or distributing this software product.

