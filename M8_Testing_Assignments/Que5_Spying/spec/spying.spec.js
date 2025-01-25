const { JSDOM } = require("jsdom");
const { toggleVisibility } = require("../src/spying");

describe("toggleVisibility function", () => {
  let element;

  beforeEach(() => {
    const { window } = new JSDOM(`<!DOCTYPE html><p>Hello world</p>`);
    element = window.document.createElement("div");
  });

  // Test when the element is initially visible
  it("should set display to 'none' when the element is initially visible", () => {
    element.style.display = "block";
    toggleVisibility(element);
    expect(element.style.display).toBe("none");
  });

  // Test when the element is initially hidden
  it("should set display to 'block' when the element is initially hidden", () => {
    element.style.display = "none";
    toggleVisibility(element);
    expect(element.style.display).toBe("block");
  });

  // Additional test for default style (no display set)
  it("should set display to 'none' if the element's display property is not explicitly set", () => {
    toggleVisibility(element);
    expect(element.style.display).toBe("none");
  });
});
