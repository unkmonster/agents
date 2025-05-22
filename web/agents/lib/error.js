export class NotLoggedInError extends Error {
  constructor(message = "User is not logged in") {
    super(message);
    this.name = "NotLoggedInError";
  }
}

export class TokenInvalidError extends Error {
  constructor(message = "Token is invalid") {
    super(message);
    this.name = "TokenInvalidError";
  }
}

export class AppError extends Error {
  constructor(data) {
    super(data.message);
    this.name = "AppError";
    this.reason = data.reason;
    this.code = data.code;
    this.metadata = data.metadata;
  }

  toJSON() {
    return {
      message: this.message,
      reason: this.reason,
      code: this.code,
    };
  }
}
