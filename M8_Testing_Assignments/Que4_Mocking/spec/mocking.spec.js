const { sendNotification } = require("../src/mocking");

describe("sendNotification function", () => {
  let notificationService;

  beforeEach(() => {
    notificationService = {
      send: jasmine.createSpy(),
    };
  });

  // Test for successful notification sending
  it("should return 'Notification Sent' when the notification is sent successfully", () => {
    notificationService.send.and.returnValue(true);

    const result = sendNotification(notificationService, "Test message");
    expect(result).toBe("Notification Sent");
    expect(notificationService.send).toHaveBeenCalledWith("Test message");
  });

  // Test for failed notification sending
  it("should return 'Failed to Send' when the notification fails to send", () => {
    notificationService.send.and.returnValue(false);

    const result = sendNotification(notificationService, "Test message");
    expect(result).toBe("Failed to Send");
    expect(notificationService.send).toHaveBeenCalledWith("Test message");
  });
});
