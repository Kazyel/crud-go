import type User from "./user";

class GlobalState {
  #user: User | null;
  #csrfToken: string | null;

  constructor(user: User | null = null, csrfToken: string | null = null) {
    this.#user = user;
    this.#csrfToken = csrfToken;
  }

  getUser() {
    return this.#user || null;
  }

  getCSRFToken() {
    return this.#csrfToken || null;
  }

  setUser(user: User) {
    this.#user = user;
  }

  setCSRFToken(csrfToken: string) {
    this.#csrfToken = csrfToken;
  }
}

export default GlobalState;
