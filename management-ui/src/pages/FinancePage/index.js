// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { Button } from '@commonground/design-system'
import { useTranslation } from 'react-i18next'
import PageTemplate from '../../components/PageTemplate'
import useStores from '../../hooks/use-stores'
import { IconDownload } from '../../icons'
import { CenteredExport } from './index.styles'

const FinancePage = () => {
  const { financeStore } = useStores()
  const { t } = useTranslation()

  async function download(e) {
    e.preventDefault()

    const result = await financeStore.downloadExport()
    const a = document.createElement('a')

    a.download = 'report.csv'
    a.href = `data:text/csv;base64,${result.data}`

    document.body.append(a)

    a.click()

    document.body.removeChild(a)
  }

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

        <Button variant="secondary" onClick={download}>
          <IconDownload inline />
          {t('Export report')}
        </Button>
      </CenteredExport>
    </PageTemplate>
  )
}

export default FinancePage
