export class User {
  private id: string;
  name: string;
  email: string;
  isLoggedIn: boolean;

  constructor(id: string, name: string, email: string) {
    this.id = id;
    this.name = name;
    this.email = email;
    this.isLoggedIn = false;
  }

  getID() {
    return this.id;
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
}

export const createUserContext = (): UserState => {
  const userState = new UserState();
  userState.loadState();
  return userState;
};
