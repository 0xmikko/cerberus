import * as actionTypes from './actionTypes';
import {
  createDataloaderDetailAction,
  createDataloaderCreateUpdateDataAction,
} from './dataloader';

// Get User Prorile
export const getProfile = createDataloaderDetailAction(
  '/api/user/',
  actionTypes.PROFILE_PREFIX,
);
export const updateProfile = createDataloaderCreateUpdateDataAction(
  '/api/user/',
  actionTypes.PROFILE_PREFIX,
);
