import jwtDecode from 'jwt-decode';
import { updateState } from '../utils/updateState';
import * as auth from '../actions/actionTypes';
import * as status from '../utils/status';

const initialState = {
  access: undefined,
  refresh: undefined,
  error: '',
  authStatus: status.STATUS_UPDATE_NEEDED,
  tokenStatus: status.STATUS_UPDATE_NEEDED,
};

export default (state = initialState, action) => {
  console.log(action);
  switch (action.type) {
    case auth.LOGIN_REQUEST:
    case auth.SIGNUP_REQUEST:
      return updateState(state, {
        authStatus: status.STATUS_LOADING,
      });

    case auth.LOGIN_SUCCESS:
    case auth.SIGNUP_SUCCESS:
    case auth.TOKEN_RECEIVED:
      return updateState(state, {
        access: {
          token: action.payload.access,
          ...jwtDecode(action.payload.access),
        },
        refresh: {
          token: action.payload.refresh,
          ...jwtDecode(action.payload.refresh),
        },
        error: '',
        authStatus: status.STATUS_SUCCESS,
        tokenStatus: status.STATUS_SUCCESS,
      });

    case auth.LOGIN_FAILED:
    case auth.SIGNUP_FAILED:
      const error = action.payload.response
        ? action.payload.response.message
        : action.payload.message;

      return updateState(state, {
        access: undefined,
        refresh: undefined,
        error: error || 'Authentification error',
        authStatus: status.STATUS_FAILURE,
        tokenStatus: status.STATUS_FAILURE,
      });

    case auth.TOKEN_FAILURE:
      const tokenError = action.payload.response
        ? action.payload.response.message
        : action.payload.message;
      return updateState(state, {
        access: undefined,
        refresh: undefined,
        error: tokenError || 'Authentification error',
        tokenStatus: status.STATUS_FAILURE,
        authStatus: status.STATUS_UPDATE_NEEDED,
      });

    case auth.LOGOUT:
      return updateState(state, {
        access: undefined,
        refresh: undefined,
        tokenStatus: status.STATUS_UPDATE_NEEDED,
        authStatus: status.STATUS_UPDATE_NEEDED,
      });
    default:
      return state;
  }
};

export function accessToken(state) {
  if (state.access) {
    return state.access.token;
  }
}

export function isAccessTokenExpired(state) {
  if (state.access && state.access.exp) {
    return 1000 * state.access.exp - new Date().getTime() < 5000;
  }
  return true;
}

export function refreshToken(state) {
  if (state.refresh) {
    return state.refresh.token;
  }
}

export function isRefreshTokenExpired(state) {
  if (state.refresh && state.refresh.exp) {
    return 1000 * state.refresh.exp - new Date().getTime() < 5000;
  }
  return true;
}

export function isAuthenticated(state) {
  return !isRefreshTokenExpired(state);
}

export function errors(state) {
  return state.errors;
}

export const refreshTime = state => state;

export const authStatus = state => state.authStatus;
export const tokenStatus = state => state.tokenStatus;
export const error = state => state.error;
