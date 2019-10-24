import * as actionTypes from './actionTypes';
import {
  createDataloaderDetailAction,
  createDataloaderListAction,
} from './dataloader';

export const getTransactionsList = createDataloaderListAction(
  '/api/transactions/',
  actionTypes.TRANSACTIONS_PREFIX,
);
export const getTransactionDetails = createDataloaderDetailAction(
  '/api/transactions/:id/',
  actionTypes.TRANSACTIONS_PREFIX,
);
