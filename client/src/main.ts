import "./css/style.css";
import "./css/reset.css";

import { createUserContext } from "./context/user-context.ts";
import buildForms from "./modules/forms.ts";
import renderProfileView from "./modules/profile.ts";

export default function renderMainView() {
  const app = document.querySelector<HTMLDivElement>("#app");

  if (app) {
    app.innerHTML = `  
      <div class="grid">
        <div id="form-container" class="grid-item">
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

  return app;
}

document.addEventListener("DOMContentLoaded", () => {
  const userContext = createUserContext();
  renderMainView();

  if (userContext.getUser()?.isLoggedIn) {
    renderProfileView(userContext);
  }

  buildForms(userContext);
});
