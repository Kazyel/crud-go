import type { App } from "../main";
import type { LoginResponse } from "../types/types";

import renderResponse from "./render-response";

type BodyObject = {
  [key: string]: FormDataEntryValue;
};

const createBody = (formData: FormData) => {
  let bodyObject: BodyObject = {};
  formData.forEach((value, key) => {
    bodyObject[key] = value;
  });

  return bodyObject;
};

const createFormComponent = (
  id: string,
  method: string,
  url: string,
  innerHTML: string,
  app: App
) => {
  const form = document.createElement("form");
  form.id = id;
  form.method = method;
  form.innerHTML = innerHTML;
  form.addEventListener("submit", (event) => {
    event.preventDefault();
    submitForm(app, form, url);
  });
  return form;
};

const createLink = (id: string, innerHTML: string, app: App) => {
  const link = document.createElement("a");
  link.id = id;
  link.innerHTML = innerHTML;
  link.addEventListener("click", (event) => {
    event.preventDefault();
    switchForms(app);
  });
  return link;
};

const buildLoginForm = (
  app: App,
  currentForm: HTMLFormElement,
  currentAnchor: HTMLAnchorElement
) => {
  currentForm.replaceWith(
    createFormComponent(
      "create-account-form",
      "POST",
      "http://localhost:8080/api/v1/users/",
      ` <label for="name">Name</label>
        <input type="text" name="name" placeholder="Your name" />
        <label for="email">Email</label>
        <input type="email" name="email" placeholder="Your email" />
        <label for="password">Password</label>
        <input type="password" name="password" placeholder="Your password" />
        <button type="submit">Submit</button>
        `,
      app
    )
  );
  currentAnchor.replaceWith(
    createLink("login-link", "Already have an account? Login here.", app)
  );
};

const buildCreateAccountForm = (
  app: App,
  currentForm: HTMLFormElement,
  currentAnchor: HTMLAnchorElement
) => {
  currentForm.replaceWith(
    createFormComponent(
      "login-form",
      "POST",
      "http://localhost:8080/api/v1/auth/login",
      `
      <label for="email">Email</label>
      <input type="email" name="email" placeholder="Your email" />
      <label for="password">Password</label>
      <input type="password" name="password" placeholder="Your password" />
      <button type="submit">Submit</button>
      `,
      app
    )
  );
  currentAnchor.replaceWith(
    createLink("create-account-link", "Or create an account...", app)
  );
};

export const submitForm = async (state: App, form: HTMLFormElement, url: string) => {
  const formData = new FormData(form);

  const response = await fetch(url, {
    method: form.method.toUpperCase(),
    body: JSON.stringify(createBody(formData)),
  });

  const data: LoginResponse = await response.json();
  if (!response.ok) {
    renderResponse(data);
    return;
  }

  renderResponse(data);
  state.login(data);
};

const switchForms = (app: App) => {
  const formContainer = document.querySelector<HTMLDivElement>("#container");
  if (!formContainer) {
    return;
  }

  const currentForm = formContainer.querySelector<HTMLFormElement>("form");
  const currentAnchor = formContainer.querySelector<HTMLAnchorElement>("a");
  if (!currentForm || !currentAnchor) {
    return;
  }

  switch (currentForm.id) {
    case "login-form":
      buildLoginForm(app, currentForm, currentAnchor);
      break;
    case "create-account-form":
      buildCreateAccountForm(app, currentForm, currentAnchor);
      break;
    default:
      return;
  }
};

const buildForms = (app: App) => {
  const loginForm = document.querySelector<HTMLFormElement>("#login-form");
  if (loginForm) {
    loginForm.addEventListener("submit", (event) => {
      event.preventDefault();
      submitForm(app, loginForm, "http://localhost:8080/api/v1/auth/login");
    });
  }

  const initialAnchor = document.querySelector<HTMLAnchorElement>("#create-account-link");
  initialAnchor?.addEventListener("click", (event) => {
    event.preventDefault();
    switchForms(app);
  });
};

export default buildForms;
