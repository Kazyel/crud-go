import type { LoginResponse } from "../types/types";

import renderProfileView from "../modules/profile";

export class User {
  #id: string;
  name: string;
  email: string;
  isLoggedIn: boolean;

  constructor(id: string, name: string, email: string) {
    this.#id = id;
    this.name = name;
    this.email = email;
    this.isLoggedIn = false;
  }

  getID() {
    return this.#id;
  }

  setIsLoggedIn(isLoggedIn: boolean) {
    this.isLoggedIn = isLoggedIn;
  }
}

export class UserState {
  #user: User | null;
  #csrfToken: string | null;

  constructor(user: User | null = null, csrfToken: string | null = null) {
    this.#user = user;
    this.#csrfToken = csrfToken;
  }

  getUser() {
    return this.#user || null;
  }

  setUser(user: User | null) {
    this.#user = user;
  }

  getCSRFToken() {
    return this.#csrfToken || null;
  }

  setCSRFToken(csrfToken: string | null) {
    this.#csrfToken = csrfToken;
  }

  loadState() {
    const user = window.localStorage.getItem("user");
    const csrfToken = window.localStorage.getItem("csrfToken");

    if (user && csrfToken) {
      this.#user = new User(user, "", "");
      this.#user.setIsLoggedIn(true);
      this.#csrfToken = csrfToken;
    }
  }

  login(data: LoginResponse) {
    this.setUser(new User(data.data.user_id, "", ""));
    this.setCSRFToken(data.data.csrf_token);
    this.getUser()?.setIsLoggedIn(true);

    window.localStorage.setItem("user", JSON.stringify(this.getUser()!.getID()));
    window.localStorage.setItem("csrfToken", this.getCSRFToken()!);
    renderProfileView(this);
  }

  async logout() {
    const response = await fetch("http://localhost:8080/api/v1/auth/logout", {
      method: "POST",
    });

    if (!response.ok) {
      return;
    }

    this.setUser(null);
    this.setCSRFToken(null);
    this.getUser()?.setIsLoggedIn(false);

    window.localStorage.removeItem("user");
    window.localStorage.removeItem("csrfToken");
    window.location.reload();
  }
}

export const createUserContext = (): UserState => {
  const userState = new UserState();
  userState.loadState();
  return userState;
};
