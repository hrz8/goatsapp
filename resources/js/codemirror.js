import {basicSetup} from '@codemirror/basic-setup';
import {defaultKeymap, history, historyKeymap} from '@codemirror/commands';
import {json} from '@codemirror/lang-json';
import {defaultHighlightStyle, syntaxHighlighting} from '@codemirror/language';
import {EditorState} from '@codemirror/state';
import {EditorView, keymap} from '@codemirror/view';
import {dracula} from 'thememirror';

function CodeMirror() {
  let view = null;
  return {
    init(container) {
      if (!container) {
        return;
      }

      const extensions = [
        dracula,
        basicSetup,
        json(),
        history(),
        syntaxHighlighting(defaultHighlightStyle),
        keymap.of([defaultKeymap, historyKeymap]),
      ];
      const state = EditorState.create({
        doc: `{\n  "key": "value"\n}`,
        extensions,
      });
      view = new EditorView({state});
      container.innerHTML = '';
      container.append(view.dom);
    },
    value() {
      return view.state.sliceDoc();
    },
  };
}

export {CodeMirror};
