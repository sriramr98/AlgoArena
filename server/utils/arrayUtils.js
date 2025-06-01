const areArraysEqual = (arr1, arr2) => {
  if (!Array.isArray(arr1) || !Array.isArray(arr2)) {
    return false;
  }

  if (arr1.length !== arr2.length) {
    return false;
  }

  const valueOccurences = {};
  for (const value of arr1) {
    if (valueOccurences[value]) {
      valueOccurences[value]++;
    } else {
      valueOccurences[value] = 1;
    }
  }
  for (const value of arr2) {
    if (valueOccurences[value]) {
      valueOccurences[value]--;
    } else {
      return false;
    }
  }
  for (const key in valueOccurences) {
    if (valueOccurences[key] !== 0) {
      return false;
    }
  }
  return true;
};

module.exports = {
  areArraysEqual,
};
