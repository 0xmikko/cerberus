/**
 * Sample React Native App
 * https://github.com/facebook/react-native
 *
 * @format
 * @flow
 */

import React from 'react';
import {Provider} from 'react-redux';
import {createAppContainer, createSwitchNavigator} from 'react-navigation';
import {createStackNavigator} from 'react-navigation-stack';
import {createBottomTabNavigator} from 'react-navigation-tabs';
import {Icon} from 'react-native-elements';
import {ThemeProvider} from 'react-native-elements';

import Notifications from './notifications';
import configureStore from './src/store/configureStore';

// AuthStack
import AuthSwitcherScreen from './src/screens/Auth/AuthSwitcherScreen';
import LoginScreen from './src/screens/Auth/LoginScreen';
import SignUpScreen from './src/screens/Auth/SignUpScreen';
import ErrorScreen from './src/screens/Auth/ErrorScreen';

// Transactions Stack
import TransactionsListScreen from './src/screens/Transactions/TransactionsListScreen';
import TransactionDetailsScreen from './src/screens/Transactions/TransactionDetailsScreen';

// Accounts stack
import AccountsListScreen from './src/screens/Accounts/AccountsListScreen';

// More
import ProfileScreen from './src/screens/More/ProfileScreen';
import {theme} from './styles';
import PushNotification from 'react-native-push-notification';

const store = configureStore();

const TransactionsStack = createStackNavigator(
  {
    TransactionsListScreen,
    TransactionDetailsScreen,
  },
  {
    initialRouteName: 'TransactionsListScreen',
  },
);

const AccountsStack = createStackNavigator(
  {
    AccountsListScreen,
  },
  {
    initialRouteName: 'AccountsListScreen',
  },
);

const MoreStack = createStackNavigator(
  {
    ProfileScreen: {
      screen: ProfileScreen,
      title: 'Profile',
    },
  },
  {
    initialRouteName: 'ProfileScreen',
  },
);

const AppStack = createBottomTabNavigator(
  {
    TransactionsStack: {
      screen: TransactionsStack,
      navigationOptions: {
        tabBarLabel: 'Transactions',
        tabBarIcon: (
          <Icon
            name="format-list-bulleted"
            type="material-community"
            color="grey"
          />
        ), // Todo: add focused color
      },
    },

    AccountsStack: {
      screen: AccountsStack,
      navigationOptions: {
        tabBarLabel: 'Account',
        tabBarIcon: (
          <Icon name="credit-card" type="material-community" color="grey" />
        ), // Todo: add focused color
      },
    },
    MoreStack: {
      screen: MoreStack,
      navigationOptions: {
        tabBarLabel: 'More',
        tabBarIcon: (
          <Icon name="dots-horizontal" type="material-community" color="grey" />
        ), // Todo: add focused color
      },
    },
  },
  {
    initialRouteName: 'TransactionsStack',
    tabBarOptions: {
      activeTintColor: 'grey',
      labelStyle: {
        fontSize: 12,
      },
      style: {
        backgroundColor: 'white',
      },
    },
  },
);

const AuthStack = createSwitchNavigator(
  {
    AuthSwitcherScreen,
    LoginScreen,
    SignUpScreen,
    ErrorScreen,
  },
  {
    initialRouteName: 'AuthSwitcherScreen',
  },
);

const AppContainer = createAppContainer(
  createSwitchNavigator(
    {
      AuthStack,
      App: AppStack,
    },
    {
      initialRouteName: 'AuthStack',
    },
  ),
);

class App extends React.Component {

  // UNSAFE_componentWillMount() {
  //   PushNotificationIOS.addEventListener('register', this._onRegistered);
  //   PushNotificationIOS.addEventListener(
  //     'registrationError',
  //     this._onRegistrationError,
  //   );
  //   PushNotificationIOS.addEventListener(
  //     'notification',
  //     this._onRemoteNotification,
  //   );
  //   PushNotificationIOS.addEventListener(
  //     'localNotification',
  //     this._onLocalNotification,
  //   );
  //
  //   PushNotificationIOS.requestPermissions();
  // }
  //
  // _sendLocalNotification() {
  //   PushNotificationIOS.presentLocalNotification({
  //     alertBody: 'Sample local notification',
  //     applicationIconBadgeNumber: 1,
  //   });
  // }
  //
  // _onRegistered(deviceToken) {
  //   Alert.alert('Registered For Remote Push', `Device Token: ${deviceToken}`, [
  //     {
  //       text: 'Dismiss',
  //       onPress: null,
  //     },
  //   ]);
  // }
  //
  // _onRegistrationError(error) {
  //   Alert.alert(
  //     'Failed To Register For Remote Push',
  //     `Error (${error.code}): ${error.message}`,
  //     [
  //       {
  //         text: 'Dismiss',
  //         onPress: null,
  //       },
  //     ],
  //   );
  // }
  //
  // _onRemoteNotification(notification) {
  //   const result = `Message: ${notification.getMessage()};\n
  //     badge: ${notification.getBadgeCount()};\n
  //     sound: ${notification.getSound()};\n
  //     category: ${notification.getCategory()};\n
  //     content-available: ${notification.getContentAvailable()}.`;
  //
  //   Alert.alert('Push Notification Received', result, [
  //     {
  //       text: 'Dismiss',
  //       onPress: null,
  //     },
  //   ]);
  // }
  //
  // _onLocalNotification(notification) {
  //   Alert.alert(
  //     'Local Notification Received',
  //     'Alert message: ' + notification.getMessage(),
  //     [
  //       {
  //         text: 'Dismiss',
  //         onPress: null,
  //       },
  //     ],
  //   );
  // }
  //
  // _showPermissions() {
  //   PushNotificationIOS.checkPermissions(permissions => {
  //     this.setState({permissions});
  //   });
  // }

  render() {
    return (
      <Provider store={store}>
        <Notifications />
        <ThemeProvider theme={theme}>
          <AppContainer />
        </ThemeProvider>
      </Provider>
    );
  }
}

export default App;
