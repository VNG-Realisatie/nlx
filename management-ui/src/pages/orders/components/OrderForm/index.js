// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//
import React, { useEffect } from 'react'
import { array, bool, func, instanceOf, shape, string } from 'prop-types'
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
import { observer } from 'mobx-react'
import DateInput, { isoDateSchema } from '../../../../components/DateInput'
import {
  useAccessProofStore,
  useDirectoryServiceStore,
  useOutwayStore,
} from '../../../../hooks/use-stores'
import { DateInputsWrapper, DateInputWrapper, StyledForm } from './index.styles'

const formatDateToYYYYMMDD = (date) =>
  date.getFullYear() +
  '-' +
  `${date.getMonth() + 1}`.padStart(2, 0) +
  '-' +
  `${date.getDate()}`.padStart(2, 0)

const DEFAULT_INITIAL_VALUES = {
  description: '',
  reference: '',
  delegatee: '',
  publicKeyPEM: '',
  validFrom: '',
  validUntil: '',
  accessProofIds: [],
}

const OrderForm = ({
  initialValues,
  onSubmitHandler,
  submitButtonText,
  isEditMode,
  ...props
}) => {
  const { t } = useTranslation()
  const outwayStore = useOutwayStore()
  const directoryServiceStore = useDirectoryServiceStore()
  const accessProofStore = useAccessProofStore()

  useEffect(() => {
    outwayStore.fetchAll()
    directoryServiceStore.fetchAll()
  }, []) // eslint-disable-line react-hooks/exhaustive-deps

  const validationSchema = Yup.object().shape({
    description: Yup.string()
      .max(100, t('Maximum of n characters allowed', { n: 100 }))
      .required(t('This field is required')),
    reference: Yup.string().required(t('This field is required')),
    delegatee: Yup.string()
      .max(20, t('Maximum of n characters allowed', { n: 20 }))
      .required(t('This field is required')),
    publicKeyPEM: Yup.string().required(t('This field is required')),
    validFrom: isoDateSchema(t('Invalid date')).required(
      t('This field is required'),
    ),
    validUntil: isoDateSchema().required(t('This field is required')),
    accessProofIds: Yup.array().min(1, t('This field is required')),
  })

  const selectableServices = directoryServiceStore.servicesWithAccess.reduce(
    (previousValue, service) => {
      service.accessStatesWithAccess.forEach((accessState) => {
        const { accessRequest, accessProof } = accessState

        const outwayNames = outwayStore
          .getByPublicKeyFingerprint(accessRequest.publicKeyFingerprint)
          .map((outway) => outway.name)
          .join(', ')

        previousValue.push({
          value: accessProof.id,
          label: `${service.serviceName} - ${service.organization.name} (${service.organization.serialNumber}) - via ${outwayNames} (${accessRequest.publicKeyFingerprint})`,
        })
      })

      return previousValue
    },
    [],
  )

  const handleSubmit = (values) => {
    onSubmitHandler(values)
  }

  const combinedInitialValues = {
    ...DEFAULT_INITIAL_VALUES,
    ...initialValues,
  }

  if (combinedInitialValues.validFrom) {
    combinedInitialValues.validFrom = formatDateToYYYYMMDD(
      combinedInitialValues.validFrom,
    )
  }

  if (combinedInitialValues.validUntil) {
    combinedInitialValues.validUntil = formatDateToYYYYMMDD(
      combinedInitialValues.validUntil,
    )
  }

  const isServiceOptionDisabledCheck = (option, selectedOptions) => {
    const accessProof = accessProofStore.getById(option.value)

    const selectedAccessProofModels = selectedOptions.map((option) =>
      accessProofStore.getById(option.value),
    )

    const selectedAccessProofsForSameService = selectedAccessProofModels.find(
      (selectedAccessProof) => {
        return (
          selectedAccessProof.organization.serialNumber ===
            accessProof.organization.serialNumber &&
          selectedAccessProof.serviceName === accessProof.serviceName
        )
      },
    )

    return !!selectedAccessProofsForSameService
  }

  return (
    <Formik
      initialValues={combinedInitialValues}
      validationSchema={validationSchema}
      onSubmit={(values) => {
        handleSubmit({
          ...values,
          validFrom: new Date(values.validFrom),
          validUntil: new Date(values.validUntil),
        })
      }}
    >
      <StyledForm {...props}>
        <Fieldset>
          <TextInput name="description" size="l">
            <FieldLabel
              label={t('Order description')}
              small={t('For your own reference')}
            />
          </TextInput>

          <TextInput name="reference" size="l" disabled={isEditMode}>
            <FieldLabel
              label={t('Reference')}
              small={t('This identifier is sent with each request')}
            />
          </TextInput>

          <TextInput name="delegatee" size="l" disabled={isEditMode}>
            <FieldLabel
              label={t('Delegated organization')}
              small={t('OIN of the delegatee')}
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
            name="accessProofIds"
            options={selectableServices}
            size="xl"
            isMulti
            placeholder={t('Select a service…')}
            isOptionDisabled={isServiceOptionDisabledCheck}
          >
            <FieldLabel
              label={t('Services')}
              small={t(
                'Is the service not listed? Then please first request access to this service in the directory',
              )}
            />
          </Select>
        </Fieldset>

        <Button type="submit">{submitButtonText}</Button>
      </StyledForm>
    </Formik>
  )
}

OrderForm.propTypes = {
  initialValues: shape({
    description: string,
    reference: string,
    delegatee: string,
    publicKeyPEM: string,
    validFrom: instanceOf(Date),
    validUntil: instanceOf(Date),
    accessProofIds: array,
  }),
  submitButtonText: string.isRequired,
  onSubmitHandler: func.isRequired,
  isEditMode: bool,
}

OrderForm.defaultProps = {
  isEditMode: false,
}

export default observer(OrderForm)
