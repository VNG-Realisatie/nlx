// i18next-parser.config.js

const options = require('./src/i18n.options')

module.exports = {
  contextSeparator: '_',
  // Key separator used in your translation keys

  createOldCatalogs: true,
  // Save the \_old files

  defaultNamespace: options.defaultNS,
  // Default namespace used in your i18next config

  defaultValue: '',
  // Default value to give to empty keys

  indentation: 2,
  // Indentation of the catalog files

  keepRemoved: true,
  // Keep keys from the catalog that are no longer in code

  keySeparator: options.keySeparator,
  // Key separator used in your translation keys
  // If you want to use plain english keys, separators such as `.` and `:` will conflict. You might want to set `keySeparator: false` and `namespaceSeparator: false`. That way, `t('Status: Loading...')` will not think that there are a namespace and three separator dots for instance.

  // see below for more details
  lexers: {
    htm: ['HTMLLexer'],
    html: ['HTMLLexer'],

    js: ['JsxLexer'],
    default: ['JavascriptLexer'],
  },

  lineEnding: 'auto',
  // Control the line ending. See options at https://github.com/ryanve/eol

  locales: [options.lng, options.fallbackLng],
  // An array of the locales in your applications

  namespaceSeparator: options.nsSeparator,
  // Namespace separator used in your translation keys
  // If you want to use plain english keys, separators such as `.` and `:` will conflict. You might want to set `keySeparator: false` and `namespaceSeparator: false`. That way, `t('Status: Loading...')` will not think that there are a namespace and three separator dots for instance.

  output: 'public/i18n/$LOCALE/$NAMESPACE.json',
  // Supports $LOCALE and $NAMESPACE injection
  // Supports JSON (.json) and YAML (.yml) file formats
  // Where to write the locale files relative to process.cwd()

  input: ['src/**/*.js'],
  // An array of globs that describe where to look for source files
  // relative to the location of the configuration file

  reactNamespace: false,
  // For react file, extract the defaultNamespace - https://react.i18next.com/components/translate-hoc.html
  // Ignored when parsing a `.jsx` file and namespace is extracted from that file.

  sort: true,
  // Whether or not to sort the catalog

  useKeysAsDefaultValue: true,
  // Whether to use the keys as the default value; ex. "Hello": "Hello", "World": "World"
  // The option `defaultValue` will not work if this is set to true

  verbose: true,
  // Display info about the parsing including some stats
}
