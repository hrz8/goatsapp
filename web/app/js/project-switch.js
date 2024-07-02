import {initFlowbite} from 'flowbite';

import {htmxForm} from './htmx-form';

function clearCheckedRadioBtnProject() {
  document.querySelectorAll('.select-project.checked').forEach((checkedEl) => {
    checkedEl.classList.remove('checked');
  });
}

function switchProject(radio) {
  const form = radio.closest('form');
  const submit = form.querySelector('button[type=submit]');
  submit.click();
}

function onLoadProjectSelector(dropdown) {
  document.querySelectorAll(dropdown + ' .select-project').forEach((el) => {
    el.addEventListener('click', function () {
      clearCheckedRadioBtnProject();
      switchProject(this);
    });
  });
}

function initProjectSelector() {
  htmxForm(document.querySelector('#add-project-form'), {
    before(e) {
      const btn = e.detail.elt.querySelector('button[type="submit"');
      btn.classList.add('hidden');
    },
    after(e) {
      initFlowbite();
      const btn = e.detail.elt.querySelector('button[type="submit"');
      const toast = FlowbiteInstances.getInstance('Dismiss', 'toast');
      const modal = FlowbiteInstances.getInstance('Modal', 'add-project-modal');
      btn.classList.remove('hidden');
      btn.classList.add('block');

      if (e.detail.xhr.status === 200 || e.detail.xhr.status === 422) {
        if (e.detail.xhr.status === 200) {
          modal.hide();
          e.detail.elt.reset();
        }

        setTimeout(() => {
          toast.hide();
          location.reload();
        }, 2000);
      }
    },
  });
}

export {initProjectSelector, onLoadProjectSelector};
