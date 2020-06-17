// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { func, string } from 'prop-types'
import { useParams, useHistory } from 'react-router-dom'
import { Alert, Drawer } from '@commonground/design-system'
import { useTranslation } from 'react-i18next'
import ServiceRepository from '../../domain/service-repository'
import ServiceDetails from '../../components/ServiceDetails'
import usePromise from '../../hooks/use-promise'
import LoadingMessage from '../../components/LoadingMessage'

const ServiceDetailPage = ({
  getServiceByName,
  removeService,
  refreshHandler,
  parentUrl,
}) => {
  const { name } = useParams()
  const { t } = useTranslation()
  const history = useHistory()
  const { isReady, error, result: service } = usePromise(getServiceByName, name)
  const close = () => history.push(parentUrl)

  const handleRemove = () => {
    removeService(service)
    refreshHandler()
  }

  return (
    <Drawer noMask closeHandler={close}>
      <Drawer.Header
        as="header"
        title={name}
        closeButtonLabel={t('Close')}
        data-testid="service-name"
      />

      <Drawer.Content>
        {!isReady || (!error && !service) ? (
          <LoadingMessage />
        ) : error ? (
          <Alert variant="error" data-testid="error-message">
            {t('Failed to load the service.', { name })}
          </Alert>
        ) : service ? (
          <ServiceDetails service={service} removeHandler={handleRemove} />
        ) : null}
      </Drawer.Content>
    </Drawer>
  )
}

ServiceDetailPage.propTypes = {
  getServiceByName: func,
  refreshHandler: func,
  removeService: func,
  parentUrl: string,
}

ServiceDetailPage.defaultProps = {
  getServiceByName: ServiceRepository.getByName,
  removeService: ServiceRepository.remove,
  refreshHandler: () => {},
  parentUrl: '/services',
}

export default ServiceDetailPage
