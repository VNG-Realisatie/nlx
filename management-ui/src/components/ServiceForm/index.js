// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { arrayOf, bool, func, shape, string } from 'prop-types'
import { FieldArray, Formik } from 'formik'
import * as Yup from 'yup'
import { useTranslation } from 'react-i18next'
import {
  Button,
  Checkbox,
  Fieldset,
  Label,
  Legend,
  TextInput,
} from '@commonground/design-system'

import FormikFocusError from '../FormikFocusError'
import InwayRepository from '../../domain/inway-repository'
import usePromise from '../../hooks/use-promise'
import { showServiceVisibilityAlert } from '../ServiceVisibilityAlert'
import {
  CheckboxGroup,
  Form,
  InwaysEmptyMessage,
  InwaysLoadingMessage,
  ServiceNameWrapper,
  StyledServiceVisibilityAlert,
} from './index.styles'

const DEFAULT_INITIAL_VALUES = {
  name: '',
  endpointURL: '',
  documentationURL: '',
  apiSpecificationURL: '',
  publishedInDirectory: true,
  techSupportContact: '',
  publicSupportContact: '',
  inways: [],
}

const ServiceForm = ({
  initialValues,
  onSubmitHandler,
  submitButtonText,
  disableName,
  getInways,
  ...props
}) => {
  const { t } = useTranslation()

  const { isReady: allInwaysIsReady, result: allInways } = usePromise(getInways)
  const validationSchema = Yup.object().shape({
    name: Yup.string().required(t('This field is required')),
    endpointURL: Yup.string().required(t('Invalid endpoint URL')),
    documentationURL: Yup.string(),
    apiSpecificationURL: Yup.string(),
    publishedInDirectory: Yup.boolean(),
    techSupportContact: Yup.string(),
    publicSupportContact: Yup.string(),
    inways: Yup.array().of(Yup.string()),
  })

  return (
    <Formik
      initialValues={{
        ...DEFAULT_INITIAL_VALUES,
        ...initialValues,
        publishedInDirectory: !initialValues.internal,
      }}
      validationSchema={validationSchema}
      onSubmit={({ publishedInDirectory, ...values }) => {
        return onSubmitHandler({
          ...values,
          internal: !publishedInDirectory,
        })
      }}
    >
      {({ handleSubmit, values: { inways, publishedInDirectory } }) => (
        <Form onSubmit={handleSubmit} data-testid="form" {...props}>
          <ServiceNameWrapper>
            <TextInput
              name="name"
              data-testid="name"
              size="l"
              {...(disableName ? { disabled: true } : {})}
            >
              {t('Service name')}
            </TextInput>
          </ServiceNameWrapper>

          <Fieldset>
            <Legend>{t('API details')}</Legend>

            <TextInput name="endpointURL" data-testid="endpointURL" size="xl">
              {t('API endpoint URL')}
            </TextInput>

            <TextInput
              name="documentationURL"
              data-testid="documentationURL"
              size="xl"
            >
              {t('API documentation URL')}
            </TextInput>

            <TextInput
              name="apiSpecificationURL"
              data-testid="apiSpecificationURL"
              size="xl"
            >
              {t('API specification URL')}
            </TextInput>
          </Fieldset>

          <Fieldset>
            <Legend>{t('Contact')}</Legend>

            <TextInput
              name="techSupportContact"
              data-testid="techSupportContact"
              size="l"
            >
              {t('Tech support email')}
            </TextInput>

            <TextInput
              name="publicSupportContact"
              data-testid="publicSupportContact"
              size="l"
            >
              {t('Public support email')}
            </TextInput>
          </Fieldset>

          <Fieldset>
            <Legend>{t('Inways')}</Legend>

            {!allInwaysIsReady ? (
              <InwaysLoadingMessage />
            ) : !allInways || allInways.length === 0 ? (
              <InwaysEmptyMessage data-testid="inways-empty">
                {t('There are no inways registered to be connected')}
              </InwaysEmptyMessage>
            ) : (
              <CheckboxGroup>
                <Label>
                  {t('Connected inways for this service (optional)')}
                </Label>
                <FieldArray
                  name="inways"
                  render={({ push, remove }) => {
                    const handleChangeFor = (name) => (e) => {
                      if (e.target.checked) {
                        push(name)
                      } else {
                        const idx = inways.indexOf(name)
                        remove(idx)
                      }
                    }

                    return allInways.map(({ name }) => (
                      <Checkbox
                        key={name}
                        name="inways"
                        value={name}
                        checked={inways.includes(name)}
                        onChange={handleChangeFor(name)}
                      >
                        {name}
                      </Checkbox>
                    ))
                  }}
                />
              </CheckboxGroup>
            )}
          </Fieldset>

          <Fieldset>
            <Legend>{t('Visibility')}</Legend>

            <Checkbox name="publishedInDirectory">
              {t('Publish to central directory')}
            </Checkbox>

            {showServiceVisibilityAlert({
              internal: !publishedInDirectory,
              inways,
            }) ? (
              <StyledServiceVisibilityAlert data-testid="publishedInDirectory-warning" />
            ) : null}
          </Fieldset>

          <Button type="submit">{submitButtonText}</Button>

          <FormikFocusError />
        </Form>
      )}
    </Formik>
  )
}

ServiceForm.propTypes = {
  onSubmitHandler: func,
  initialValues: shape({
    name: string,
    endpointURL: string,
    documentationURL: string,
    apiSpecificationURL: string,
    internal: bool,
    techSupportContact: string,
    publicSupportContact: string,
    inways: arrayOf(string),
  }),
  submitButtonText: string.isRequired,
  disableName: bool,
  getInways: func,
}

ServiceForm.defaultProps = {
  initialValues: DEFAULT_INITIAL_VALUES,
  disableName: false,
  getInways: InwayRepository.getAll,
}

export default ServiceForm
