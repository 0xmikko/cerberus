import {combineReducers} from 'redux';
import { connectRouter } from 'connected-react-router'
import * as actionTypes from '../actions/actionTypes';
import web3, * as fromWeb3 from './web3';


export default history => combineReducers({
  router:  connectRouter(history),
  web3,

});

export const Web3 = state => fromWeb3.web3(state.web3)
export const accounts = state => fromWeb3.accounts(state.web3)
export const contractDeployStatus = state => fromWeb3.contractDeployStatus(state.web3)
export const contractAddress = state => fromWeb3.contractAddress(state.web3)

