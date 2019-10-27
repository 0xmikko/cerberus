export const numberWithCommas = x => {
  if (!x) return '0.00';
  let parts = x.toFixed(2).split('.');
  parts[0] = parts[0].replace(/\B(?=(\d{3})+(?!\d))/g, ',');
  return parts.join('.');
};

export const numberIntegerPartWithCommas = x => {
  if (!x) return '0';
  let parts = x.toString().split('.');
  parts[0] = parts[0].replace(/\B(?=(\d{3})+(?!\d))/g, ',');
  return parts[0];
};

export const fractionalPart = x => {
  if (!x) return '.00';
  let parts = x.toFixed(2).split('.');
  return parts[1];
};
