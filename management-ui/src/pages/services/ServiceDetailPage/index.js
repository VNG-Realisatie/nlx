// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { string, shape } from 'prop-types'
import { useParams, useHistory } from 'react-router-dom'
import { Alert, Drawer } from '@commonground/design-system'
import { useTranslation } from 'react-i18next'
import pick from 'lodash.pick'
import { serviceModelPropTypes } from '../../../models/ServiceModel'
import { useServicesStore } from '../../../hooks/use-stores'
import serviceActions from '../ServicesPage/serviceActions'
import ServiceDetailView from './ServiceDetailView'

const ServiceDetailPage = ({ parentUrl, service }) => {
  const { name } = useParams()
  const { t } = useTranslation()
  const history = useHistory()
  const { removeService } = useServicesStore()

  const close = () => history.push(parentUrl)
  const handleRemove = async () => {
    await removeService(service)
    history.push(`/services/${name}?lastAction=${serviceActions.REMOVED}`)
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
        {service ? (
          <ServiceDetailView service={service} removeHandler={handleRemove} />
        ) : (
          <Alert variant="error" data-testid="error-message">
            {t('Failed to load the service', { name })}
          </Alert>
        )}
      </Drawer.Content>
    </Drawer>
  )
}

ServiceDetailPage.propTypes = {
  parentUrl: string,
  service: shape(pick(serviceModelPropTypes, ['name'])),
}

ServiceDetailPage.defaultProps = {
  parentUrl: '/services',
}

export default ServiceDetailPage
