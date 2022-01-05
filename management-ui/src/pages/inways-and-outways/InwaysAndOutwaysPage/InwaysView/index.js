// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//
import React from 'react'
import { Alert } from '@commonground/design-system'
import { useTranslation } from 'react-i18next'
import { Route, Routes, useParams } from 'react-router-dom'
import LoadingMessage from '../../../../components/LoadingMessage'
import { useInwayStore } from '../../../../hooks/use-stores'
import InwayDetailPage from '../../InwayDetailPage'
import Inways from './components/Inways'

const InwaysView = () => {
  const { t } = useTranslation()
  const inwayStore = useInwayStore()
  const { name } = useParams()

  return inwayStore.isFetching ? (
    <LoadingMessage />
  ) : inwayStore.error ? (
    <Alert variant="error" data-testid="error-message">
      {t('Failed to load the inways')}
    </Alert>
  ) : (
    <>
      <Inways inways={inwayStore.inways} selectedInwayName={name} />

      <Routes>
        <Route path=":name" element={<InwayDetailPage />} />
      </Routes>
    </>
  )
}

export default InwaysView
