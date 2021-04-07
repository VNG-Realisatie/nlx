// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React, { useContext, useState } from 'react'
import { useTranslation } from 'react-i18next'
import { Alert, ToasterContext } from '@commonground/design-system'
import PageTemplate from '../../components/PageTemplate'
import { useOrderStore } from '../../hooks/use-stores'
import OrderFormContainer from './components/OrderFormContainer'

const AddOrderPage = () => {
  const { t } = useTranslation()
  const { create } = useOrderStore()
  const { showToast } = useContext(ToasterContext)
  const [error, setError] = useState(null)

  const submitOrder = async (formData) => {
    try {
      await create(formData)
      showToast({
        title: t('Order created successfully'),
        variant: 'success',
      })
    } catch (err) {
      setError(err.message)
    }
  }

  return (
    <PageTemplate>
      <PageTemplate.HeaderWithBackNavigation
        backButtonTo="/"
        title={t('New order')}
      />

      {error ? (
        <Alert
          title={t('Failed to add order')}
          variant="error"
          data-testid="error-message"
        >
          {error}
        </Alert>
      ) : null}

      <OrderFormContainer onSubmitHandler={submitOrder} />
    </PageTemplate>
  )
}

export default AddOrderPage
