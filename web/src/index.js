import React from "react";
import ReactDOM from "react-dom";
import "./index.css";
import App from "./App";
import "./App.css";
import * as serviceWorker from "./serviceWorker";
import {createBrowserHistory } from "history";
import { ConnectedRouter } from "connected-react-router";
import { Provider } from "react-redux";
import { ReactReduxContext } from "react-redux";
import configureStore from "./store/configureStore";

const history = createBrowserHistory({ basename: '/wallet' });
const store = configureStore(history);
console.log(store)
ReactDOM.render(
  <Provider store={store}>
    <ConnectedRouter history={history}>
      <App />
    </ConnectedRouter>
  </Provider>,
  document.getElementById("root")
);

// If you want your app to work offline and load faster, you can change
// unregister() to register() below. Note this comes with some pitfalls.
// Learn more about service workers: https://bit.ly/CRA-PWA
serviceWorker.unregister();
