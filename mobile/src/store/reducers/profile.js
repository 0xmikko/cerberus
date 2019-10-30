import {updateState} from '../utils/updateState';
import * as actionTypes from '../actions/actionTypes';

const initialState = {
  APNToken: null,
};

export default (state = initialState, action) => {
  switch (action.type) {
    case actionTypes.PROFILE_PREFIX_SAVE_TOKEN:
      return updateState(state, {
        APNToken: action.payload,
      });
    default:
      return state;
  }
};

export const apnToken = state => state.APNToken;
