// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//
import React from 'react'
import { arrayOf, func, shape, string } from 'prop-types'
import { Formik } from 'formik'
import * as Yup from 'yup'
import { useTranslation } from 'react-i18next'
import {
  Button,
  FieldLabel,
  Fieldset,
  Select,
  TextInput,
} from '@commonground/design-system'
import DateInput, { isoDateSchema } from '../../../../../components/DateInput'
import { DateInputsWrapper, DateInputWrapper, StyledForm } from './index.styles'

const serviceMapper = (services) =>
  services.map((service) => ({
    value: service,
    label: `${service.organization.name} (${service.organization.serialNumber}) - ${service.service}`,
  }))

const EditOrderForm = ({ order, services, onSubmitHandler }) => {
  const { t } = useTranslation()

  const initialValues = {
    description: order.description,
    publicKeyPEM: order.publicKeyPem,
    validFrom: new Date(order.validFrom).toISOString().split('T')[0],
    validUntil: new Date(order.validUntil).toISOString().split('T')[0],
    services: [
      {
        organisation: {
          name: 'Organization-B',
          serialNumber: '12345678901234567891',
        },
        service: 'tester',
      },
    ],
  }

  const validationSchema = Yup.object().shape({
    description: Yup.string()
      .max(100, t('Maximum of n characters allowed', { n: 100 }))
      .required(t('This field is required')),
    publicKeyPEM: Yup.string().required(t('This field is required')),
    validFrom: isoDateSchema(t('Invalid date')).required(
      t('This field is required'),
    ),
    validUntil: isoDateSchema().required(t('This field is required')),
    services: Yup.array()
      .of(
        Yup.object().shape({
          organization: Yup.object().shape({
            serialNumber: Yup.string(),
            name: Yup.string(),
          }),
          service: Yup.string(),
        }),
      )
      .min(1, t('This field is required')),
  })

  const selectableServices = serviceMapper(services)

  const handleSubmit = (values) => {
    onSubmitHandler(values)
  }

  return (
    <Formik
      initialValues={initialValues}
      validationSchema={validationSchema}
      onSubmit={(values) => {
        handleSubmit({
          ...values,
          validFrom: new Date(values.validFrom),
          validUntil: new Date(values.validUntil),
        })
      }}
    >
      <StyledForm>
        <Fieldset>
          <TextInput name="description" size="l">
            <FieldLabel
              label={t('Order description')}
              small={t('For your own reference')}
            />
          </TextInput>

          <TextInput name="publicKeyPEM" size="l" type="textarea">
            <FieldLabel
              label={t('Public key PEM')}
              small={t('The public key of the delegated organization as PEM')}
            />
          </TextInput>

          <DateInputsWrapper>
            <DateInputWrapper>
              <DateInput name="validFrom" label={t('Valid from')} size="s" />
            </DateInputWrapper>
            <DateInputWrapper>
              <DateInput name="validUntil" label={t('Valid until')} size="s" />
            </DateInputWrapper>
          </DateInputsWrapper>

          <Select
            name="services"
            options={selectableServices}
            size="l"
            isMulti
            placeholder={t('Select a service…')}
          >
            <FieldLabel
              label={t('Services')}
              small={t(
                'Is the service not listed? Then please first request access to this service in the directory',
              )}
            />
          </Select>
        </Fieldset>

        <Button type="submit">{t('Update order')}</Button>
      </StyledForm>
    </Formik>
  )
}

const serviceShape = shape({
  organization: shape({
    serialNumber: string.isRequired,
    name: string.isRequired,
  }).isRequired,
  service: string,
})

EditOrderForm.propTypes = {
  order: shape({
    description: string,
    publicKeyPEM: string,
    validFrom: string,
    validUntil: string,
    services: arrayOf(serviceShape),
  }),
  services: arrayOf(serviceShape),
  onSubmitHandler: func.isRequired,
}

EditOrderForm.defaultProps = {
  services: [],
}

export default EditOrderForm
