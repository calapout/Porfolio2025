module.exports = {
  root: true,
  env: {
    browser: true,
    es2022: true,
    node: true,
    // Enables defineProps/defineEmits/etc in <script setup>
    'vue/setup-compiler-macros': true,
  },
  parser: 'vue-eslint-parser',
  parserOptions: {
    // Let vue-eslint-parser delegate <script> blocks to TS parser
    parser: '@typescript-eslint/parser',
    ecmaVersion: 'latest',
    sourceType: 'module',
    extraFileExtensions: ['.vue'],
  },
  extends: [
    'eslint:recommended',
    'plugin:vue/vue3-recommended',
    'plugin:@typescript-eslint/recommended',
  ],
  plugins: ['vue', '@typescript-eslint'],
  rules: {
    // Prefer the TS version of no-unused-vars
    'no-unused-vars': 'off',
    '@typescript-eslint/no-unused-vars': [
      'warn',
      { argsIgnorePattern: '^_', varsIgnorePattern: '^_' },
    ],

    // Commonly disabled for App.vue, Index.vue, etc.
    'vue/multi-word-component-names': 'off',
  },
  ignorePatterns: ['node_modules/', 'dist/'],
};
