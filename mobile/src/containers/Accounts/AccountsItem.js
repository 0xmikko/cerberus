import React from 'react';
import ListItem from '../../components/ListItem';
import {TouchableOpacity} from 'react-native';

const TransactionItem = ({item, onSelect}) => {

  return (
    <ListItem
      icon={item.icon}
      title={item.address || ''}
      onSelect={() => onSelect(item.id)}
    />
  );
};

export default TransactionItem;
