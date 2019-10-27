/**
 *
 * Cerberus Project for ChainLink Hackathon
 * https://github.com/MikaelLazarev/cerberus
 *
 * @format
 * @flow
 */

/* eslint-disable react-native/no-inline-styles */
import React from 'react';
import {connect} from 'react-redux';
import {
  StyleSheet,
  Text,
  ScrollView,
  View,
  Dimensions,
  SafeAreaView,
} from 'react-native';
import * as actions from '../../store/actions';
import * as reducers from '../../store/reducers';
import TemplateScreen from '../../components/templateScreen';

const AccountsListScreen = ({navigation}) => {
  return (
    <TemplateScreen title={'Accounts'}>
      <Text>Accounts here</Text>
    </TemplateScreen>
  );
};

const styles = StyleSheet.create({
  container: {
    flex: 1,

    alignContent: 'flex-start',
    alignItems: 'flex-start',
    justifyContent: 'flex-start',
  },
  containerBalance: {
    flex: 1,
    backgroundColor: '#F6F7F8',
    alignContent: 'center',
    alignItems: 'center',
    justifyContent: 'center',
    height: Math.round(Dimensions.get('window').height * 0.2),
  },
  scrollContainer: {
    width: '100%',
  },
});

// Navifation options
AccountsListScreen.navigationOptions = ({navigation}) => ({
  header: null,
});

const mapStateToProps = state => ({
  accounts: reducers.accountsList(state),
});

const mapDispatchToProps = dispatch => ({
  getAccountsList: () => dispatch(actions.getAccountsList()),
});

export default connect(
  mapStateToProps,
  mapDispatchToProps,
)(AccountsListScreen);
