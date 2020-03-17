import React from 'react'
import { useTranslation } from 'react-i18next'
import PageTemplate from '../../components/PageTemplate'

const ServicesPage = () => {
  const { t } = useTranslation()

  return (
    <PageTemplate title={t('Services')}>
      <p>{t('An overview of the services here.')}</p>
    </PageTemplate>
  )
}

export default ServicesPage
