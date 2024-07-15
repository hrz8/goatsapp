import 'flowbite';
import './dark-mode';
import './sidebar';
import './project-switch';

import {initFlowbite} from 'flowbite';
import {Signal} from 'signal-polyfill';

import {CodeMirror} from './codemirror';

window.Signal = Signal;
window.initFlowbite = initFlowbite;
window.CodeMirror = CodeMirror;

function InitProjectForm() {
  const form = document.querySelector('#add-project-form');
  form.addEventListener('htmx:beforeOnLoad', function (e) {
    if (e.detail.xhr.status === 422) {
      e.detail.shouldSwap = true;
      e.detail.isError = false;
    }
  });
  form.addEventListener('htmx:beforeRequest', function (e) {
    const btn = e.detail.elt.querySelector('button[type="submit"');
    btn.classList.add('hidden');
  });
  form.addEventListener('htmx:afterRequest', function (e) {
    const btn = e.detail.elt.querySelector('button[type="submit"');
    btn.classList.remove('hidden');
    btn.classList.add('block');

    if (e.detail.xhr.status === 200 || e.detail.xhr.status === 422) {
      initFlowbite(); // need to re-init since gonna doing getInstance()
      const modal = FlowbiteInstances.getInstance('Modal', 'add-project-modal');

      if (e.detail.xhr.status === 200) {
        modal.hide();
        e.detail.elt.reset();
      }

      setTimeout(() => location.reload(), 2000);
    }
  });
}

window.addEventListener('load', function () {
  InitProjectForm();
});

document.body.addEventListener('htmx:pushedIntoHistory', function () {
  // because the form (Modal) removed and replaced every time url history pushed
  // need to re-init the listener
  InitProjectForm();
});

document.body.addEventListener('htmx:load', function () {
  initFlowbite();
});
