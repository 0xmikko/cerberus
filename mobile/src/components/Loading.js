import React, {Fragment} from 'react';
import {
  StyleSheet,
  SafeAreaView,
  ActivityIndicator,
  StatusBar,
} from 'react-native';

const loadingView = () => (
  <Fragment>
    <StatusBar barStyle="default" />
    <SafeAreaView style={styles.container}>
      <ActivityIndicator />
    </SafeAreaView>
  </Fragment>
);

const styles = StyleSheet.create({
  container: {
    flex: 1,
    alignItems: 'center',
    alignContent: 'center',
    justifyContent: 'space-around',
  },
});

export default loadingView;
