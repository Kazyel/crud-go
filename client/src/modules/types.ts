export type LoginResponse = {
  success: boolean;
  message: string;
  data: {
    user_id: string;
    csrf_token: string;
  };
};

export type CreateAccountResponse = {
  success: boolean;
  message: string;
  data: {
    user_id: string;
  };
};

export type ErrorResponse = {
  success: boolean;
  error: string;
};
