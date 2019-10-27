/**
 *
 * Cerberus Project for ChainLink Hackathon
 * https://github.com/MikaelLazarev/cerberus
 *
 * @format
 * @flow
 */

/* eslint-disable react-native/no-inline-styles */
import React, {useState, useEffect} from 'react';
import {connect} from 'react-redux';
import {
  StyleSheet,
  ScrollView,
  View,
  Dimensions,
  SafeAreaView,
} from 'react-native';
import {Text} from 'react-native-elements';
import TransactionList from '../../containers/Transactions/TransactionList';
import {isDataLoaded} from '../../components/isDataLoaded';
import * as actions from '../../store/actions';
import * as reducers from '../../store/reducers';

const TransactionsListScreen = ({
  transactions,
  getTransactionsList,
  navigation,
}) => {
  useEffect(() => {
    getTransactionsList();
    return () => {};
  }, [getTransactionsList]);

  let notReadyStatus = isDataLoaded(transactions);
  if (notReadyStatus) {
    return notReadyStatus;
  }
  const {data} = transactions;

  if (!data) {
    return <View />;
  }

  const onSelect = id => {
    console.log(id);
    navigation.push('TransactionDetailsScreen', {id});
  };

  return (
    <SafeAreaView>
      <ScrollView style={styles.scrollContainer}>
        <View
          style={{
            flex: 1,
            borderRadius: 10,
            backgroundColor: '#F6F7F8',
            height: Dimensions.get('window').height,
          }}>
          <Text h1>Transactions</Text>
          <TransactionList data={data} onSelect={onSelect} />
        </View>
      </ScrollView>
    </SafeAreaView>
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
TransactionsListScreen.navigationOptions = ({navigation}) => ({
  header: null,
});

const mapStateToProps = state => ({
  transactions: reducers.transactionsList(state),
});

const mapDispatchToProps = dispatch => ({
  getTransactionsList: () => dispatch(actions.getTransactionsList()),
});

export default connect(
  mapStateToProps,
  mapDispatchToProps,
)(TransactionsListScreen);
