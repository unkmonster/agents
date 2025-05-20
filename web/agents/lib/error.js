export class NotLoggedInError extends Error {
  constructor(message = "User is not logged in") {
    super(message);
    this.name = "NotLoggedInError";
  }
}
