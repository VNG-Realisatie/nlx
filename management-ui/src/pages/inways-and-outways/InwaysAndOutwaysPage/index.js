// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import React, { useEffect } from 'react'
import { useTranslation } from 'react-i18next'
import { Alert, Button } from '@commonground/design-system'
import { Route, useParams } from 'react-router-dom'
import { observer } from 'mobx-react'
import PageTemplate from '../../../components/PageTemplate'
import InwayDetailPage from '../InwayDetailPage'
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
  const inwayStore = useInwayStore()
  const outwayStore = useOutwayStore()
  const params = useParams()

  useEffect(() => {
    const fetchData = async () => {
      await inwayStore.fetchInways()
      await outwayStore.fetchAll()
    }

    fetchData()
  }, []) // eslint-disable-line react-hooks/exhaustive-deps

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

      {params.type === 'inways' ? (
        inwayStore.isFetching ? (
          <LoadingMessage />
        ) : inwayStore.error ? (
          <Alert variant="error" data-testid="error-message">
            {t('Failed to load the inways')}
          </Alert>
        ) : (
          <Inways inways={inwayStore.inways} selectedInwayName={params.name} />
        )
      ) : params.type === 'outways' ? (
        outwayStore.isFetching ? (
          <LoadingMessage />
        ) : outwayStore.error ? (
          <Alert variant="error" data-testid="error-message">
            {t('Failed to load the outways')}
          </Alert>
        ) : (
          <Outways
            outways={outwayStore.outways}
            selectedOutwayName={params.name}
          />
        )
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
    </PageTemplate>
  )
}

export default observer(InwaysAndOutwaysPage)
