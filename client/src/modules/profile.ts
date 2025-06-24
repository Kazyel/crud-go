import type { App } from "../main";

const renderProfileView = (app: App) => {
  const user = app.getUserContext().getUser();
  if (!user || !user.isLoggedIn) {
    return;
  }

  const formContainer = document.querySelector<HTMLDivElement>("#container");
  if (!formContainer) {
    return;
  }

  const profileContainer = document.createElement("div");
  profileContainer.id = "profile-container";
  profileContainer.classList.add("grid-item");

  profileContainer.innerHTML = `
    <div class="profile-header">
      <h1>You are logged in, this is your profile</h1>
      <button id="logout-button">Logout</button>
    </div>
  `;

  formContainer.replaceWith(profileContainer);
  const logoutButton = document.querySelector<HTMLButtonElement>("#logout-button");

  if (logoutButton) {
    logoutButton.addEventListener("click", (event) => {
      event.preventDefault();
      app.logout();
    });
  }
};

export default renderProfileView;
