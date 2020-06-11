// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
const SemanticReleaseError = require('@semantic-release/error')
const AggregateError = require('aggregate-error')
const glob = require('glob')
const replace = require('replace-in-file')

let files = []

function verifyConditions(pluginConfig, context) {
  const { charts } = pluginConfig

  if (!charts) {
    throw new SemanticReleaseError('Invalid `charts` option.', 'EINVALIDOPTION')
  }

  files = glob.sync(`${charts}/Chart.yaml`)

  if (!files.length) {
    throw new SemanticReleaseError('No Chart.yaml files found.', 'ENOCHARTS')
  }
}

async function prepare(pluginConfig, context) {
  const { dryRun } = pluginConfig
  const { version } = context.nextRelease

  const options = {
    files: files,
    from: [/^version: .*$/m, /^appVersion: .*$/m],
    to: [`version: ${version}`, `appVersion: ${version}`],
    disableGlobs: true,
    countMatches: true,
    dry: dryRun,
  }

  const isVersionReplaced = ({ numMatches, numReplacements }) =>
    numMatches === 2 && numReplacements === 2

  let results

  try {
    results = await replace(options)
  } catch (error) {
    throw new SemanticReleaseError(
      'Faild to replace versions',
      'EFAILEDREPLACEVERSION',
      error,
    )
  }

  const errors = results.reduce(
    (errors, result) =>
      !isVersionReplaced(result)
        ? [
            ...errors,
            new SemanticReleaseError(
              `Failed to update versions in: ${result.file}`,
              'EFAILEDREPLACEVERSION',
            ),
          ]
        : errors,
    [],
  )

  if (errors.length > 0) {
    throw new AggregateError(errors)
  }
}

module.exports = { verifyConditions, prepare }
