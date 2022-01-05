// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//
import React from 'react'
import { Alert } from '@commonground/design-system'
import { useTranslation } from 'react-i18next'
import { Route, Routes, useParams } from 'react-router-dom'
import LoadingMessage from '../../../../components/LoadingMessage'
import { useOutwayStore } from '../../../../hooks/use-stores'
import OutwayDetailPage from '../../OutwayDetailPage'
import Outways from './components/Outways'

const OutwaysView = () => {
  const { t } = useTranslation()
  const outwayStore = useOutwayStore()
  const { name } = useParams()

  return outwayStore.isFetching ? (
    <LoadingMessage />
  ) : outwayStore.error ? (
    <Alert variant="error" data-testid="error-message">
      {t('Failed to load the outways')}
    </Alert>
  ) : (
    <>
      <Outways outways={outwayStore.outways} selectedOutwayName={name} />
      <Routes>
        <Route path=":name" element={<OutwayDetailPage />} />
      </Routes>
    </>
  )
}

export default OutwaysView
