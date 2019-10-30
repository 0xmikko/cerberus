import React, {useEffect} from 'react';
import {View} from 'react-native';
import {connect} from 'react-redux';
import PushNotification from 'react-native-push-notification';
import PushNotificationIOS from '@react-native-community/push-notification-ios';
import * as reducers from './src/store/reducers';
import * as actions from './src/store/actions';

export const Notifications = ({saveAPNToken}) => {
  useEffect(() => {
    PushNotification.configure({
      // (optional) Called when Token is generated (iOS and Android)
      onRegister: function(token) {
        console.log('TOKEN:', token);
        saveAPNToken(token);
      },

      onError: function(e) {
        console.log('ERROR', e);
      },

      // (required) Called when a remote or local notification is opened or received
      onNotification: function(notification) {
        console.log('NOTIFICATION:', notification);

        // process the notification

        // required on iOS only (see fetchCompletionHandler docs: https://github.com/react-native-community/react-native-push-notification-ios)
        notification.finish(PushNotificationIOS.FetchResult.NoData);
      },

      // ANDROID ONLY: GCM or FCM Sender ID (product_number) (optional - not required for local notifications, but is need to receive remote push notifications)
      senderID: 'YOUR GCM (OR FCM) SENDER ID',

      // IOS ONLY (optional): default: all - Permissions to register.
      permissions: {
        alert: true,
        badge: true,
        sound: true,
      },

      // Should the initial notification be popped automatically
      // default: true
      popInitialNotification: true,

      /**
       * (optional) default: true
       * - Specified if permissions (ios) and token (android and ios) will requested or not,
       * - if not, you must call PushNotificationsHandler.requestPermissions() later
       */
      requestPermissions: true,
    });
  }, [saveAPNToken]);

  return <View />;
};

const mapStateToProps = state => ({});

const mapDispatchToProps = dispatch => ({
  saveAPNToken: token => dispatch(actions.saveAPNToken(token)),
});

export default connect(
  mapStateToProps,
  mapDispatchToProps,
)(Notifications);
