// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { func, shape, string, oneOf, bool } from 'prop-types'
import { Formik, Field, useField } from 'formik'
import { Button, Input } from '@commonground/design-system'

import { ThemeProvider } from 'styled-components'
import * as Yup from 'yup'
import { useTranslation } from 'react-i18next'
import theme from '../../../theme'
import {
  AUTHORIZATION_TYPE_NONE,
  AUTHORIZATION_TYPE_WHITELIST,
} from '../../../vocabulary'
import FieldValidationMessage from './FieldValidationMessage'
import {
  Form,
  StyledField,
  Fieldset,
  Label,
  Legend,
  StyledLabelWithInput,
} from './index.styles'

const DEFAULT_INITIAL_VALUES = {
  name: '',
  endpointURL: '',
  documentationURL: '',
  apiSpecificationURL: '',
  internal: true,
  techSupportContact: '',
  publicSupportContact: '',
  authorizationSettings: {
    mode: AUTHORIZATION_TYPE_WHITELIST,
  },
}

const FieldWithValidation = ({ label, ...props }) => {
  const [field, meta] = useField(props)
  const { error, touched } = meta

  return (
    <>
      <Label htmlFor={field.name}>{label}</Label>
      <StyledField
        {...field}
        {...props}
        className={error && touched ? 'invalid' : null}
      />
      {error && touched ? (
        <FieldValidationMessage data-testid={`${props.id}-error`}>
          {error}
        </FieldValidationMessage>
      ) : null}
    </>
  )
}

FieldWithValidation.propTypes = {
  id: string,
  label: string,
}

const AddServiceForm = ({
  initialValues,
  onSubmitHandler,
  submitButtonText,
  ...props
}) => {
  const { t } = useTranslation()

  const validationSchema = Yup.object().shape({
    name: Yup.string().required(t('This field is required.')),
    endpointURL: Yup.string().required(t('Invalid endpoint URL.')),
    documentationURL: Yup.string(),
    apiSpecificationURL: Yup.string(),
    internal: Yup.boolean(),
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
        initialValues={{ ...DEFAULT_INITIAL_VALUES, ...initialValues }}
        validationSchema={validationSchema}
        onSubmit={(values) => onSubmitHandler(values)}
      >
        {({ handleSubmit }) => (
          <Form onSubmit={handleSubmit} data-testid="form" {...props}>
            <Fieldset>
              <Legend>{t('API details')}</Legend>

              <FieldWithValidation
                label={t('API name')}
                name="name"
                id="name"
                size="s"
              />

              <FieldWithValidation
                label={t('API endpoint URL')}
                name="endpointURL"
                id="endpointURL"
                size="m"
              />

              <FieldWithValidation
                label={t('API documentation URL')}
                name="documentationURL"
                id="documentationURL"
                size="m"
              />

              <FieldWithValidation
                label={t('API specification URL')}
                name="apiSpecificationURL"
                id="apiSpecificationURL"
                size="m"
              />

              <StyledLabelWithInput>
                <Field
                  as={Input}
                  id="internal"
                  name="internal"
                  type="checkbox"
                />
                {t('Publish to central directory')}
              </StyledLabelWithInput>
            </Fieldset>

            <Fieldset>
              <Legend>{t('Contact')}</Legend>

              <FieldWithValidation
                label={t('Tech support email')}
                name="techSupportContact"
                id="techSupportContact"
                size="s"
              />

              <FieldWithValidation
                label={t('Public support email')}
                name="publicSupportContact"
                id="publicSupportContact"
                size="s"
              />
            </Fieldset>

            <Fieldset>
              <Legend>{t('Authorization')}</Legend>

              <Label>{t('Type of authorization')}</Label>
              <StyledLabelWithInput>
                <Field
                  id="authorizationModeWhitelist"
                  name="authorizationSettings.mode"
                  type="radio"
                  as={Input}
                  value={AUTHORIZATION_TYPE_WHITELIST}
                />
                {t('Whitelist for authorized organizations')}
              </StyledLabelWithInput>

              <StyledLabelWithInput>
                <Field
                  id="authorizationModeNone"
                  name="authorizationSettings.mode"
                  type="radio"
                  as={Input}
                  value={AUTHORIZATION_TYPE_NONE}
                />
                {t('Allow all organizations')}
              </StyledLabelWithInput>
            </Fieldset>

            <Button type="submit">{submitButtonText}</Button>
          </Form>
        )}
      </Formik>
    </ThemeProvider>
  )
}

AddServiceForm.propTypes = {
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
}

AddServiceForm.defaultProps = {
  initialValues: DEFAULT_INITIAL_VALUES,
}

export default AddServiceForm
