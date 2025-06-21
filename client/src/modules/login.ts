import type GlobalState from "./global-state";
import type { LoginResponse } from "./types";

import renderResponse from "./render-response";
import User from "./user";

export const login = async (state: GlobalState, data: LoginResponse) => {
  state.setUser(new User(data.data.user_id, "", ""));
  state.setCSRFToken(data.data.csrf_token);
  state.getUser()?.setIsLoggedIn(true);

  window.localStorage.setItem("user", JSON.stringify(state.getUser()!.getID()));
  window.localStorage.setItem("csrfToken", state.getCSRFToken()!);
};

const initializeLogin = (state: GlobalState) => {
  const submitForm = async (form: HTMLFormElement) => {
    const formData = new FormData(form);

    const response = await fetch("http://localhost:8080/api/v1/auth/login", {
      method: form.method.toUpperCase(),
      body: JSON.stringify({
        email: formData.get("email"),
        password: formData.get("password"),
      }),
    });

    const data = await response.json();

    if (response.ok) {
      renderResponse(data);
      login(state, data);
      return;
    }

    renderResponse(data);
  };

  const loginForm = document.querySelector<HTMLFormElement>("#login-form");

  if (loginForm) {
    loginForm.addEventListener("submit", (event) => {
      event.preventDefault();
      submitForm(loginForm);
    });
  }
};

export default initializeLogin;
