// Copyright © VNG Realisatie 2019
// Licensed under the EUPL

import React, { useState, useEffect } from 'react'
import PropTypes from 'prop-types';
 
const VersionLogger = ({logger}) => {
    const [versionTag, setVersionTag] = useState(null)
    
    useEffect(
        () => {
        const load = async () => {
            const result = await fetch('/version.json')
            const {tag} = await result.json()
            setVersionTag(tag)
        }
        load()
    },
    [], // prevent inifinite rerenders
    )
    
    if (!!versionTag) {
        logger(versionTag)
    }
    return null
}

VersionLogger.propTypes = {
    logger: PropTypes.func,
}

VersionLogger.defaultProps = {
    logger: console.log,
}

export default VersionLogger
