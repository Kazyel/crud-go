import "./css/style.css";
import "./css/reset.css";

import type { LoginResponse } from "./types/types.ts";

import buildForms from "./modules/forms.ts";
import renderProfileView from "./modules/profile.ts";
import { createUserContext, User, UserState } from "./context/user-context.ts";

interface AppDependencies {
  userContext?: UserState;
  container?: HTMLElement;
}

function renderMainView(container: HTMLElement) {
  if (container) {
    container.innerHTML = `  
      <div class="grid">
        <div id="container" class="grid-item">
              <form id="login-form" method="POST">
                <label for="email">Email</label>
                <input type="email" name="email" placeholder="Your email" />

                <label for="password">Password</label>
                <input type="password" name="password" placeholder="Your password" />
              
                <button type="submit">Submit</button>
              </form>

              <a id="create-account-link">Or create an account...</a>
          </div>
          
          <div id="response-container" class="grid-item">
          </div>
      </div>`;
  }
}

export class App {
  private userContext: UserState;
  private container: HTMLElement;

  constructor(dependencies: AppDependencies) {
    this.userContext = dependencies.userContext || createUserContext();
    this.container = dependencies.container || document.querySelector("#app")!;
  }

  getUserContext() {
    return this.userContext;
  }

  getContainer() {
    return this.container;
  }

  render() {
    renderMainView(this.container);

    if (this.userContext.getUser()?.isLoggedIn) {
      renderProfileView(this);
      return;
    }

    buildForms(this);
  }

  async login(data: LoginResponse) {
    this.userContext.setUser(new User(data.data.user_id, "", ""));
    this.userContext.setCSRFToken(data.data.csrf_token);
    this.userContext.getUser()?.setIsLoggedIn(true);

    window.localStorage.setItem(
      "user",
      JSON.stringify(this.userContext.getUser()!.getID())
    );
    window.localStorage.setItem("csrfToken", this.userContext.getCSRFToken()!);
    renderProfileView(this);
  }

  async logout() {
    const response = await fetch("http://localhost:8080/api/v1/auth/logout", {
      method: "POST",
    });

    if (!response.ok) {
      return;
    }

    this.userContext.setUser(null);
    this.userContext.setCSRFToken(null);
    this.userContext.getUser()?.setIsLoggedIn(false);

    window.localStorage.removeItem("user");
    window.localStorage.removeItem("csrfToken");
    this.render();
  }
}

document.addEventListener("DOMContentLoaded", () => {
  const app = new App({
    container: document.querySelector<HTMLDivElement>("#app")!,
  });
  app.render();
});
