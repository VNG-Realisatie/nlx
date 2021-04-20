// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React, { useContext, useEffect, useState } from 'react'
import { observer } from 'mobx-react'
import { useTranslation } from 'react-i18next'
import { Alert, Button, ToasterContext } from '@commonground/design-system'
import { Link, useLocation, useHistory } from 'react-router-dom'
import { useOrderStore } from '../../../hooks/use-stores'
import PageTemplate from '../../../components/PageTemplate'
import LoadingMessage from '../../../components/LoadingMessage'
import { IconPlus } from '../../../icons'
import OrdersViewPage from './OrdersViewPage'
import OrdersEmptyView from './OrdersEmptyView'
import { StyledActionsBar } from './index.styles'

const OrdersPage = () => {
  const { t } = useTranslation()
  const { showToast } = useContext(ToasterContext)
  const location = useLocation()
  const history = useHistory()
  const orderStore = useOrderStore()
  const [error, setError] = useState()

  useEffect(() => {
    const fetchData = async () => {
      try {
        await orderStore.fetchAll()
      } catch (err) {
        setError(err.message)
      }
    }

    fetchData()
  }, []) // eslint-disable-line react-hooks/exhaustive-deps

  useEffect(() => {
    const searchParams = new URLSearchParams(location.search)
    const lastAction = searchParams.get('lastAction')
    if (!lastAction) return

    showToast({
      title: t('Order created successfully'),
      variant: 'success',
    })

    history.replace('/orders')
  }, [location.search, history, showToast, t])

  return (
    <PageTemplate>
      <PageTemplate.Header
        title={t('Orders')}
        description={t('Consume services on behalf of another organization.')}
      />

      <StyledActionsBar>
        <Button as={Link} to="/orders/add-order" aria-label={t('Add order')}>
          <IconPlus inline />
          {t('Add order')}
        </Button>
      </StyledActionsBar>

      {orderStore.isLoading ? (
        <LoadingMessage />
      ) : error ? (
        <Alert
          variant="error"
          data-testid="error-message"
          title={t('Failed to load orders')}
        >
          {error}
        </Alert>
      ) : orderStore.orders.length ? (
        <OrdersViewPage orders={orderStore.orders} />
      ) : (
        <OrdersEmptyView />
      )}
    </PageTemplate>
  )
}

export default observer(OrdersPage)
