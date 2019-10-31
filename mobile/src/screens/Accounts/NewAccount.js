import React, {createRef} from 'react';
import {connect} from 'react-redux';
import * as actions from '../../store/actions';
import * as reducers from '../../store/reducers';
import {Alert, Text, StyleSheet, View} from 'react-native';
import {Button} from 'react-native-elements';
import QRCodeScanner from 'react-native-qrcode-scanner';

const NewAccount = ({navigation, createAccount}) => {
  const ref = createRef();

  const getAddress = address => {
    console.log(address);
    createAccount({account: address});
    navigation.pop();
  };

  const onSuccess = e => {
    // alert(e ? e.data : "NO E");
    try {
      if (e !== undefined && e.data !== undefined) {
        Alert.alert('Please check the correct addres', e.data, [
          {text: 'Retake QR', onPress: () => ref.current.reactivate()},
          {text: 'Use this', onPress: () => getAddress(e.data)},
        ]);
      }

      //
    } catch (e) {
      console.log(e);
      alert('Error' + e);
    }
  };

  // onSuccess({ data: JSON.stringify({ recipient_id: "fff", amount: 3434.12 })})
  // At mount
  return (
    <View style={{flex: 1, height: '100%'}}>
      <QRCodeScanner
        ref={ref}
        onRead={onSuccess}
        // flashMode={QRCodeScanner.Constants.FlashMode.torch}
        topContent={
          <View>
            <Text style={styles.centerText}>
              Please scan your Desktop Cerberus Wallet
            </Text>
          </View>
        }
        bottomContent={<Button title={'OK. Got it!'} />}
      />
    </View>
  );
};

const styles = StyleSheet.create({
  centerText: {
    flex: 1,
    fontSize: 18,
    padding: 32,
    color: '#777',
  },
  textBold: {
    fontWeight: '500',
    color: '#000',
  },
  buttonText: {
    fontSize: 21,
    color: 'rgb(0,122,255)',
  },
  buttonTouchable: {
    padding: 16,
  },
});

const mapDispatchToProps = dispatch => ({
  createAccount: account => dispatch(actions.createAccount('new', account)),
});

export default connect(
  null,
  mapDispatchToProps,
)(NewAccount);
