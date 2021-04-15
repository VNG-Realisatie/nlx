// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React, { useState } from 'react'
import { useTranslation } from 'react-i18next'
import { useHistory } from 'react-router-dom'
import { Alert } from '@commonground/design-system'
import PageTemplate from '../../../components/PageTemplate'
import { useOrderStore } from '../../../hooks/use-stores'
import OrderFormContainer from './components/OrderFormContainer'

const AddOrderPage = () => {
  const { t } = useTranslation()
  const { create } = useOrderStore()
  const history = useHistory()
  const [error, setError] = useState(null)

  const submitOrder = async (formData) => {
    try {
      await create(formData)
      history.push(`/orders?lastAction=added`)
    } catch (err) {
      setError(err.message)
    }
  }

  return (
    <PageTemplate>
      <PageTemplate.HeaderWithBackNavigation
        backButtonTo="/orders"
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
