import React, { useState } from 'react';
import logo from './logo.svg';
import { Button } from 'react-bootstrap';
import './App.css';
import CerberusWalletContract from './contracts/CerberusWallet.json';
import { QRCode } from "react-qr-svg";
import getWeb3 from "./utils/getweb3";

function App() {

    const [web3, setWeb3] = useState(null);
    const [account, setAccount] = useState("");
    const [address, setAddress] = useState(null);

    const [gasPrice, setGasPrice] = useState(0);
    getWeb3()
        .then(web3 => {
            setWeb3(web3);
            web3.eth.getAccounts().then(accounts =>
                setAccount(accounts)
            )
            web3.eth.getGasPrice().then(gasPrice => setGasPrice(gasPrice))
            }

        )

    const deploySmartContract = () => {

        const contract = new web3.eth.Contract(CerberusWalletContract.abi);
        const p = contract.deploy({
            data: CerberusWalletContract.bytecode,
            arguments: [
                "0x20fE562d797A42Dcb3399062AE9546cd06f63280",
                "0xc99B3D447826532722E41bc36e644ba3479E4365",
                web3.utils.fromAscii("2ebb1c1a4b1e4229adac24ee0b5f784f"),
                web3.utils.fromAscii("100000000000000000"),
                "0x83F00b902cbf06E316C95F51cbEeD9D2572a349a",
                web3.utils.fromAscii("d4b02e3a2c354111911739c5dd3264a9"),
                web3.utils.fromAscii("100000000000000000")
            ],
        })
        p.estimateGas().then(e => console.log(e))
        p.send({
            from: account[0],
            gas: 2500000,
            gasPrice: gasPrice*2,
        }, function(error, transactionHash){
            console.log(transactionHash)
        })
        .on('error', function(error){ console.log("Error", error) })
        .on('transactionHash', function(transactionHash) { console.log("THASH", transactionHash)})
        .on('receipt', function(receipt){
            console.log(receipt.contractAddress) // contains the new contract address
        })
        .on('confirmation', function(confirmationNumber, receipt) {
            console.log(confirmationNumber) })
        .then(function (newContractInstance) {
            console.log(newContractInstance.options.address) // instance with the new contract address
        });
    }

    return (
    <div className="App">
      <header className="App-header">
        Cerberus Wallet
          <br /><br />
          { account }<br />
          Contract: { address }<br/>
          GasPrice: {gasPrice}
          <Button onClick={deploySmartContract}>Hello</Button>
          <QRCode
              bgColor="#FFFFFF"
              fgColor="#000000"
              level="Q"
              style={{ width: 256 }}
              value={address || "0"}
          />
      </header>

    </div>
  );
}

export default App;
