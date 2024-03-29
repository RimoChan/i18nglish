export const i18n = (str) => {
  if (typeof str !== "string") {
    throw new Error("Not string");
  }

  return str.replace(/\w+/g, ellipsis);
};

const ellipsis = (word) => {
  if (word.length < 3) {
    return word;
  }

  return `${word.slice(0, 1)}${word.length - 2}${word.slice(-1)}`;
};
