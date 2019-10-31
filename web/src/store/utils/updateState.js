// Updating state immutable

export const updateState = (oldState, newValues) => ({
  ...oldState,
  ...newValues,
});
