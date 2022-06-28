// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
const SemanticReleaseError = require('@semantic-release/error')
const glob = require('glob')
const replace = require('replace-in-file')

let filePaths = []

function verifyConditions(pluginConfig, context) {
  const { files } = pluginConfig

  if (!files) {
    throw new SemanticReleaseError('Invalid `files` option.', 'EINVALIDOPTION')
  }

  filePaths = glob.sync(`${files}`)

  if (!filePaths.length) {
    throw new SemanticReleaseError('No yaml files found.', 'ENOFILES')
  }
}

async function prepare(pluginConfig, context) {
  const { dryRun } = pluginConfig
  const { version } = context.nextRelease

  const options = {
    files: filePaths,
    from: [
      /^(\s*)image: (.*)nlxio\/(.*):v.*$/m,
      /^(\s*)tag: "v.*"$/m
    ],
    to: [
      `$1image: $2nlxio/$3:v${version}`,
      `$1tag: "v${version}"`
    ],
    disableGlobs: true,
    dry: dryRun
  }

  try {
    await replace(options)
  } catch (error) {
    throw new SemanticReleaseError(
      'Failed to replace versions',
      'EFAILEDREPLACEVERSION',
      error,
    )
  }
}

module.exports = { verifyConditions, prepare }
