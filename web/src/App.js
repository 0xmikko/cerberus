import React, { useEffect } from "react";
import { Route, Switch, Redirect } from "react-router";
import { Container } from "react-bootstrap";
import { withRouter } from "react-router";
import { connect } from "react-redux";
import * as actions from "./store/actions";
import * as reducers from "./store/reducers";

import WelcomeScreen from "./screens/WelcomeScreen";
import WalletScreen from "./screens/WalletScreen";
import NoWeb3FoundScreen from './screens/NoWeb3FoundScreen';


export function App({getWeb3, Web3}) {

  useEffect(() => {
   if (!Web3) { getWeb3(); }
  }, [Web3]);

  if (!Web3) {
      return <NoWeb3FoundScreen />;
  }
  return (
    <Switch>
      <Route path="/wallet/:id/" component={WalletScreen} />
      <Route path="/welcome/" component={WelcomeScreen} />
    </Switch>
  );
}

const mapStateToProps = state => ({
  Web3: reducers.Web3(state)
});

const mapDispatchToProps = dispatch => ({
    getWeb3: () => dispatch(actions.getWeb3())
});

export default withRouter(
  connect(
    mapStateToProps,
    mapDispatchToProps
  )(App)
);
