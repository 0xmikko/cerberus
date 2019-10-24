import React, {Fragment} from 'react';
import {StyleSheet, SafeAreaView, StatusBar, View} from 'react-native';
import {Icon, Text} from 'react-native-elements';

const FailureView = ({error}) => (
  <Fragment>
    <StatusBar barStyle="default" />
    <SafeAreaView style={styles.container}>
      <View style={styles.containerText}>
        <Icon name="ios-bug" type="ionicon" size={50} color="blue" />
        <Text h4 style={{textAlign: 'center'}}>
          {error}
        </Text>
      </View>
    </SafeAreaView>
  </Fragment>
);

const styles = StyleSheet.create({
  container: {
    flex: 1,
    alignItems: 'center',
    alignContent: 'center',
    justifyContent: 'center',
  },
  containerText: {
    alignItems: 'center',
    alignContent: 'center',
    padding: 20,
  },
});

export default FailureView;
