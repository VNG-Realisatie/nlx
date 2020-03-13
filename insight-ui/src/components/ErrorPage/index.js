// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

import React from 'react'
import { string, node } from 'prop-types'
import ErrorMessage from '../ErrorMessage'

const ErrorPage = ({ ...props }) => <ErrorMessage {...props} />

ErrorPage.propTypes = {
  title: string.isRequired,
  children: node,
}

export default ErrorPage
