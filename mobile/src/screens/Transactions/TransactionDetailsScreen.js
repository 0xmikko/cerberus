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
    Alert,
} from 'react-native';
import {Button, Icon, Text} from 'react-native-elements';
import TemplateScreen from '../../components/templateScreen';
import {isDataLoaded} from '../../components/isDataLoaded';
import * as actions from '../../store/actions';
import * as reducers from '../../store/reducers';
import Blockies from "../../components/blockies";

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
    if (status) {
        Alert.alert("You confirmed transaction", "You still could it cancel during confirmation period.",
            [{text: 'Ok', onPress: () => navigation.pop()},])
    } else
    {
        Alert.alert("You cancelled transaction", "Once transaction is cancelled it could not be confirmed back.",
            [{text: 'Ok', onPress: () => navigation.pop()},])
    }

  };

  console.log('TRANS DATA', data);

  return (
    <TemplateScreen title={'Confirmation'}>
        <View style={{
            flex: 1,
            width: '100%',
            alignItems: 'center',
            alignContent: 'center',
            justifyContent: 'space-between'}}
        >
            <View></View>

        <View style={{
            width: '100%',
            alignItems: 'center',
            alignContent: 'center',
            justifyContent: 'center',
        }}>
            <View style={{
                alignItems: 'center',
                alignContent: 'center',
                justifyContent: 'flex-end',
            paddingBottom: 10}}>
                <Text h3>Some, probably you initiated process of sending crypto. Sending transaction was signed by your private key!</Text>
                <Text h3>Press "Confirm" if you send this crypto</Text>
                <Text h3>Press "Cancel" if it's you send this crypto</Text>
            </View>
            <View style={{
                width: '100%',
                alignItems: 'center',
                alignContent: 'center',
                justifyContent: 'center',
            }}>
            <View style={{
                width: '100%',
                alignItems: 'center',
                alignContent: 'center',
                justifyContent: 'center',
                flexDirection: 'row'
            }}>
                <View style={{padding: 15}} >
                <Blockies
                    blockies={data.from} //string content to generate icon
                    size={16} // blocky icon size
                    scale={6}
                    style={{width: 32, height: 32}} // style of the view will wrap the icon
                />
                </View>

                    <Icon name={'keyboard-arrow-right'}  />
                <View style={{padding: 15}} >
                <Blockies
                blockies={data.to} //string content to generate icon
                size={16} // blocky icon size
                scale={6}
                style={{width: 32, height: 32}} // style of the view will wrap the icon
            />
                </View>
            </View>
            </View>
        <View style={{
            width: '100%',
            alignItems: 'center',
            alignContent: 'center',
            justifyContent: 'center',
            paddingTop: 15,
        }}>
            <Text h3>From: {data.from}</Text>
            <Text h3>To: {data.to}</Text>
            <Text h2>Amount: {data.amount}</Text>
        </View>
        </View>
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
          buttonStyle={{backgroundColor: 'green', borderRadius: 7,}}
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
          buttonStyle={{backgroundColor: 'red', borderRadius: 7}}
          onPress={() => onConfirm(false)}
        />
      </View>
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
