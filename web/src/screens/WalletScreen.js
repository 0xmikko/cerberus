import React, { useState } from 'react';
import {connect} from "react-redux";
import {Button, FormControl} from 'react-bootstrap';


import * as reducers from "../store/reducers";
import * as actions from "../store/actions";
import {QRCode} from "react-qr-svg";
import * as status from '../store/utils/status';

function WalletScreen({match: {params: {id}}, send, deposit, accounts, paymentStatus, Web3}) {

    console.log(id)

    const [depositAmount, setDepositAmount] = useState(undefined);
    const [sendRecipient, setSendRecipient] = useState(undefined);
    const [sendAmount, setSendAmount] = useState(undefined);

    if (!Web3.utils.isAddress(id)) {
        return <div className="App">
            <header className="App-header">
                <div>Incorrect contract address: <br /><br />
                    {id}
                </div>
            </header>
        </div>
    }

    const depositMoney = () => {
        console.log(depositAmount)
        if (!depositAmount) {
            alert("Please, provide amount of money you want to deposit");
            return
        }
        deposit(id, accounts[0], depositAmount);
    }
    const sendMoney = () => {
        console.log(sendRecipient, sendAmount)
        if (!sendRecipient || !sendAmount) {
            alert("Please, provide recipient and amount");
            return
        }
        send(id, accounts[0], sendRecipient, sendAmount);
    }

    if (paymentStatus === status.STATUS_LOADING) {
        return (
            <div className="App">
                <header className="App-header">
                    Operations is in progress<br />
                    Please, check your metamask plugin, probably it waits your signing.
                </header>
            </div>
        );
    }

    return (
        <div className="App">
        <header className="App-header">
            <div>Contract <br /><br />
            <a href={`https://ropsten.etherscan.io/address/${id}#events`} style={{ color: 'white'}}>{id}</a></div>
            <br />
            <FormControl style={{fontSize: 30, marginRight: 20, width: '70%'}}
                                      placeholder={"Amount"}
                                      value={depositAmount}
                                      onChange={w => setDepositAmount(w.target.value)}

                    /><br />
                        <Button onClick={depositMoney}  style={{fontSize: 30, paddingLeft: 10 }}>Deposit Money</Button><br /><br />
                    <FormControl style={{fontSize: 30, marginRight: 20, width: '70%'}} placeholder={"To"}
                                     value={sendRecipient}
                                     onChange={w => setSendRecipient(w.target.value)}

                    /><br />
                        <FormControl style={{fontSize: 30, marginRight: 20, width: '70%'}} placeholder={"Amount"}
                                     value={sendAmount}
                                     onChange={w => setSendAmount(w.target.value)}

                        /><br />
                        <Button onClick={sendMoney}  style={{fontSize: 30, paddingLeft: 10}}>Send Money</Button><br /><br />




            <QRCode
                bgColor="#FFFFFF"
                fgColor="#000000"
                level="Q"
                style={{ width: 200 }}
                value={id}
            />
        </header>
        </div>


    )


}

const mapStateToProps = state => ({
    Web3: reducers.Web3(state),
    accounts: reducers.accounts(state),
    paymentStatus: reducers.paymentStatus(state),
});

const mapDispatchToProps = dispatch => ({
    deposit: (contract, from, amount) => dispatch(actions.depositMoney(contract, from, amount)),
    send: (contract, from, to, amount) => dispatch(actions.sendMoney(contract, from, to, amount))
});

export default connect(
    mapStateToProps,
    mapDispatchToProps
)(WalletScreen);
