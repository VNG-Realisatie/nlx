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
import OutwayDetailPage from '../OutwayDetailPage'
import LoadingMessage from '../../../components/LoadingMessage'
import { useInwayStore, useOutwayStore } from '../../../hooks/use-stores'
import {
  ActionsBar,
  ActionsBarButton,
  StyledIconInway,
  StyledIconOutway,
} from './index.styles'
import Inways from './Inways'
import Outways from './Outways'

const InwaysAndOutwaysPage = () => {
  const { t } = useTranslation()
  const { isInitiallyFetched, error } = useInwayStore()
  const inwayStore = useInwayStore()
  const outwayStore = useOutwayStore()
  const params = useParams()

  return (
    <PageTemplate>
      <PageTemplate.Header
        title={t('Inways and Outways')}
        description={t(
          'Gateways to provide (Inways) or consume (Outways) services.',
        )}
      />

      <ActionsBar>
        <Button
          as={ActionsBarButton}
          aria-label={t('Show Inways')}
          variant="secondary"
          to="/inways-and-outways/inways"
        >
          <StyledIconInway /> {t('Inways')} ({inwayStore.inways.length})
        </Button>
        <Button
          as={ActionsBarButton}
          aria-label={t('Show Outways')}
          variant="secondary"
          to="/inways-and-outways/outways"
        >
          <StyledIconOutway /> {t('Outways')} ({outwayStore.outways.length})
        </Button>
      </ActionsBar>

      {!isInitiallyFetched ? (
        <LoadingMessage />
      ) : error ? (
        <Alert variant="error" data-testid="error-message">
          {t('Failed to load the inways')}
        </Alert>
      ) : params.type === 'inways' ? (
        <Inways inways={inwayStore.inways} selectedInwayName={params.name} />
      ) : params.type === 'outways' ? (
        <Outways
          outways={outwayStore.outways}
          selectedOutwayName={params.name}
        />
      ) : null}

      <Route
        path="/inways-and-outways/inways/:name"
        render={({ match }) => {
          const inway = inwayStore.getInway({ name: match.params.name })

          if (inway) {
            inway.fetch()
          }

          return (
            <InwayDetailPage
              parentUrl="/inways-and-outways/inways"
              inway={inway}
            />
          )
        }}
      />

      <Route
        path="/inways-and-outways/outways/:name"
        render={({ match }) => {
          const outway = outwayStore.getByName({ name: match.params.name })

          return (
            <OutwayDetailPage
              parentUrl="/inways-and-outways/outways"
              outway={outway}
            />
          )
        }}
      />
    </PageTemplate>
  )
}

export default observer(InwaysAndOutwaysPage)
