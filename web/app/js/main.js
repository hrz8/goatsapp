import 'flowbite';
import './dark-mode';
import './sidebar';

import {initFlowbite} from 'flowbite';
import {Signal} from 'signal-polyfill';

import {CodeMirror} from './codemirror';
import {htmxForm} from './htmx-form';
import {initProjectSelector, onLoadProjectSelector} from './project-switch';

window.Signal = Signal;
window.initFlowbite = initFlowbite;
window.htmxForm = htmxForm;
window.CodeMirror = CodeMirror;
window.onLoadProjectSelector = onLoadProjectSelector;

document.body.addEventListener('htmx:load', function () {
  initFlowbite();
  initProjectSelector();
});

document.body.addEventListener('htmx:pushedIntoHistory', function () {
  initProjectSelector();
});
