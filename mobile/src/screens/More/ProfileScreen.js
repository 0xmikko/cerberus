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
import {StyleSheet, View, Text,  Dimensions} from 'react-native';
import { Button } from 'react-native-elements'
import * as actions from '../../store/actions';
import * as reducers from '../../store/reducers';
import TemplateScreen from '../../components/templateScreen';

const ProfileScreen = ({
  navigation,
  logout,
  refreshToken,
  apnToken,
  registerToken,
}) => {
  useEffect(() => {
    navigation.setParams({logout});
  }, [logout]);

  useEffect(() => {
    if (!refreshToken) {
      navigation.navigate('AuthStack');
    }
  }, [refreshToken]);

  const registerDevice = () => {
    if (apnToken) {
      registerToken(apnToken.token);
    }
  };

  return (
    <TemplateScreen title={'Settings'}>
      <View style={{
        width: '100%',
        alignItems: 'center',
        justifyContent: 'center',
        alignContent: 'center',}}>
      <View style={{
        width: '90%',
        paddingTop: 40,

      }}>
          <Button title={'Register device'} onPress={registerDevice} containerStyle={{width: '100%', paddingBottom: 20}}/>
        <Button title={'Panic mode'} onPress={registerDevice} containerStyle={{width: '100%'}} buttonStyle={{backgroundColor: 'red'}}/>
      </View>
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

ProfileScreen.navigationOptions = ({navigation}) => ({
  headerRight: () => (
    <Button type="clear" title={'Logout'} onPress={navigation.getParam('logout')} />
  ),
});

const mapStateToProps = state => ({
  refreshToken: reducers.refreshToken(state),
  apnToken: reducers.apnToken(state),
});

const mapDispatchToProps = dispatch => ({
  logout: () => dispatch(actions.logout()),
  registerToken: token => dispatch(actions.registerToken(token)),
});

export default connect(
  mapStateToProps,
  mapDispatchToProps,
)(ProfileScreen);
