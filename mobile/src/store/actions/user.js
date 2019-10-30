import * as actionTypes from './actionTypes';
import {
  createDataloaderDetailAction,
  createDataloaderCreateUpdateDataAction,
} from './dataloader';
import {RSAA} from 'redux-api-middleware';
import {getApiById} from '../utils/api';
import {withAuth} from '../reducers';

// Get User Prorile
export const getProfile = createDataloaderDetailAction(
  '/api/user/',
  actionTypes.PROFILE_PREFIX,
);
export const updateProfile = createDataloaderCreateUpdateDataAction(
  '/api/user/',
  actionTypes.PROFILE_PREFIX,
);

export const registerToken = (token, hash) => ({
  [RSAA]: {
    endpoint: getApiById('/api/user/register/ios/'),
    method: 'POST',
    body: JSON.stringify({token}),
    headers: withAuth({'Content-Type': 'application/json'}),
    types: [
      {type: actionTypes.PROFILE_PREFIX + '_REQUEST', meta: { hash}},
      {type: actionTypes.PROFILE_PREFIX + '_SUCCESS', meta: { hash}},
      {type: actionTypes.PROFILE_PREFIX + '_FAILURE', meta: { hash}},
    ],
  },
});

export const saveAPNToken = token => ({
  type: actionTypes.PROFILE_PREFIX_SAVE_TOKEN,
  payload: token,
});
