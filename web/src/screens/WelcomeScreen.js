import React, {useEffect, useState} from 'react';
import { Button } from 'react-bootstrap';
import { QRCode } from "react-qr-svg";
import * as reducers from "../store/reducers";
import * as actions from "../store/actions";
import {withRouter} from "react-router";
import {connect} from "react-redux";
import * as status from '../store/utils/status';

function WelcomeScreen({Web3, accounts, gasPrice, deployContract, contractDeployStatus, contractAddress, history}) {

    const address = "";

    const accountsRendered = accounts.map(acc => <>{acc}<br /></>)

    useEffect(() => {
        if (contractDeployStatus === status.STATUS_SUCCESS) {
            console.log("redirect to", contractAddress)
            history.push('/wallet/' + contractAddress + '/')

        }

    }, [contractDeployStatus]);


    const onDeployContract = () => {
        deployContract(accounts[0])
    }

    return (
        <div className="App">
            <header className="App-header">
                Cerberus Wallet
                <br /><br />
                { accountsRendered }
                Contract: { address }<br/>
                <Button onClick={onDeployContract}>Deploy</Button><br />

            </header>

        </div>
    );
}

const mapStateToProps = state => ({
    Web3: reducers.Web3(state),
    accounts: reducers.accounts(state),
    contractDeployStatus: reducers.contractDeployStatus(state),
    contractAddress: reducers.contractAddress(state),

});

const mapDispatchToProps = dispatch => ({
    deployContract: from => dispatch(actions.deployContract(from)),
});

export default connect(
        mapStateToProps,
        mapDispatchToProps
    )(WelcomeScreen);
