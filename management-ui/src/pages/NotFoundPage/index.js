// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { useTranslation, Trans } from 'react-i18next'
import PageTemplate from '../../components/PageTemplate'

import { NotFoundContainer, StyledIconErrorCircle } from './index.styles'

const NotFoundPage = () => {
  const { t } = useTranslation()
  return (
    <PageTemplate>
      <NotFoundContainer>
        <StyledIconErrorCircle title={t('Error')} />

        <h1>{t('Page not found')}</h1>
        <p>
          <Trans i18nKey="404 body" />
        </p>
      </NotFoundContainer>
    </PageTemplate>
  )
}

export default NotFoundPage
