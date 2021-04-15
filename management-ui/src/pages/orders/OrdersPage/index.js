// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React, { useContext, useEffect } from 'react'
import { useTranslation } from 'react-i18next'
import { Button, ToasterContext } from '@commonground/design-system'
import { Link, useLocation, useHistory } from 'react-router-dom'
import PageTemplate from '../../../components/PageTemplate'
import { IconPlus } from '../../../icons'
import { StyledActionsBar, Centered } from './index.styles'

const OrdersPage = () => {
  const { t } = useTranslation()
  const { showToast } = useContext(ToasterContext)
  const location = useLocation()
  const history = useHistory()

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

      <Centered>
        <h3>
          <small>{t('Overview not yet available')}</small>
        </h3>
        <p>
          <small>
            {t(
              'In the current version of NLX Management it is not yet possible to view the orders you created.',
            )}
          </small>
        </p>
      </Centered>
    </PageTemplate>
  )
}

export default OrdersPage
