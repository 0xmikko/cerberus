import React from 'react';
import * as statuses from '../store/utils/status';
import LoadingView from './Loading';
import FailureView from './Failure';

export const isDataLoaded = (item) => {

  if (!item) {
    return <LoadingView />;
  }
  const { data, status } = item;
  if (
    (status === statuses.STATUS_UPDATE_NEEDED ||
      status === statuses.STATUS_LOADING) &&
    (!data || data.length === 0)
  ) {
    return <LoadingView />;
  }

  if (status === statuses.STATUS_FAILURE) {
    return <FailureView error="Oops! It's a problem connecting server" />;
  }

  return null;
};
