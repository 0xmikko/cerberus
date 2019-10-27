/**
 *
 * Cerberus Project for ChainLink Hackathon
 * https://github.com/MikaelLazarev/cerberus
 *
 * @format
 * @flow
 */

import React, {useState, useEffect} from 'react';
import {connect} from 'react-redux';
import {
  StyleSheet,
  View,
  SafeAreaView,
  Image,
  ActivityIndicator,
  TouchableOpacity,
} from 'react-native';
import {Text, Button, Input} from 'react-native-elements';
import * as yup from 'yup';
import * as actions from '../../store/actions';
import * as reducers from '../../store/reducers';
import * as status from '../../store/utils/status';

const LoginScreen = ({
  authStatus,
  authError,
  isAuthenticated,
  navigation,
  login,
}) => {
  // State
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [valid, setValid] = useState(false);
  const [isSubmitted, setIsSubmitted] = useState(false);
  const [authErrorDisplay, setAuthErrorDisplay] = useState('');

  // Validation Hook
  useEffect(() => {
    const scheme = yup.object().shape({
      email: yup
        .string()
        .email()
        .required(),
      password: yup.string().required(),
    });
    scheme.isValid({email, password}).then(valid => setValid(valid));
  }, [email, password]);

  // Show Loading Modal during loading phase
  // and update Error message
  useEffect(() => {
    setIsSubmitted(authStatus === status.STATUS_LOADING);
    if (authStatus === status.STATUS_FAILURE) {
      setAuthErrorDisplay(authError);
    }
    if (isAuthenticated) {
      navigation.navigate('App');
    }
  }, [authError, authStatus, isAuthenticated, navigation]);

  return (
    <SafeAreaView style={styles.container}>
      <View style={styles.logo}>
        <Image source={require('../../../logo.png')} style={styles.image} />
        <Text h2>Cerberus Wallet</Text>
      </View>

      <View style={styles.input}>
        <Input placeholder="email" onChangeText={setEmail} value={email} />
      </View>
      <View style={styles.input}>
        <Input
          placeholder="password"
          secureTextEntry={true}
          onChangeText={setPassword}
          value={password}
        />
      </View>
      <Text>{authErrorDisplay}</Text>
      <View style={styles.button}>
        <Button
          title="Login"
          onPress={() => login(email, password)}
          disabled={!valid}
          loading={isSubmitted}
        />
      </View>

      <View>
        <Button
          title="Sign Up"
          onPress={() => navigation.navigate('SignUpScreen')}
          type={'clear'}
        />
      </View>
    </SafeAreaView>
  );
};

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: 'white',
    alignContent: 'center',
    alignItems: 'center',
    justifyContent: 'center',
  },
  logo: {
    alignItems: 'center',
    alignContent: 'center',
    justifyContent: 'center',
    paddingBottom: 30,
  },
  image: {
    width: 100,
    height: 100,
    resizeMode: 'contain',
  },
  input: {
    paddingTop: 5,
    width: '80%',
  },

  button: {
    width: '75%',
    paddingTop: 20,
    paddingBottom: 140,
  },
});

const mapStateToProps = state => ({
  authStatus: reducers.authStatus(state),
  authError: reducers.authError(state),
  isAuthenticated: reducers.isAuthenticated(state),
});

const mapDispatchToProps = dispatch => ({
  login: (email, password) => dispatch(actions.login(email, password)),
});

export default connect(
  mapStateToProps,
  mapDispatchToProps,
)(LoginScreen);
