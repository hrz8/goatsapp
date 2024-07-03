document.addEventListener('project:switch', function (e) {
  // uncheck all checked
  document.querySelectorAll('.select-project.checked').forEach((checkedEl) => {
    checkedEl.classList.remove('checked');
  });
  // do htmx request to set project_id cookie
  const form = e.detail.elt.closest('form');
  const submit = form.querySelector('button[type=submit]');
  submit.click();
});
