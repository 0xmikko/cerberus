import * as actionTypes from './actionTypes';
import { RSAA } from 'redux-api-middleware';
import { getFullAPIAddress } from '../utils/api';
import AsyncStorage from '@react-native-community/async-storage';

export const login = (email, password) => {
  return async dispatch => {
    const actionResponse = await dispatch({
      [RSAA]: {
        endpoint: getFullAPIAddress('/auth/login/'),
        method: 'POST',
        body: JSON.stringify({ email, password }),
        headers: { 'Content-Type': 'application/json' },
        types: [
          actionTypes.LOGIN_REQUEST,
          actionTypes.LOGIN_SUCCESS,
          actionTypes.LOGIN_FAILED,
        ],
      },
    });
    if (!actionResponse.error) {
      // Save token in AsyncStorage
      dispatch(saveRefreshTokenInStorage(actionResponse.payload.refresh));
    }
  };
};

export const signup = (email, password) => {
  return async dispatch => {
    const actionResponse = await dispatch({
      [RSAA]: {
        endpoint: getFullAPIAddress('/auth/signup/'),
        method: 'POST',
        body: JSON.stringify({ email, password }),
        headers: { 'Content-Type': 'application/json' },
        types: [
          actionTypes.SIGNUP_REQUEST,
          actionTypes.SIGNUP_SUCCESS,
          actionTypes.SIGNUP_FAILED,
        ],
      },
    });
    if (!actionResponse.error) {
      // Save token in AsyncStorage
      dispatch(saveRefreshTokenInStorage(actionResponse.payload.refresh));
    }
  };
};

export const refreshAccessToken = token => ({
  [RSAA]: {
    endpoint: getFullAPIAddress('/auth/token/refresh/'),
    method: 'POST',
    body: JSON.stringify({ refresh: token }),
    headers: { 'Content-Type': 'application/json' },
    types: [
      actionTypes.TOKEN_REQUEST,
      actionTypes.TOKEN_RECEIVED,
      actionTypes.TOKEN_FAILURE,
    ],
  },
});

export const saveRefreshTokenInStorage = token => {
  console.log('Saving TOKEN', token);
  return async dispatch => {
    AsyncStorage.setItem('refreshToken', token)
      .then(() => {
        dispatch({ type: actionTypes.TOKEN_SAVE_SUCCESS });
      })
      .catch(() => {
        dispatch({ type: actionTypes.TOKEN_SAVE_FAILURE });
      });
  };
};

export const logout = () => {
  return async dispatch => {
    AsyncStorage.removeItem('refreshToken')
      .then(
        await dispatch({ type: actionTypes.LOGOUT }),
        // Clean up profile after logout
        await dispatch({
          type: actionTypes.PROFILE_PREFIX + actionTypes.DETAIL_REQUEST,
          meta: { id: 'user' },
        }),
      )
      .catch(console.log);
  };
};
