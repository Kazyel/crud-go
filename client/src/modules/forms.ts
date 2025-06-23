import type { UserState } from "../context/user-context";
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

export const submitForm = async (
  state: UserState,
  form: HTMLFormElement,
  url: string
) => {
  const formData = new FormData(form);

  const response = await fetch(url, {
    method: form.method.toUpperCase(),
    body: JSON.stringify(createBody(formData)),
  });

  const data = await response.json();
  if (!response.ok) {
    renderResponse(data);
    return;
  }

  renderResponse(data);
  state.login(data);
};

const createFormComponent = (
  id: string,
  method: string,
  url: string,
  innerHTML: string,
  state: UserState
) => {
  const form = document.createElement("form");
  form.id = id;
  form.method = method;
  form.innerHTML = innerHTML;
  form.addEventListener("submit", (event) => {
    event.preventDefault();
    submitForm(state, form, url);
  });
  return form;
};

const createLink = (id: string, innerHTML: string, state: UserState) => {
  const link = document.createElement("a");
  link.id = id;
  link.innerHTML = innerHTML;
  link.addEventListener("click", (event) => {
    event.preventDefault();
    switchForms(state);
  });
  return link;
};

const buildLoginForm = (
  state: UserState,
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
      state
    )
  );
  currentAnchor.replaceWith(
    createLink("login-link", "Already have an account? Login here.", state)
  );
};

const buildCreateAccountForm = (
  state: UserState,
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
      state
    )
  );
  currentAnchor.replaceWith(
    createLink("create-account-link", "Or create an account...", state)
  );
};

const switchForms = (state: UserState) => {
  const formContainer = document.querySelector<HTMLDivElement>("#form-container");
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
      buildLoginForm(state, currentForm, currentAnchor);
      break;
    case "create-account-form":
      buildCreateAccountForm(state, currentForm, currentAnchor);
      break;
    default:
      return;
  }
};

const buildForms = (state: UserState) => {
  const loginForm = document.querySelector<HTMLFormElement>("#login-form");
  if (loginForm) {
    loginForm.addEventListener("submit", (event) => {
      event.preventDefault();
      submitForm(state, loginForm, "http://localhost:8080/api/v1/auth/login");
    });
  }

  const initialAnchor = document.querySelector<HTMLAnchorElement>("#create-account-link");
  initialAnchor?.addEventListener("click", (event) => {
    event.preventDefault();
    switchForms(state);
  });
};

export default buildForms;
