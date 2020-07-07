// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { useTranslation } from 'react-i18next'
import PageTemplate from '../../components/PageTemplate'

import { NotFoundContainer, StyledIconErrorCircle } from './index.styles'

const NotFoundPage = () => {
  const { t } = useTranslation()
  return (
    <PageTemplate>
      <NotFoundContainer>
        <StyledIconErrorCircle title={t('Error')} />

        <h1>{t('Page not found')}</h1>
        <p>{t('We could not find what you were looking for.')}</p>
        <p>
          {t(
            'Please contact the person or organisation that linked you here and let them know their link is broken.',
          )}
        </p>
      </NotFoundContainer>
    </PageTemplate>
  )
}

export default NotFoundPage
