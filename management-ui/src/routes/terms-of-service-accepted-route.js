// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React, { useContext } from 'react'
import { Navigate } from 'react-router-dom'
import { node } from 'prop-types'
import ToSContext from '../tos-context'

const TermsOfServiceAcceptedRoute = ({ children }) => {
  const { isReady, tos } = useContext(ToSContext)

  if (!isReady) {
    return null
  }

  if (!tos.enabled || tos.accepted) {
    return children
  }

  return <Navigate to="/terms-of-service" />
}

TermsOfServiceAcceptedRoute.propTypes = {
  children: node,
}

export default TermsOfServiceAcceptedRoute
