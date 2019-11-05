import React, {useEffect} from 'react';
import {Button, Form, FormControl} from 'react-bootstrap';
import * as reducers from '../store/reducers';
import * as actions from '../store/actions';
import {connect} from 'react-redux';
import * as status from '../store/utils/status';
import {Helmet} from "react-helmet";

function WelcomeScreen({
  Web3,
  accounts,
  gasPrice,
  deployContract,
  contractDeployStatus,
  contractAddress,
  history,
}) {
  // const address = "";
  //
  // const accountsRendered = accounts.map(acc => <>{acc}<br /></>)

  useEffect(() => {
    if (contractDeployStatus === status.STATUS_SUCCESS) {
      console.log('redirect to', contractAddress);
      history.push('/wallet/' + contractAddress + '/');
    }
  }, [contractDeployStatus]);

  const onDeployContract = () => {
    deployContract(accounts[0]);
  };

  if (contractDeployStatus === status.STATUS_LOADING) {
    return (
      <div className="App">
        <Helmet>
          <title>Cerberus Wallet: Deploying</title>
        </Helmet>
        <header className="App-header">
          Deploying new wallet, please wait...
        </header>
        |
      </div>
    );
  }

  return (
    <div className="App">
      <Helmet>
        <title>Welcome to Cerberus Wallet</title>
      </Helmet>
      <header className="App-header">
        Welcome to Cerberus Wallet
        <br />
        <br />
        Please connect existing wallet <br />
        <br />
        <Form>
          <FormControl
            style={{fontSize: 30, marginRight: 20}}
            placeholder={'Your wallet address'}
          />
          <Button
            onClick={onDeployContract}
            style={{fontSize: 30, paddingLeft: 10}}>
            Connect
          </Button>
        </Form>
        <br />
        Or
        <br />
        <Button
          onClick={onDeployContract}
          style={{fontSize: 30, paddingLeft: 10}}>
          Deploy
        </Button>
        <br />
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
  mapDispatchToProps,
)(WelcomeScreen);
