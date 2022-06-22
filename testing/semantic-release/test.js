const prepare = require('../../scripts/semantic-release-installation-guide-yaml').prepare

prepare({
  dryRun: false,
  files: 'technical-docs/nlx-helm-installation-guide/nlx-outway-values.yaml'
}, {
  nextRelease: {
    version: '0.0.42'
  }
})
