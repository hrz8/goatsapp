import eslint from '@eslint/js';
import eslintConfigPrettier from 'eslint-config-prettier';
import prettierPlugin from 'eslint-plugin-prettier';
import simpleImportSortPlugin from 'eslint-plugin-simple-import-sort';
import globals from 'globals';

const customConfig = {
  files: ['**/*.js'],
  languageOptions: {
    sourceType: 'module',
    globals: {
      ...globals.browser,
      ...globals.jest,
      ...globals.node,
    },
  },
  plugins: {
    'simple-import-sort': simpleImportSortPlugin,
    prettier: prettierPlugin,
  },
  rules: {
    'prettier/prettier': 'error',
    'quote-props': ['error', 'as-needed'],
    curly: 'error',
    'brace-style': [
      'error',
      '1tbs',
      {
        allowSingleLine: false,
      },
    ],
    'keyword-spacing': 'error',
    'semi-spacing': 'error',
    'block-spacing': 'error',
    'space-infix-ops': 'error',
    'space-in-parens': ['error', 'never'],
    'space-before-blocks': 'error',
    'implicit-arrow-linebreak': 'off',
    'linebreak-style': 'off',
    'prefer-const': 'error',
    'arrow-body-style': ['error', 'as-needed'],
    'operator-linebreak': 'off',
    'function-call-argument-newline': ['error', 'consistent'],
    'newline-per-chained-call': ['off'],
    'import/no-cycle': 'off',
    'class-methods-use-this': 'off',
    'import/prefer-default-export': 'off',
    'arrow-parens': ['error', 'always'],
    camelcase: [
      'warn',
      {
        properties: 'never',
      },
    ],
    'comma-dangle': [
      'error',
      {
        arrays: 'only-multiline',
        objects: 'only-multiline',
        imports: 'only-multiline',
        exports: 'only-multiline',
        functions: 'only-multiline',
      },
    ],
    complexity: ['warn', 16],
    'constructor-super': 'error',
    'object-curly-spacing': ['error', 'always'],
    'eol-last': 'error',
    eqeqeq: ['warn', 'smart'],
    'guard-for-in': 'warn',
    'id-blacklist': [
      'error',
      'any',
      'Number',
      'number',
      'String',
      'string',
      'Boolean',
      'boolean',
      'Undefined',
      'undefined',
    ],
    'id-match': 'error',
    'import/no-internal-modules': 'off',
    'max-classes-per-file': 'off',
    'new-parens': 'error',
    'no-unused-vars': 'warn',
    'no-multi-spaces': [
      'error',
      {
        exceptions: {
          ImportDeclaration: true,
          VariableDeclarator: true,
          Property: true,
        },
      },
    ],
    'no-bitwise': 'off',
    'no-caller': 'error',
    'no-cond-assign': 'error',
    'no-console': [
      'error',
      {
        allow: [
          'warn',
          'dir',
          'trace',
          'clear',
          'table',
          'debug',
          'info',
          'error',
          'Console',
        ],
      },
    ],
    'no-debugger': 'error',
    'no-empty': 'error',
    'no-eval': 'error',
    'no-fallthrough': 'off',
    'no-invalid-this': 'off',
    'no-new-wrappers': 'error',
    'no-shadow': [
      'error',
      {
        hoist: 'all',
      },
    ],
    'no-throw-literal': 'error',
    'no-trailing-spaces': 'error',
    'no-whitespace-before-property': 'error',
    'no-undef': 'error',
    'no-undef-init': 'error',
    'no-underscore-dangle': 'off',
    'no-unsafe-finally': 'error',
    'no-unused-labels': 'error',
    'object-shorthand': 'error',
    'one-var': ['off', 'never'],
    radix: 'error',
    'spaced-comment': [
      'error',
      'always',
      {
        markers: ['/'],
      },
    ],
    'use-isnan': 'error',
    'valid-typeof': [
      'error',
      {
        requireStringLiterals: true,
      },
    ],
    'max-len': [
      'warn',
      {
        ignoreComments: true,
        ignoreTrailingComments: true,
        ignoreStrings: true,
        ignoreRegExpLiterals: true,
        code: 80,
      },
    ],
    'padded-blocks': ['error', 'never'],
    'padding-line-between-statements': [
      'error',
      {
        blankLine: 'always',
        prev: '*',
        next: 'block',
      },
      {
        blankLine: 'always',
        prev: 'block',
        next: '*',
      },
      {
        blankLine: 'always',
        prev: '*',
        next: 'block-like',
      },
      {
        blankLine: 'always',
        prev: 'block-like',
        next: '*',
      },
    ],
    'simple-import-sort/imports': 'error',
    'simple-import-sort/exports': 'error',
  },
};

export default [eslint.configs.recommended, customConfig, eslintConfigPrettier];
