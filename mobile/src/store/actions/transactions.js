import {RSAA} from 'redux-api-middleware';
import * as actionTypes from './actionTypes';

import {
  createDataloaderDetailAction,
  createDataloaderListAction,
} from './dataloader';
import {getApiById} from '../utils/api';
import {withAuth} from '../reducers';

export const getTransactionsList = createDataloaderListAction(
  '/api/transactions/',
  actionTypes.TRANSACTIONS_PREFIX,
);
export const getTransactionDetails = createDataloaderDetailAction(
  '/api/transactions/:id/',
  actionTypes.TRANSACTIONS_PREFIX,
);

export const confirmTransaction = (id, status, hash) => ({
  [RSAA]: {
    endpoint: getApiById('/api/transactions/:id/confirm/', id),
    method: 'PUT',
    body: JSON.stringify({status}),
    headers: withAuth({'Content-Type': 'application/json'}),
    types: [
      {type: actionTypes.TRANSACTION_CONFIRM_REQUEST, meta: {id, hash}},
      {type: actionTypes.TRANSACTIONS_PREFIX + actionTypes.LIST_SUCCESS , meta: {id, hash}},
      {type: actionTypes.TRANSACTION_CONFIRM_FAILURE, meta: {id, hash}},
    ],
  },
});
