import * as actionTypes from './actionTypes';
import {RSAA} from 'redux-api-middleware';
import {getApiById} from '../utils/api';
import {withAuth} from '../reducers';

import {
  createDataloaderDetailAction,
  createDataloaderListAction,
} from './dataloader';

export const getAccountsList = createDataloaderListAction(
  '/api/accounts/',
  actionTypes.ACCOUNTS_PREFIX,
);
export const getAccountDetails = createDataloaderDetailAction(
  '/api/accounts/:id/',
  actionTypes.ACCOUNTS_PREFIX,
);
