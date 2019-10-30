import React, {useEffect} from 'react';
import {connect} from 'react-redux';
import AsyncStorage from '@react-native-community/async-storage';
import LoadingView from '../../components/Loading';
import * as actions from '../../store/actions';
import * as reducers from '../../store/reducers';
import * as status from '../../store/utils/status';
import PushNotificationIOS from '@react-native-community/push-notification-ios';
import PushNotification from 'react-native-push-notification';

const AuthSwitcherScreen = ({navigation, refreshToken, tokenStatus}) => {
  // Check next action based on refreshToken

  useEffect(() => {
    AsyncStorage.getItem('refreshToken').then(savedRefreshToken => {
      if (savedRefreshToken) {
        console.log(
          '[AuthLoadingScreen]: Restore refreshToken from AsyncStorage',
        );
        refreshToken(savedRefreshToken);
        return;
      } else {
        navigation.navigate('LoginScreen');
      }
    });
  }, [navigation, refreshToken]);

  // Swith screens depends on status of refresh token auth
  useEffect(() => {
    switch (tokenStatus) {
      case status.STATUS_FAILURE:
        navigation.navigate('LoginScreen');
        return;
      case status.STATUS_SUCCESS:
        navigation.navigate('App');
        return;
    }
  }, [navigation, tokenStatus]);

  return <LoadingView />;
};

const mapStateToProps = state => ({
  tokenStatus: reducers.tokenStatus(state),
});

const mapDispatchToProps = dispatch => ({
  refreshToken: token => dispatch(actions.refreshAccessToken(token)),
});

export default connect(
  mapStateToProps,
  mapDispatchToProps,
)(AuthSwitcherScreen);
