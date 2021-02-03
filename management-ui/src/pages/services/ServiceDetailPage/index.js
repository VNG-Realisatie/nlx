// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React, { useContext } from 'react'
import { string, shape } from 'prop-types'
import { useParams, useHistory } from 'react-router-dom'
import { Alert, Drawer, ToasterContext } from '@commonground/design-system'
import { useTranslation } from 'react-i18next'
import { useServicesStore } from '../../../hooks/use-stores'
import serviceActions from '../ServicesPage/serviceActions'
import ServiceDetailView from './ServiceDetailView'

const ServiceDetailPage = ({ parentUrl, service }) => {
  const { name } = useParams()
  const { t } = useTranslation()
  const { showToast } = useContext(ToasterContext)
  const history = useHistory()
  const { removeService } = useServicesStore()

  const close = () => history.push(parentUrl)
  const handleRemove = async () => {
    try {
      await removeService(service.name)
      history.push(`/services/${name}?lastAction=${serviceActions.REMOVED}`)
    } catch (err) {
      showToast({
        title: t('Failed to remove the service'),
        body: err.message,
        variant: 'error',
      })
    }
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
  service: shape({
    name: string.isRequired,
  }),
}

ServiceDetailPage.defaultProps = {
  parentUrl: '/services',
}

export default ServiceDetailPage
