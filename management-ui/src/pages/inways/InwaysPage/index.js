// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import React from 'react'
import { useTranslation } from 'react-i18next'
import { Alert } from '@commonground/design-system'
import { Route } from 'react-router-dom'

import { observer } from 'mobx-react'
import PageTemplate from '../../../components/PageTemplate'
import InwayDetailPage from '../InwayDetailPage'
import LoadingMessage from '../../../components/LoadingMessage'
import { useInwaysStore } from '../../../hooks/use-stores'
import InwaysPageView from './InwaysPageView'

const InwaysPage = () => {
  const { t } = useTranslation()
  const { isInitiallyFetched, inways, error, selectInway } = useInwaysStore()

  return (
    <PageTemplate>
      <PageTemplate.Header
        title={t('Inways')}
        description={t('Gateways to provide services')}
      />

      {!isInitiallyFetched ? (
        <LoadingMessage />
      ) : error ? (
        <Alert variant="error" data-testid="error-message">
          {t('Failed to load the inways')}
        </Alert>
      ) : (
        <InwaysPageView inways={inways} />
      )}

      <Route
        path="/inways/:name"
        render={({ match }) => (
          <InwayDetailPage
            parentUrl="/inways"
            inway={selectInway(match.params.name)}
          />
        )}
      />
    </PageTemplate>
  )
}

export default observer(InwaysPage)
