import React from 'react';
import ListItem from '../../components/ListItem';
import {TouchableOpacity} from 'react-native';

const TransactionItem = ({item, onSelect}) => {
  const states = {
    0: {
      title: 'Undefined',
      color: 'grey',
    },
    1: {
      title: 'Cancelled',
      color: 'red',
    },
    2: {
      title: 'Confirmed',
      color: 'green',
    },
  };

  return (
    <ListItem
      icon={item.icon}
      title={item.to ? item.to : ''}
      subtitle={states[item.state] ? states[item.state].title : '???'}
      amount={item.amount}
      onSelect={() => onSelect(item.id)}
      color={states[item.state] ? states[item.state].color : 'grey'}
    />
  );
};

export default TransactionItem;
