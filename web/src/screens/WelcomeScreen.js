import React, { useState } from 'react';
import { Button } from 'react-bootstrap';
import { QRCode } from "react-qr-svg";
import * as reducers from "../store/reducers";
import * as actions from "../store/actions";
import {withRouter} from "react-router";
import {connect} from "react-redux";

function WelcomeScreen({Web3, accounts, gasPrice, deployContract}) {

    const address = "";

    const accountsRendered = accounts.map(acc => <>{acc}<br /></>)

    return (
        <div className="App">
            <header className="App-header">
                Cerberus Wallet
                <br /><br />
                { accountsRendered }
                Contract: { address }<br/>
                <Button onClick={() => deployContract(accounts[0])}>Deploy</Button><br />
                <QRCode
                    bgColor="#FFFFFF"
                    fgColor="#000000"
                    level="Q"
                    style={{ width: 256 }}
                    value={accounts[0] || "0"}
                />
            </header>

        </div>
    );
}

const mapStateToProps = state => ({
    Web3: reducers.Web3(state),
    accounts: reducers.accounts(state),


});

const mapDispatchToProps = dispatch => ({
    deployContract: from => dispatch(actions.deployContract(from)),
});

export default connect(
        mapStateToProps,
        mapDispatchToProps
    )(WelcomeScreen);
