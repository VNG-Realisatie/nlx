// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React, { useContext, useEffect } from 'react'
import { useParams, useNavigate } from 'react-router-dom'
import { observer } from 'mobx-react'
import { Alert, Drawer, ToasterContext } from '@commonground/design-system'
import { useTranslation } from 'react-i18next'
import { useServiceStore } from '../../../hooks/use-stores'
import serviceActions from '../ServicesPage/serviceActions'
import ServiceDetailView from './ServiceDetailView'

const ServiceDetailPage = () => {
  const { name } = useParams()
  const { t } = useTranslation()
  const { showToast } = useContext(ToasterContext)
  const navigate = useNavigate()
  const serviceStore = useServiceStore()

  useEffect(() => {
    serviceStore.fetch({ name })
  }, [name, serviceStore])

  const close = () => navigate('/services')
  const service = serviceStore.getByName(name)

  const handleRemove = async () => {
    try {
      await serviceStore.removeService(service.name)
      navigate(`../${name}?lastAction=${serviceActions.REMOVED}`)
    } catch (err) {
      let message = err.message

      if (err.response && err.response.status === 403) {
        message = t(`You don't have the required permission.`)
      }

      showToast({
        title: t('Failed to remove the service'),
        body: message,
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

export default observer(ServiceDetailPage)
