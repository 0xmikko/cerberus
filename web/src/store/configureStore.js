import {createStore, compose, applyMiddleware} from 'redux';
import rootReducer from './reducers';
import thunk from 'redux-thunk';

import {routerMiddleware} from 'react-router-redux';

export default history => {
  const composeEnhancers =
    window.__REDUX_DEVTOOLS_EXTENSION_COMPOSE__ || compose;

  return createStore(
    rootReducer(history),
    composeEnhancers(applyMiddleware(thunk, routerMiddleware(history))),
  );

  // return configureStore;
};
