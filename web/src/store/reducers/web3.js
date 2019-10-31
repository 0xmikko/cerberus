import {updateState} from '../utils/updateState';
import * as actionTypes from '../actions/actionTypes';

const initialState = {
  web3: null,
  accounts: null,
  gasPrice: null,
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
    default:
      return state;
  }
};

export const web3= state => state.web3;
export const accounts = state => state.accounts;
