// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React, { useContext, useEffect, useState } from 'react'
import { observer } from 'mobx-react'
import { useTranslation } from 'react-i18next'
import {
  Alert,
  Button,
  ToasterContext,
  Spinner,
} from '@commonground/design-system'
import { Link, useLocation, useHistory, useParams } from 'react-router-dom'
import { useOrderStore } from '../../../hooks/use-stores'
import PageTemplate from '../../../components/PageTemplate'
import LoadingMessage from '../../../components/LoadingMessage'
import { IconPlus, IconRefresh } from '../../../icons'
import Outgoing from './Outgoing'
import Incoming from './Incoming'
import { ActionsBar, StyledButton } from './index.styles'

const OrdersPage = () => {
  const { t } = useTranslation()
  const { showToast } = useContext(ToasterContext)
  const location = useLocation()
  const params = useParams()
  const history = useHistory()
  const orderStore = useOrderStore()
  const [error, setError] = useState()
  const [isRefreshLoading, setRefreshLoading] = useState(false)

  useEffect(() => {
    const fetchOrders = async () => {
      try {
        await Promise.all([
          orderStore.fetchOutgoing(),
          orderStore.fetchIncoming(),
        ])
      } catch (err) {
        setError(err.message)
      }
    }
    fetchOrders()
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

  const updateIncomingOrders = async () => {
    setRefreshLoading(true)

    const totalIncomingOrders = orderStore.incomingOrders?.length
    await orderStore.updateIncoming()

    const newIncomingOrders =
      orderStore.incomingOrders?.length - totalIncomingOrders

    setTimeout(() => {
      setRefreshLoading(false)
      showToast({
        title: t(`Overview updated`),
        body: `${newIncomingOrders || t('No')} ${t('new orders found')}`,
        variant: 'success',
      })
    }, 400)
  }

  return (
    <PageTemplate>
      <PageTemplate.Header
        title={t('Orders')}
        description={t('Consume services on behalf of another organization.')}
      />

      <ActionsBar>
        <Button
          as={StyledButton}
          aria-label={t('Issued')}
          to="/orders/outgoing"
          variant="secondary"
        >
          {t('Issued')} ({orderStore.outgoingOrders.length})
        </Button>
        <Button
          as={StyledButton}
          aria-label={t('Received')}
          variant="secondary"
          to="/orders/incoming"
        >
          {t('Received')} ({orderStore.incomingOrders.length})
        </Button>

        <Button
          aria-label={t('Update overview')}
          disabled={isRefreshLoading}
          onClick={updateIncomingOrders}
          variant="secondary"
        >
          {isRefreshLoading ? <Spinner /> : <IconRefresh inline />}
          {t('Update overview')}
        </Button>
        <Button as={Link} to="/orders/add-order" aria-label={t('Add order')}>
          <IconPlus inline />
          {t('Add order')}
        </Button>
      </ActionsBar>

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
      ) : params.type === 'outgoing' ? (
        <Outgoing orders={orderStore.outgoingOrders} />
      ) : params.type === 'incoming' ? (
        <Incoming orders={orderStore.incomingOrders} />
      ) : null}
    </PageTemplate>
  )
}

export default observer(OrdersPage)
