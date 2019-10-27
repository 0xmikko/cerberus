import { StyleSheet } from 'react-native';

export const styles = StyleSheet.create({
  container: {
    flex: 1,
    paddingTop: 18,
    paddingBottom: 10,
    paddingLeft: 15,
    paddingRight: 5,
    marginTop: 1,
    flexDirection: 'row',
    alignContent: 'space-between',
    justifyContent: 'space-between',
    backgroundColor: 'white',
  },
  leftContainer: {
    alignItems: 'center',
    padding: 0,
    flexDirection: 'row',
  },
  textContainer: {
    paddingLeft: 15,
    paddingRight: 10,
    alignItems: 'stretch',
    alignContent: 'space-between',
    marginBottom: 5,
  },
  rightContainer: {
    width: 100,
    alignItems: 'flex-end',
    alignContent: 'flex-end',
    paddingRight: 15,
  },
});
