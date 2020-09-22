// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'

export default (Component, ignoredProps = []) => {
  // eslint-disable-next-line react/display-name
  return (props) => {
    const passableProps = {}

    for (const prop in props) {
      if (!ignoredProps.includes(prop)) {
        passableProps[prop] = props[prop]
      }
    }

    return <Component {...passableProps} />
  }
}
