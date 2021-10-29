// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import React from 'react'
import { useTranslation } from 'react-i18next'
import { Alert } from '@commonground/design-system'
import { Route, useLocation, useParams } from 'react-router-dom'
import { observer } from 'mobx-react'
import PageTemplate from '../../../components/PageTemplate'
import InwayDetailPage from '../InwayDetailPage'
import LoadingMessage from '../../../components/LoadingMessage'
import { useInwayStore } from '../../../hooks/use-stores'
import Inways from './Inways'

const OverviewPage = () => {
  const { t } = useTranslation()
  const { isInitiallyFetched, inways, error, getInway } = useInwayStore()
  const { name } = useParams()
  const { pathname } = useLocation()

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
        <Inways inways={inways} selectedInwayName={name} />
      )}

      <Route
        path={`${pathname}/:name`}
        render={({ match }) => {
          const inway = getInway({ name: match.params.name })

          if (inway) {
            inway.fetch()
          }

          return <InwayDetailPage parentUrl={pathname} inway={inway} />
        }}
      />
    </PageTemplate>
  )
}

export default observer(OverviewPage)
