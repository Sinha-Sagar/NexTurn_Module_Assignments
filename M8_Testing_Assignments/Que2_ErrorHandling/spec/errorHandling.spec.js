const { getElement } = require("../src/errorHandling");

describe("getElement function", () => {
  // Tests for valid index values
  it("should return the correct element for valid indices", () => {
    const arr = [1, 2, 3, 4, 5];
    expect(getElement(arr, 0)).toBe(1);
    expect(getElement(arr, 2)).toBe(3);
  });

  // Tests to check if the error is thrown for negative indices
  it("should throw an error for negative indices", () => {
    const arr = [1, 2, 3, 4, 5];
    expect(() => getElement(arr, -1)).toThrowError("Index out of bounds");
    expect(() => getElement(arr, -5)).toThrowError("Index out of bounds");
  });

  // Tests to check if the error is thrown for out-of-range indices
  it("should throw an error for out-of-range indices", () => {
    const arr = [1, 2, 3, 4, 5];
    expect(() => getElement(arr, 5)).toThrowError("Index out of bounds");
    expect(() => getElement(arr, 10)).toThrowError("Index out of bounds");
  });
});
