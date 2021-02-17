// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { observer } from 'mobx-react'
import { Button } from '@commonground/design-system'
import { useTranslation, Trans } from 'react-i18next'
import PageTemplate from '../../components/PageTemplate'
import { useFinanceStore } from '../../hooks/use-stores'
import { IconDownload } from '../../icons'
import { Centered, StyledIconExternalLink } from './index.styles'

const FinancePage = () => {
  const financeStore = useFinanceStore()
  const { t } = useTranslation()

  async function download(e) {
    e.preventDefault()

    const result = await financeStore.downloadExport()
    const a = document.createElement('a')

    a.download = 'report.csv'
    a.href = `data:text/csv;base64,${result.data || ''}`

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

      {financeStore.enabled ? (
        <Centered>
          <p>
            <small>{t('This report is only available as CSV file')}</small>
          </p>

          <Button variant="secondary" onClick={download}>
            <IconDownload inline />
            {t('Export report')}
          </Button>
        </Centered>
      ) : (
        <Centered>
          <h3>
            <small>{t('Configure the transaction log')}</small>
          </h3>
          <p>
            <small>
              <Trans i18nKey="finance_configure">
                To create a financial report, you need to configure the
                transaction log.
              </Trans>
            </small>
          </p>
          <Button
            as="a"
            variant="link"
            href="https://docs.nlx.io/use-nlx/enable-pricing"
          >
            {t('To NLX Docs')} <StyledIconExternalLink />
          </Button>
        </Centered>
      )}
    </PageTemplate>
  )
}

export default observer(FinancePage)
