/**
 *
 * Cerberus Project for ChainLink Hackathon
 * https://github.com/MikaelLazarev/cerberus
 *
 * @format
 * @flow
 */

/* eslint-disable react-native/no-inline-styles */
import React, {useState, useEffect, useMemo} from 'react';
import {connect} from 'react-redux';
import {
  StyleSheet,
  ScrollView,
  View,
  Dimensions,
  SafeAreaView,
} from 'react-native';
import {Button} from 'react-native-elements';
import {Text} from 'react-native-elements';
import TemplateScreen from '../../components/templateScreen';
import TransactionList from '../../containers/Transactions/TransactionList';
import {isDataLoaded} from '../../components/isDataLoaded';
import * as actions from '../../store/actions';
import * as reducers from '../../store/reducers';

const TransactionDetailsScreen = ({
  transactionDetails,
  getTransactionDetails,
  confirmTransaction,
  navigation,
}) => {
  const id = useMemo(() => navigation.getParam('id'), [navigation]);
  console.log('OPOPOPOP', id, transactionDetails);

  const [hash, setHash] = useState(0);
  useEffect(() => {
    getTransactionDetails(id);
  }, [getTransactionDetails, id]);

  const notReadyStatus = isDataLoaded(transactionDetails[id]);
  if (notReadyStatus) {
    return notReadyStatus;
  }

  const {data} = transactionDetails[id];

  const onConfirm = status => {
    const hash = '1500';
    confirmTransaction(id, status, hash);
  };

  console.log('TRANS DATA', data);

  return (
    <TemplateScreen title={'Confirmation'}>
      <Text>To: {data.to}</Text>
      <View
        style={{
          width: '100%',
          alignItems: 'center',
          alignContent: 'center',
          justifyContent: 'center',
        }}>
        <Button
          title={'Confirm'}
          containerStyle={{
            width: '100%',
            paddingBottom: 20,
            paddingLeft: 20,
            paddingRight: 20,
          }}
          buttonStyle={{backgroundColor: 'green'}}
          onPress={() => onConfirm(true)}
        />
        <Button
          title={'Cancel'}
          containerStyle={{
            width: '100%',
            paddingBottom: 40,
            paddingLeft: 20,
            paddingRight: 20,
          }}
          buttonStyle={{backgroundColor: 'red'}}
          onPress={() => onConfirm(false)}
        />
      </View>
    </TemplateScreen>
  );
};

const styles = StyleSheet.create({
  container: {
    flex: 1,
    width: '100%',

    alignContent: 'flex-start',
    alignItems: 'flex-start',
    justifyContent: 'space-between',
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

// Navigation options
TransactionDetailsScreen.navigationOptions = ({navigation}) => ({
  header: null,
});

const mapStateToProps = state => ({
  transactionDetails: reducers.transactionDetails(state),
});

const mapDispatchToProps = dispatch => ({
  getTransactionDetails: id => dispatch(actions.getTransactionDetails(id)),
  confirmTransaction: (id, status, hash) =>
    dispatch(actions.confirmTransaction(id, status, hash)),
});

export default connect(
  mapStateToProps,
  mapDispatchToProps,
)(TransactionDetailsScreen);
