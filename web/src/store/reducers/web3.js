import {updateState} from '../utils/updateState';
import * as actionTypes from '../actions/actionTypes';
import * as status from '../utils/status';

const initialState = {
  web3: null,
  accounts: null,
  gasPrice: null,
  contractDeployStatus: status.STATUS_UPDATE_NEEDED,
  contractAddress: null,
  paymentStatus: status.STATUS_SUCCESS,
};

export default (state = initialState, action) => {
  switch (action.type) {
    case actionTypes.WEB3_REQUEST:
      return updateState(state, {
        web3: null,
      });

    case actionTypes.WEB3_SUCCESS:
      console.log(action)
        return updateState(state, {
          ...action.payload,
        });

    case actionTypes.CONTRACT_DEPLOY_REQUEST:
      console.log(action)
      return updateState(state, {
        contractDeployStatus: status.STATUS_LOADING,
      })


    case actionTypes.CONTRACT_DEPLOY_SUCCESS:
      console.log(action)
      return updateState(state, {
        contractAddress: action.payload,
        contractDeployStatus: status.STATUS_SUCCESS,
      })

    case actionTypes.PAYMENT_REQUEST:
      console.log(action)
      return updateState(state, {
        paymentStatus: status.STATUS_LOADING,
      })


    case actionTypes.PAYMENT_SUCCESS:
      console.log(action)
      return updateState(state, {
        paymentStatus: status.STATUS_SUCCESS,
      })

    default:
      return state;
  }


};

export const web3= state => state.web3;
export const accounts = state => state.accounts;
export const contractDeployStatus = state => state.contractDeployStatus;
export const contractAddress = state => state.contractAddress;
export const paymentStatus = state => state.paymentStatus;
