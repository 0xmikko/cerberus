import React, {useEffect} from 'react';
import {FlatList, SectionList, StyleSheet, View} from 'react-native';
import {Text} from 'react-native-elements';
import {toHumanDate} from '../../utils/humanDate';
import AccountsItem from './AccountsItem';

const AccountsList = ({data, onSelect}) => {
  if (!data) {
    return <View />;
  }

  const renderItem = ({item}) => {
    console.log(item);
    return <AccountsItem key={item.id} item={item} onSelect={onSelect} />;
  };

  return (
    <FlatList
      styles={styles.container}
      data={data}
      renderItem={renderItem}
      // renderSectionHeader={({section: {date}}) => (
      //   <View style={styles.header}>
      //     <Text>{toHumanDate(date)}</Text>
      //   </View>
      // )}
      // onRefresh={() => console.log("refresh")}
      keyExtractor={item => item.id}
    />
  );
};

const styles = StyleSheet.create({
  header: {
    paddingLeft: 15,
    paddingTop: 20,
    paddingBottom: 4,
  },
});

export default AccountsList;
