// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

import React from 'react'
import { oneOf, array } from 'prop-types'
import { useTranslation } from 'react-i18next'

const AuthorizationMode = ({ mode, authorizations }) => {
  const { t } = useTranslation()

  return mode === 'whitelist' ? (
    <span>
      {t('Whitelist')} ({authorizations.length})
    </span>
  ) : (
    <span>{t('Open')}</span>
  )
}

AuthorizationMode.propTypes = {
  mode: oneOf(['none', 'whitelist']).isRequired,
  authorizations: array.isRequired,
}

export default AuthorizationMode
