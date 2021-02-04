// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { useTranslation } from 'react-i18next'
import PageTemplate from '../../components/PageTemplate'

const AuditLogPage = () => {
  const { t } = useTranslation()
  return (
    <PageTemplate>
      <PageTemplate.Header
        title={t('Audit log')}
        description={t('History of all mutations within your NLX setup.')}
      />
    </PageTemplate>
  )
}

export default AuditLogPage
