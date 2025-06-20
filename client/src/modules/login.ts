import renderResponse from "./render-response";

const initializeLogin = () => {
  const submitForm = async (form: HTMLFormElement) => {
    const formData = new FormData(form);

    const response = await fetch("http://localhost:8080/api/v1/auth/login", {
      method: form.method.toUpperCase(),
      body: JSON.stringify({
        email: formData.get("email"),
        password: formData.get("password"),
      }),
    });

    if (response.ok) {
      renderResponse(await response.json());
      return;
    }

    renderResponse(await response.json());
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
