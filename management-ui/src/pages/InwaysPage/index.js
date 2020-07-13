// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import React from 'react'
import { func } from 'prop-types'
import { useTranslation } from 'react-i18next'
import { Alert } from '@commonground/design-system'
import { Route } from 'react-router-dom'

import PageTemplate from '../../components/PageTemplate'
import usePromise from '../../hooks/use-promise'
import InwayRepository from '../../domain/inway-repository'
import InwayDetailPage from '../InwayDetailPage'

import LoadingMessage from '../../components/LoadingMessage'

import InwaysPageView from './InwaysPageView'

const InwaysPage = ({ getInways }) => {
  const { t } = useTranslation()
  const { isReady, error, result: inways } = usePromise(getInways)

  return (
    <PageTemplate>
      <PageTemplate.Header
        title={t('Inways')}
        description={t('Gateways to provide services.')}
      />

      {!isReady ? (
        <LoadingMessage />
      ) : error ? (
        <Alert variant="error" data-testid="error-message">
          {t('Failed to load the inways.')}
        </Alert>
      ) : (
        <InwaysPageView inways={inways} />
      )}

      <Route exact path="/inways/:name">
        <InwayDetailPage parentUrl="/inways" />
      </Route>
    </PageTemplate>
  )
}

InwaysPage.propTypes = {
  getInways: func,
}

InwaysPage.defaultProps = {
  getInways: InwayRepository.getAll,
}

export default InwaysPage
