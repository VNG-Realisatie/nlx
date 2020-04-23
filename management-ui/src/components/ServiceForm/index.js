// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { func, shape, string, oneOf, bool } from 'prop-types'
import { Formik } from 'formik'
import {
  Button,
  TextInput,
  Radio,
  Checkbox,
  Fieldset,
  Legend,
} from '@commonground/design-system'

import { ThemeProvider } from 'styled-components'
import * as Yup from 'yup'
import { useTranslation } from 'react-i18next'
import theme from '../../theme'
import {
  AUTHORIZATION_TYPE_NONE,
  AUTHORIZATION_TYPE_WHITELIST,
} from '../../vocabulary'
import { Form } from './index.styles'

const DEFAULT_INITIAL_VALUES = {
  name: '',
  endpointURL: '',
  documentationURL: '',
  apiSpecificationURL: '',
  publishedInDirectory: true,
  techSupportContact: '',
  publicSupportContact: '',
  authorizationSettings: {
    mode: AUTHORIZATION_TYPE_WHITELIST,
  },
}

const ServiceForm = ({
  initialValues,
  onSubmitHandler,
  submitButtonText,
  disableName,
  ...props
}) => {
  const { t } = useTranslation()

  const validationSchema = Yup.object().shape({
    name: Yup.string().required(t('This field is required.')),
    endpointURL: Yup.string().required(t('Invalid endpoint URL.')),
    documentationURL: Yup.string(),
    apiSpecificationURL: Yup.string(),
    publishedInDirectory: Yup.boolean(),
    techSupportContact: Yup.string(),
    publicSupportContact: Yup.string(),
    authorizationSettings: Yup.object().shape({
      mode: Yup.mixed().oneOf([
        AUTHORIZATION_TYPE_WHITELIST,
        AUTHORIZATION_TYPE_NONE,
      ]),
    }),
  })

  return (
    <ThemeProvider theme={theme}>
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
        {({ handleSubmit }) => (
          <Form onSubmit={handleSubmit} data-testid="form" {...props}>
            <Fieldset>
              <Legend>{t('API details')}</Legend>

              <TextInput
                name="name"
                id="name"
                size="l"
                {...(disableName ? { disabled: true } : {})}
              >
                {t('API name')}
              </TextInput>

              <TextInput name="endpointURL" id="endpointURL" size="xl">
                {t('API endpoint URL')}
              </TextInput>

              <TextInput
                name="documentationURL"
                id="documentationURL"
                size="xl"
              >
                {t('API documentation URL')}
              </TextInput>

              <TextInput
                name="apiSpecificationURL"
                id="apiSpecificationURL"
                size="xl"
              >
                {t('API specification URL')}
              </TextInput>

              <Checkbox name="publishedInDirectory" id="publishedInDirectory">
                {t('Publish to central directory')}
              </Checkbox>
            </Fieldset>

            <Fieldset>
              <Legend>{t('Contact')}</Legend>

              <TextInput
                name="techSupportContact"
                id="techSupportContact"
                size="l"
              >
                {t('Tech support email')}
              </TextInput>

              <TextInput
                name="publicSupportContact"
                id="publicSupportContact"
                size="l"
              >
                {t('Public support email')}
              </TextInput>
            </Fieldset>

            <Fieldset>
              <Legend>{t('Authorization')}</Legend>

              <Radio.Group label={t('Type of authorization')}>
                <Radio
                  id="authorizationModeWhitelist"
                  name="authorizationSettings.mode"
                  value={AUTHORIZATION_TYPE_WHITELIST}
                >
                  {t('Whitelist for authorized organizations')}
                </Radio>

                <Radio
                  id="authorizationModeNone"
                  name="authorizationSettings.mode"
                  value={AUTHORIZATION_TYPE_NONE}
                >
                  {t('Allow all organizations')}
                </Radio>
              </Radio.Group>
            </Fieldset>

            <Button type="submit">{submitButtonText}</Button>
          </Form>
        )}
      </Formik>
    </ThemeProvider>
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
    authorizationSettings: shape({
      mode: oneOf([AUTHORIZATION_TYPE_WHITELIST, AUTHORIZATION_TYPE_NONE]),
    }),
  }),
  submitButtonText: string.isRequired,
  disableName: bool,
}

ServiceForm.defaultProps = {
  initialValues: DEFAULT_INITIAL_VALUES,
  disableName: false,
}

export default ServiceForm
