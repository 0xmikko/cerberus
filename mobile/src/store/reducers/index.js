import {combineReducers} from 'redux';
import * as actionTypes from '../actions/actionTypes';
import {createDataLoader} from './dataloader';
import auth, * as fromAuth from './auth';
import profileToken, * as fromProfileToken from './profile';

const rootReducer = combineReducers({
  auth,
  transactions: createDataLoader(actionTypes.TRANSACTIONS_PREFIX),
  accounts: createDataLoader(actionTypes.ACCOUNTS_PREFIX),
  user: createDataLoader(actionTypes.PROFILE_PREFIX),
  profileToken,
});

// Authentication
export const isAuthenticated = state => fromAuth.isAuthenticated(state.auth);
export const accessToken = state => fromAuth.accessToken(state.auth);
export const isAccessTokenExpired = state =>
  fromAuth.isAccessTokenExpired(state.auth);
export const refreshToken = state => fromAuth.refreshToken(state.auth);
export const isRefreshTokenExpired = state =>
  fromAuth.isRefreshTokenExpired(state.auth);
export const authError = state => fromAuth.error(state.auth);
export const refreshTime = state => fromAuth.refreshTime(state.auth);
export const authStatus = state => fromAuth.authStatus(state.auth);
export const tokenStatus = state => fromAuth.tokenStatus(state.auth);
export function withAuth(headers = {}) {
  return state => ({
    ...headers,
    Authorization: `Bearer ${accessToken(state)}`,
  });
}

// APN Token
export const apnToken = state => fromProfileToken.apnToken(state.profileToken);

// Transactions
export const transactionsList = state => state.transactions.List;
export const transactionDetails = state => state.transactions.Details;
export const transactionUpdates = state => state.transactions.Updates;

// Account
export const accountsList = state => state.accounts.List;
export const accountDetails = state => state.accounts.Details;

// User Profile
export const profile = state => state.user.Details;
export const profileUpdates = state => state.user.Updates;

export default rootReducer;
