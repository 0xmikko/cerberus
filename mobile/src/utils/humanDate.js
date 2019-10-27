import moment from 'moment';

export const toHumanDate = date => {
  const today = moment().format('YYYY-MM-DD');

  const yesterday = moment()
    .subtract(1, 'days')
    .format('YYYY-MM-DD');

  switch (date) {
    case today:
      return 'Today';
    case yesterday:
      return 'Yesterday';
    default:
      return moment(date).format('MMMM Do YYYY');
  }
};
