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
import {StyleSheet, Linking, View, Text, Button} from 'react-native';
import FailureView from '../../components/Failure';

class WelcomeScreen extends React.Component {
  render() {
    return (
      <FailureView error="Connection error. The internet connection appears to be offlie." />
    );
  }
}

export default WelcomeScreen;
