import 'flowbite';
import './dark-mode';
import './sidebar';

import {Dismiss, Modal} from 'flowbite';
import {Signal} from 'signal-polyfill';

window.Signal = Signal;

const addProjectModal = new Modal(document.querySelector('#add-project-modal'));

document.querySelectorAll('.add-project-button').forEach(function (el) {
  el.addEventListener('click', function (e) {
    e.preventDefault();
  });
});

const addProjectForm = document.querySelector('#add-project-form');
addProjectForm.addEventListener('htmx:beforeOnLoad', function (evt) {
  if (evt.detail.xhr.status === 422) {
    evt.detail.shouldSwap = true;
    evt.detail.isError = false;
  }
});
addProjectForm.addEventListener('htmx:afterRequest', function (e) {
  const backdrop = document.querySelector('body > div[modal-backdrop]');
  const btn = e.detail.elt.querySelector('button[type="submit"');
  btn.classList.remove('hidden');
  btn.classList.add('block');

  if (e.detail.xhr.status === 200 || e.detail.xhr.status === 422) {
    if (e.detail.xhr.status === 200) {
      addProjectModal.hide();
      backdrop.classList.remove('fixed');
      e.detail.elt.reset();
    }

    const dismiss = new Dismiss(
      document.querySelector('#toast'),
      document.querySelector('#toast-dismisser'),
    );
    setTimeout(() => dismiss.hide(), 5000);
  }
});
addProjectForm.addEventListener('htmx:beforeRequest', function (e) {
  const btn = e.detail.elt.querySelector('button[type="submit"');
  btn.classList.add('hidden');
});
