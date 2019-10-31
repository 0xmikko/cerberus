import * as actionTypes from './actionTypes';
import {RSAA} from 'redux-api-middleware';
import {getApiById} from '../utils/api';
import {withAuth} from '../reducers';

import {
  createDataloaderDetailAction,
  createDataloaderListAction,
  createDataloaderCreateUpdateDataAction,
} from './dataloader';

export const getAccountsList = createDataloaderListAction(
  '/api/accounts/',
  actionTypes.ACCOUNTS_PREFIX,
);

export const getAccountDetails = createDataloaderDetailAction(
  '/api/accounts/:id/',
  actionTypes.ACCOUNTS_PREFIX,
);

export const createAccount = (account, hash) => {
  console.log(account)
  return {
    [RSAA]: {
    endpoint: getApiById('/api/accounts/'),
        method: 'POST',
        body: JSON.stringify(account),
        headers: withAuth({'Content-Type': 'application/json'}),
        types: [
      {type: actionTypes.ACCOUNTS_PREFIX + actionTypes.UPLOAD_REQUEST, meta: {account, hash}},
      {type: actionTypes.ACCOUNTS_PREFIX + actionTypes.LIST_SUCCESS, meta: {account, hash}},
      {type: actionTypes.ACCOUNTS_PREFIX + actionTypes.UPLOAD_FAILURE, meta: {account, hash}},
    ],
  },
  }
};
