class User {
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

export default User;
