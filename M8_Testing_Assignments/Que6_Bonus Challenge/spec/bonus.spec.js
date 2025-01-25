const { JSDOM } = require("jsdom");
const ApiService = require("../src/apiService");
const { fetchAndDisplayUser } = require("../src/bouns");

describe("fetchAndDisplayUser function", () => {
  let element;
  let apiService;

  beforeEach(() => {
    const { window } = new JSDOM(`<!DOCTYPE html><p>Hello world</p>`);
    element = window.document.createElement("div");

    // Mock apiService
    apiService = new ApiService();
  });

  // Test for successful user fetch
  it("should display the user's name when user data is fetched successfully", async () => {
    const mockUser = { name: "John Doe" };
    spyOn(apiService, "getUser").and.returnValue(Promise.resolve(mockUser));

    await fetchAndDisplayUser(apiService, 1, element);

    expect(apiService.getUser).toHaveBeenCalledWith(1);
    expect(element.textContent).toBe("Hello, John Doe");
  });

  // Test for failed user fetch due to invalid data
  it("should display an error message if the user data is invalid", async () => {
    const mockUser = {}; // No name field
    spyOn(apiService, "getUser").and.returnValue(Promise.resolve(mockUser));

    await fetchAndDisplayUser(apiService, 1, element);

    expect(apiService.getUser).toHaveBeenCalledWith(1);
    expect(element.textContent).toBe("Invalid user data");
  });

  // Test for failed user fetch due to network error
  it("should display an error message if the user fetch fails", async () => {
    spyOn(apiService, "getUser").and.returnValue(Promise.reject(new Error("Network error")));

    await fetchAndDisplayUser(apiService, 1, element);

    expect(apiService.getUser).toHaveBeenCalledWith(1);
    expect(element.textContent).toBe("Network error");
  });
});
