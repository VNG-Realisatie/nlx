// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import { useState, useEffect } from 'react'
import PropTypes from 'prop-types'

const VersionLogger = ({ logger }) => {
  const [versionTag, setVersionTag] = useState(null)

  useEffect(
    () => {
      const load = async () => {
        try {
          const result = await fetch('/version.json')
          const { tag } = await result.json()
          setVersionTag(tag)
        } catch (e) {}
      }
      load()
    },
    [], // prevent inifinite rerenders
  )

  if (versionTag) {
    logger(versionTag)
  }
  return null
}

VersionLogger.propTypes = {
  logger: PropTypes.func,
}

VersionLogger.defaultProps = {
  // eslint-disable-next-line no-console
  logger: console.log,
}

export default VersionLogger
