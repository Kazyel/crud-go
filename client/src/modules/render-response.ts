const renderResponse = (response: any) => {
  const responseContainer = document.querySelector<HTMLDivElement>("#response-container");

  if (responseContainer) {
    responseContainer.innerHTML = `
        <h2>Response</h2>
        <pre>${JSON.stringify(response, null, 2)}</pre>
      </div>
    `;
  }
};

export default renderResponse;
