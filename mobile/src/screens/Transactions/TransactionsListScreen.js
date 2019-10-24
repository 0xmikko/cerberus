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
import {StyleSheet, Text, ScrollView, View, Dimensions} from 'react-native';
import * as actions from '../../store/actions';
import * as reducers from '../../store/reducers';

const TransactionsListScreen = ({navigation}) => {
  return (
    <ScrollView style={styles.scrollContainer}>
      <View
        style={{
          flex: 1,
          borderRadius: 10,
          backgroundColor: '#F6F7F8',
          height: Dimensions.get('window').height,
          marginTop: -10,
        }}>
        <Text>Accounts</Text>
      </View>
    </ScrollView>
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
  accounts: reducers.accountsList(state),
});

const mapDispatchToProps = dispatch => ({
  getAccountsList: () => dispatch(actions.getAccountsList()),
});

export default connect(
  mapStateToProps,
  mapDispatchToProps,
)(TransactionsListScreen);
