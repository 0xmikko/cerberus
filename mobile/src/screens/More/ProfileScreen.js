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
import {StyleSheet, Text, Button, Dimensions} from 'react-native';
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
    <TemplateScreen title={'More'}>
      <Text>Profile Here</Text>
      <Button title={'Register device'} onPress={registerDevice} />
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
    <Button title={'Logout'} onPress={navigation.getParam('logout')} />
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
