import {RSAA} from 'redux-api-middleware';
import {getApiById} from '../utils/api';
import {withAuth} from '../reducers';
import * as actionTypes from './actionTypes';

export const createDataloaderListAction = (api, actionPrefix) => {
  return id => ({
    [RSAA]: {
      endpoint: getApiById(api, id),
      method: 'GET',
      headers: withAuth({'Content-Type': 'application/json'}),
      types: [
        actionPrefix + actionTypes.LIST_REQUEST,
        actionPrefix + actionTypes.LIST_SUCCESS,
        actionPrefix + actionTypes.LIST_FAILURE,
      ],
    },
  });
};

export const createDataloaderDetailAction = (api, actionPrefix) => {
  return (id, hash) => ({
    [RSAA]: {
      endpoint: getApiById(api, id),
      method: 'GET',
      headers: withAuth({'Content-Type': 'application/json'}),
      types: [
        {type: actionPrefix + actionTypes.DETAIL_REQUEST, meta: {id, hash}},
        {type: actionPrefix + actionTypes.DETAIL_SUCCESS, meta: {id, hash}},
        {type: actionPrefix + actionTypes.DETAIL_FAILURE, meta: {id, hash}},
      ],
    },
  });
};

export const createDataloaderCreateUpdateDataAction = (api, actionPrefix) => {
  return (id, data, hash) => {
    console.log('[ACTIONS]: Update Data Loader Detail', id, data, hash);

    if (api === undefined || id === undefined) {
      throw 'Error in updateDataLoaderDetail, wrong parameters!\napi:' +
        api +
        '\nid: ' +
        id;
    }

    let headers = {'Content-Type': 'application/x-www-form-urlencoded'};
    if (hash === undefined) {
      hash = 0;
    }

    // If data is not formData we change Content-Type and JSONify our data
    if (!(data instanceof FormData)) {
      headers = {'Content-Type': 'application/json'};
      data = JSON.stringify(data);
    }

    const method = id.toString().startsWith('new') ? 'POST' : 'PUT';
    api = id.toString().startsWith('new')
      ? getApiById(api)
      : getApiById(api, id);

    console.log('DATA SENT:', data);

    return {
      [RSAA]: {
        endpoint: api,
        method: method,
        headers: withAuth(headers),
        body: data,
        types: [
          {
            type: actionPrefix + actionTypes.UPLOAD_REQUEST,
            meta: {id, hash},
          },
          {
            type: actionPrefix + actionTypes.UPLOAD_SUCCESS,
            meta: {id, hash},
          },
          {
            type: actionPrefix + actionTypes.UPLOAD_FAILURE,
            meta: {id, hash},
          },
        ],
      },
    };
  };
};
