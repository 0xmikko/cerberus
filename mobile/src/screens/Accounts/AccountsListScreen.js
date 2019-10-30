/**
 *
 * Cerberus Project for ChainLink Hackathon
 * https://github.com/MikaelLazarev/cerberus
 *
 * @format
 * @flow
 */

/* eslint-disable react-native/no-inline-styles */
import React, {useEffect} from 'react';
import {connect} from 'react-redux';
import {StyleSheet, Text, Button, Dimensions, View} from 'react-native';
import * as actions from '../../store/actions';
import * as reducers from '../../store/reducers';
import TemplateScreen from '../../components/templateScreen';
import {isDataLoaded} from '../../components/isDataLoaded';
import AccountsList from '../../containers/Accounts/AccountsList';

const AccountsListScreen = ({accounts, getAccountsList, navigation}) => {
  useEffect(() => {
    getAccountsList();
    return () => {};
  }, [getAccountsList]);

  let notReadyStatus = isDataLoaded(accounts);
  if (notReadyStatus) {
    return notReadyStatus;
  }
  console.log(accounts);
  const {data} = accounts;

  if (!data) {
    return <View />;
  }

  return (
    <TemplateScreen
      title={'Accounts'}
      rightButton={true}
      onPress={() => console.log('Oressed')}>
      <View
        style={{
          flex: 1,
          borderRadius: 10,
          backgroundColor: '#F6F7F8',
          height: Dimensions.get('window').height,
        }}>
        <AccountsList data={data} />
      </View>
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
