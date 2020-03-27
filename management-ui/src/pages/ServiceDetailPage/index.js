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
import { StyledLoadingMessage } from '../ServicesPage/index.styles'
import Spinner from '../ServicesPage/Spinner'

const ServiceDetailPage = ({ getServiceByName, parentUrl }) => {
  const { name } = useParams()
  const { t } = useTranslation()
  const history = useHistory()
  const { loading, error, result } = usePromise(getServiceByName, name)
  const close = () => history.push(parentUrl)

  return (
    <Drawer closeHandler={close}>
      {loading || (!error && !result) ? (
        <StyledLoadingMessage role="progressbar">
          <Spinner /> {t('Loading…')}
        </StyledLoadingMessage>
      ) : error ? (
        <Alert variant="error" data-testid="error-message">
          {t('Failed to load the service.', { name })}
        </Alert>
      ) : result ? (
        <ServiceDetails service={result} />
      ) : null}
    </Drawer>
  )
}

ServiceDetailPage.propTypes = {
  getServiceByName: func,
  parentUrl: string,
}

ServiceDetailPage.defaultProps = {
  getServiceByName: ServiceRepository.getByName,
  parentUrl: '/services',
}

export default ServiceDetailPage
