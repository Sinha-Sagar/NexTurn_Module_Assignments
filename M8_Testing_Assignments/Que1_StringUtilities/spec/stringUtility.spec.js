const { capitalize, reverseString } = require("../src/stringUtility");

// Ques-1
describe("capitalize function Test", () => {
    it("should capitalize the first letter of a word", () => {
      expect(capitalize("hello")).toBe("Hello");
      expect(capitalize("world")).toBe("World");
    });
  
    it("should handle empty strings", () => {
      expect(capitalize("")).toBe("");
    });
  
    it("should handle single-character words", () => {
      expect(capitalize("a")).toBe("A");
      expect(capitalize("z")).toBe("Z");
    });
});

// Ques-2
describe("reverseString function Test", () => {
    it("should reverse a string", () => {
      expect(reverseString("hello")).toBe("olleh");
      expect(reverseString("world")).toBe("dlrow");
    });
  
    it("should handle empty strings", () => {
      expect(reverseString("")).toBe("");
    });
  
    it("should handle palindromes", () => {
      expect(reverseString("madam")).toBe("madam");
    });
});
