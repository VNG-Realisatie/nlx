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
  Select,
  TextInput,
  Fieldset,
} from '@commonground/design-system'
import DateInput, { isoDateSchema } from '../../../../../components/DateInput'
import { StyledForm, DateInputsWrapper, DateInputWrapper } from './index.styles'

const OrderForm = ({ services, onSubmitHandler }) => {
  const { t } = useTranslation()

  const initialValues = {
    description: '',
    reference: '',
    publicKeyPEM: '',
    validFrom: '',
    validUntil: '',
    delegatee: '',
    services: [],
  }

  const validationSchema = Yup.object().shape({
    description: Yup.string()
      .max(100, t('Maximum of n characters allowed', { n: 100 }))
      .required(t('This field is required')),
    reference: Yup.string().required(t('This field is required')),
    publicKeyPEM: Yup.string().required(t('This field is required')),
    delegatee: Yup.string()
      .max(100, t('Maximum of n characters allowed', { n: 100 }))
      .matches(/^[a-zA-Z0-9-. _\s]{1,}$/, t('Please use a URL friendly name'))
      .required(t('This field is required')),
    validFrom: isoDateSchema(t('Invalid date')).required(
      t('This field is required'),
    ),
    validUntil: isoDateSchema().required(t('This field is required')),
    services: Yup.array()
      .of(
        Yup.object().shape({
          organization: Yup.string(),
          service: Yup.string(),
        }),
      )
      .min(1, t('This field is required')),
  })

  const selectableServices = services.map((service) => ({
    value: service,
    label: `${service.organization} - ${service.service}`,
  }))

  const handleSubmit = (values) => {
    onSubmitHandler(values)
  }

  return (
    <Formik
      initialValues={{
        ...initialValues,
      }}
      validationSchema={validationSchema}
      onSubmit={(values) => {
        values.validFrom = new Date(values.validFrom)
        values.validUntil = new Date(values.validUntil)
        handleSubmit(values)
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

          <TextInput name="reference" size="l">
            <FieldLabel
              label={t('Reference')}
              small={t('This identifier is sent with each request')}
            />
          </TextInput>

          <TextInput name="delegatee" size="l">
            <FieldLabel label={t('Delegated organization')} />
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

        <Button type="submit">{t('Add order')}</Button>
      </StyledForm>
    </Formik>
  )
}

OrderForm.propTypes = {
  services: arrayOf(
    shape({
      organization: string,
      service: string,
    }),
  ),
  onSubmitHandler: func.isRequired,
}

OrderForm.defaultProps = {
  services: [],
}

export default OrderForm
