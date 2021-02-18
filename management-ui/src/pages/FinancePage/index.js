// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React, { useState, useEffect } from 'react'
import { Link } from 'react-router-dom'
import { Button, Alert } from '@commonground/design-system'
import { useTranslation } from 'react-i18next'
import PageTemplate from '../../components/PageTemplate'
import { IconDownload } from '../../icons'
import { CenteredExport } from './index.styles'

const FinancePage = () => {
  const [csvUri, setCsvUri] = useState('')
  const [error, setError] = useState('')
  const { t } = useTranslation()

  useEffect(() => {
    const fetchData = async () => {
      try {
        setCsvUri('FAKE')
        setError('')
      } catch (err) {
        setCsvUri('')
        setError(t('Error'))
      }
    }

    fetchData()
  }, []) // eslint-disable-line react-hooks/exhaustive-deps

  return (
    <PageTemplate>
      <PageTemplate.Header
        title={t('Finances')}
        description={t(
          'Insight in how often and for what cost other organizations request your services',
        )}
      />

      <CenteredExport>
        <p>
          <small>{t('This report is only available as CSV file')}</small>
        </p>

        {csvUri && (
          <Button variant="secondary" as={Link} to={csvUri}>
            <IconDownload inline />
            {t('Export report')}
          </Button>
        )}

        {error && <Alert variant="error">{error}</Alert>}
      </CenteredExport>
    </PageTemplate>
  )
}

export default FinancePage
