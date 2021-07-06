// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//
import React from 'react'
import { useTranslation } from 'react-i18next'
import { Centered } from './index.styles'

const OrdersIncomingEmpty = (props) => {
  const { t } = useTranslation()
  return (
    <Centered>
      <h3>
        <small>{t("You haven't received any orders yet")}</small>
      </h3>
      <p>
        <small>
          {t(
            'Use this to review and accept requests made on behalve of your services',
          )}
        </small>
      </p>
    </Centered>
  )
}

export default OrdersIncomingEmpty
