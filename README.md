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

### Components & stack

* Mobile Application (React Native)
* Web wallet (ReactJS, Web3)
* SmartContract Wallet (Solidity, ChainLink)
* ChainLink external adapter (Golang)
* Backend server (Golang, MongoDB, go-ethereum, gin-gonic, APNS2)

### Disclaimer

This application is provided "as is" and "with all faults." Me as developer makes no representations or warranties of any kind concerning the safety, suitability, lack of viruses, inaccuracies, typographical errors, or other harmful components of this software. There are inherent dangers in the use of any software, and you are solely responsible for determining whether this software product is compatible with your equipment and other software installed on your equipment. You are also solely responsible for the protection of your equipment and backup of your data, and THE PROVIDER will not be liable for any damages you may suffer in connection with using, modifying, or distributing this software product.

