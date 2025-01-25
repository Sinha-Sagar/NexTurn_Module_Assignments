const { delayedGreeting } = require("../src/asyncFunction");

describe("delayedGreeting function", () => {
  // Test for the resolved greeting message
  it("should resolve with the correct greeting message", (done) => {
    delayedGreeting("John", 1000).then((message) => {
      expect(message).toBe("Hello, John!");
      done();
    });
  });

  // Test to validate the function respects the delay using mock timers
  it("should call the function after the specified delay", () => {
    jasmine.clock().install();

    const delay = 1000;
    const promise = delayedGreeting("John", delay);

    let resolvedMessage = null;
    promise.then((message) => {
      resolvedMessage = message;
    });

    jasmine.clock().tick(delay);

    setTimeout(() => {
        expect(resolvedMessage).toBe("Hello, John!");
        jasmine.clock().uninstall();
    }, 0);

  });
});
