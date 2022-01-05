// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import React, { useEffect } from 'react'
import { useTranslation } from 'react-i18next'
import { Button } from '@commonground/design-system'
import { Navigate, Route, Routes } from 'react-router-dom'
import { observer } from 'mobx-react'
import PageTemplate from '../../../components/PageTemplate'
import { useInwayStore, useOutwayStore } from '../../../hooks/use-stores'
import {
  ActionsBar,
  ActionsBarButton,
  StyledIconInway,
  StyledIconOutway,
} from './index.styles'
import InwaysView from './InwaysView'
import OutwaysView from './OutwaysView'

const InwaysAndOutwaysPage = () => {
  const { t } = useTranslation()
  const inwayStore = useInwayStore()
  const outwayStore = useOutwayStore()

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
          to="inways"
        >
          <StyledIconInway /> {t('Inways')} ({inwayStore.inways.length})
        </Button>
        <Button
          as={ActionsBarButton}
          aria-label={t('Show Outways')}
          variant="secondary"
          to="outways"
        >
          <StyledIconOutway /> {t('Outways')} ({outwayStore.outways.length})
        </Button>
      </ActionsBar>

      <Routes>
        <Route index element={<Navigate to="inways" />} />
        <Route path="inways/*" element={<InwaysView />} />
        <Route path="outways/*" element={<OutwaysView />} />
      </Routes>
    </PageTemplate>
  )
}

export default observer(InwaysAndOutwaysPage)
