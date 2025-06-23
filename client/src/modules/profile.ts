import type { UserState } from "../context/user-context";

const renderProfileView = (state: UserState) => {
  const user = state.getUser();
  if (!user || !user.isLoggedIn) {
    return;
  }

  const formContainer = document.querySelector<HTMLDivElement>("#form-container");
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
      state.logout();
    });
  }
};

export default renderProfileView;
