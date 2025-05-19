module.exports = {
  locales: ['en', 'ru'],
  output: 'public/locales/$LOCALE/$NAMESPACE.json',
  input: ['src/**/*.{js,jsx}'],
  sort: true,
  keySeparator: false,
  namespaceSeparator: false,
  defaultValue: (locale, namespace, key) => key,
  indentation: 2,
  lineEnding: '\n',
  keepRemoved: false,
  verbose: true
};