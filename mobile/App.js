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

import configureStore from './src/store/configureStore';

// AuthStack
import AuthSwitcherScreen from './src/screens/Auth/AuthSwitcherScreen';
import LoginScreen from './src/screens/Auth/LoginScreen';
import SignUpScreen from './src/screens/Auth/SignUpScreen';
import ErrorScreen from './src/screens/Auth/ErrorScreen';

// Transactions Stack
import TransactionsListScreen from './src/screens/Transactions/TransactionsListScreen';

// Accounts stack
import AccountsListScreen from './src/screens/Accounts/AccountsListScreen';

// More
import ProfileScreen from './src/screens/More/ProfileScreen';
import {theme} from './styles';

const store = configureStore();

const TransactionsStack = createStackNavigator(
  {
    TransactionsListScreen,
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
    ProfileScreen: {
      screen: ProfileScreen,
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

export default class App extends React.Component {
  render() {
    return (
      <Provider store={store}>
        <ThemeProvider theme={theme}>
          <AppContainer />
        </ThemeProvider>
      </Provider>
    );
  }
}
