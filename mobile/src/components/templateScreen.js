/**
 * Template Screen
 * DBS Lifestype App
 * https://github.com/MikaelLazarev/DBSLifestyle
 *
 * @format
 * @flow
 */

import React from 'react';
import {StyleSheet, View, SafeAreaView, Image} from 'react-native';
import {Text, Button} from 'react-native-elements';

const templateScreen = ({title, children, rightButton, onPress}) => {
  let RightButton;
  if (rightButton) {
    RightButton = (
      <Button
        icon={{
          name: 'add',
          color: 'blue',
        }}
        type={'clear'}
        onPress={onPress}
      />
    );
  }

  return (
    <SafeAreaView style={styles.container}>
      <View
        style={{
          paddingTop: 40,
          paddingLeft: 20,
          paddingBottom: 15,
          paddingRight: 10,
          flexDirection: 'row',
          justifyContent: 'space-between',
          alignContent: 'space-between',
          alignItems: 'center',
          width: '100%',
        }}>
        <Text h1>{title}</Text>
        {RightButton}
      </View>
      {children}
    </SafeAreaView>
  );
};

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: '#F6F7F8',
    alignContent: 'center',
    alignItems: 'flex-start',
    justifyContent: 'flex-start',
  },
});

export default templateScreen;
