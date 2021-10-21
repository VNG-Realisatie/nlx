// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import React from 'react'
import { useTranslation } from 'react-i18next'
import { Alert, Button } from '@commonground/design-system'
import { Route, useParams } from 'react-router-dom'
import { observer } from 'mobx-react'
import PageTemplate from '../../../components/PageTemplate'
import InwayDetailPage from '../InwayDetailPage'
import LoadingMessage from '../../../components/LoadingMessage'
import { useInwayStore, useOutwayStore } from '../../../hooks/use-stores'
import { ActionsBar, ActionsBarButton } from './index.styles'
import Inways from './Inways'

const OverviewPage = () => {
  const { t } = useTranslation()
  const { isInitiallyFetched, inways, error, getInway } = useInwayStore()
  const outwayStore = useOutwayStore()
  const { name } = useParams()

  return (
    <PageTemplate>
      <PageTemplate.Header
        title={t('Inways')}
        description={t(
          'Gateways to provide (Inways) or consume (Outways) services.',
        )}
      />

      <ActionsBar>
        <Button
          as={ActionsBarButton}
          aria-label={t('Show all')}
          to="/inways"
          variant="secondary"
        >
          {t('Show all')} ({outwayStore.outways.length + inways.length})
        </Button>
        <Button
          as={ActionsBarButton}
          aria-label={t('Show Inways')}
          variant="secondary"
          to="/inways"
        >
          {t('Inways')} ({inways.length})
        </Button>
        <Button
          as={ActionsBarButton}
          aria-label={t('Show Outways')}
          variant="secondary"
          to="/inways"
        >
          {t('Outways')} ({outwayStore.outways.length})
        </Button>
      </ActionsBar>

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
        path="/inways/:name"
        render={({ match }) => {
          const inway = getInway({ name: match.params.name })

          if (inway) {
            inway.fetch()
          }

          return <InwayDetailPage parentUrl="/inways" inway={inway} />
        }}
      />
    </PageTemplate>
  )
}

export default observer(OverviewPage)
