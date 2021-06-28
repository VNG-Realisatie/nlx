// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//
import React from 'react'
import { useTranslation } from 'react-i18next'
import { Centered } from './index.styles'

const OrdersOutgoingEmptyView = (props) => {
  const { t } = useTranslation()
  return (
    <Centered>
      <h3>
        <small>{t("You don't have any issued orders yet")}</small>
      </h3>
      <p>
        <small>
          {t(
            'Use this to allow other organizations to request certain services on your behalve',
          )}
        </small>
      </p>
    </Centered>
  )
}

export default OrdersOutgoingEmptyView
