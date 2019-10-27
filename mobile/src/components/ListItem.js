import React from 'react';
import {View, TouchableOpacity, StyleSheet} from 'react-native';
import {Text, Avatar} from 'react-native-elements';
import {styles} from './cardstyles';
import {numberWithCommas} from '../utils/numberFormat';
import Blockies from './blockies';

const ListItem = ({icon, title, subtitle, amount, onSelect}) => {
  console.log('CO', amount, typeof amount);

  let amountTag;
  if (amount !== '-' && typeof amount === 'number') {
    amountTag = <Text style={{fontSize: 14}}>${numberWithCommas(amount)}</Text>;
  }

  return (
    <TouchableOpacity onPress={onSelect}>
      <View style={styles.container}>
        <View style={styles.leftContainer}>
          <Blockies
            blockies={title} //string content to generate icon
            size={16} // blocky icon size
            scale={2}
            style={{width: 32, height: 32}} // style of the view will wrap the icon
          />
          <View style={[styles.textContainer, {width: '80%'}]}>
            <Text h4 style={{flex: 1, flexWrap: 'wrap'}}>
              {title}
            </Text>
            <Text style={{color: '#797979'}}>{subtitle}</Text>
          </View>
        </View>
        <View style={styles.rightContainer}>{amountTag}</View>
      </View>
    </TouchableOpacity>
  );
};

export default ListItem;
