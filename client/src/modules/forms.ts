import initializeAccountCreation from "./create-account";
import initializeLogin from "./login";

const buildForms = () => {
  const createLoginFormComponent = () => {
    const form = document.createElement("form");
    form.id = "login-form";
    form.method = "POST";

    form.innerHTML = `
      <label for="email">Email</label>
      <input type="email" name="email" placeholder="Your email" />
  
      <label for="password">Password</label>
      <input type="password" name="password" placeholder="Your password" />
    
      <button type="submit">Submit</button>
    `;

    return form;
  };

  const createAccountFormComponent = () => {
    const form = document.createElement("form");
    form.id = "create-account-form";
    form.method = "POST";

    form.innerHTML = `
    <label for="name">Name</label>
    <input type="text" name="name" placeholder="Your name" />
    
    <label for="email">Email</label>
    <input type="email" name="email" placeholder="Your email" />
    
    <label for="password">Password</label>
    <input type="password" name="password" placeholder="Your password" />
    
    <button type="submit">Submit</button>
    `;

    return form;
  };

  const createLoginLinkComponent = () => {
    const link = document.createElement("a");
    link.id = "login-link";
    link.innerHTML = "Already have an account? Login here.";
    link.addEventListener("click", (event) => {
      event.preventDefault();
      switchForms();
    });

    return link;
  };

  const createAccountLinkComponent = () => {
    const link = document.createElement("a");
    link.id = "create-account-link";
    link.innerHTML = "Or create an account...";
    link.addEventListener("click", (event) => {
      event.preventDefault();
      switchForms();
    });

    return link;
  };

  const switchForms = () => {
    const formContainer = document.querySelector<HTMLDivElement>("#form-container");
    if (formContainer) {
      const currentForm = formContainer.querySelector<HTMLFormElement>("form");
      const currentAnchor = formContainer.querySelector<HTMLAnchorElement>("a");

      if (currentForm && currentAnchor) {
        if (currentForm.id === "create-account-form") {
          currentForm.replaceWith(createLoginFormComponent());
          currentAnchor.replaceWith(createAccountLinkComponent());
          initializeLogin();
          return;
        }

        currentForm.replaceWith(createAccountFormComponent());
        currentAnchor.replaceWith(createLoginLinkComponent());
        initializeAccountCreation();
        return;
      }
    }
  };

  const initialAnchor = document.querySelector<HTMLAnchorElement>("#create-account-link");
  initialAnchor?.addEventListener("click", (event) => {
    event.preventDefault();
    switchForms();
  });
};

export default buildForms;
