// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { func, string } from 'prop-types'
import { useParams, useHistory } from 'react-router-dom'
import { Alert, Drawer } from '@commonground/design-system'
import { useTranslation } from 'react-i18next'

import InwayRepository from '../../domain/inway-repository'
import InwayDetails from '../../components/InwayDetails'
import usePromise from '../../hooks/use-promise'
import LoadingMessage from '../../components/LoadingMessage'

const InwayDetailPage = ({ getInwayByName, parentUrl }) => {
  const { name } = useParams()
  const { t } = useTranslation()
  const history = useHistory()
  const { isReady, error, result: inway } = usePromise(getInwayByName, name)
  const close = () => history.push(parentUrl)

  return (
    <Drawer noMask closeHandler={close}>
      <Drawer.Header
        title={name}
        closeButtonLabel={t('Close')}
        data-testid="gateway-name"
      />

      <Drawer.Content>
        {!isReady || (!error && !inway) ? (
          <LoadingMessage />
        ) : error ? (
          <Alert variant="error" data-testid="error-message">
            {t('Failed to load the details for this inway.', { name })}
          </Alert>
        ) : inway ? (
          <InwayDetails inway={inway} />
        ) : null}
      </Drawer.Content>
    </Drawer>
  )
}

InwayDetailPage.propTypes = {
  getInwayByName: func,
  parentUrl: string,
}

InwayDetailPage.defaultProps = {
  getInwayByName: InwayRepository.getByName,
  parentUrl: '/inways',
}

export default InwayDetailPage
