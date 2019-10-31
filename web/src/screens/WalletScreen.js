import React from 'react';
import {connect} from "react-redux";
import { Button } from 'react-bootstrap';


import * as reducers from "../store/reducers";
import * as actions from "../store/actions";


function WalletScreen({match: {params: {id}}, send, deposit, accounts}) {

    console.log(id)

    const sendMoney = () => {
        send(id,
            accounts[0],
            "0xB82710912a79D362F5F39eEA66e928E77655c445",
            "10000000");
    }
    return (
        <>
            <div>Contract {id}</div>
            <Button onClick={sendMoney}>Send Money</Button>
        </>


    )


}

const mapStateToProps = state => ({
    Web3: reducers.Web3(state),
    accounts: reducers.accounts(state),


});

const mapDispatchToProps = dispatch => ({
    deposit: (contract, amount) => dispatch(actions.depositMoney(contract, amount)),
    send: (contract, from, to, amount) => dispatch(actions.sendMoney(contract, from, to, amount))
});

export default connect(
    mapStateToProps,
    mapDispatchToProps
)(WalletScreen);
