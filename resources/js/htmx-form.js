function htmxForm(form, {before, after}) {
  if (!form) {
    return;
  }

  if (typeof after !== 'function' || typeof before !== 'function') {
    return;
  }

  form.addEventListener('htmx:beforeOnLoad', function (e) {
    if (e.detail.xhr.status === 422) {
      e.detail.shouldSwap = true;
      e.detail.isError = false;
    }
  });
  form.addEventListener('htmx:beforeRequest', before);
  form.addEventListener('htmx:afterRequest', after);
}

export {htmxForm};
