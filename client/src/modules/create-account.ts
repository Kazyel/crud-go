import renderResponse from "./render-response";

const initializeAccountCreation = () => {
  const submitForm = async (form: HTMLFormElement) => {
    const formData = new FormData(form);

    const response = await fetch("http://localhost:8080/api/v1/users/", {
      method: form.method.toUpperCase(),
      body: JSON.stringify({
        name: formData.get("name"),
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

  const createAccountForm =
    document.querySelector<HTMLFormElement>("#create-account-form");

  if (createAccountForm) {
    createAccountForm.addEventListener("submit", (event) => {
      event.preventDefault();
      submitForm(createAccountForm);
    });
  }
};

export default initializeAccountCreation;
