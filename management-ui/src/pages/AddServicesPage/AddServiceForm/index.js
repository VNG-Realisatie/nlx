import React from 'react'
import { func, shape, string, oneOf, bool } from 'prop-types'
import { Formik, useField } from 'formik'
import { Button } from '@commonground/design-system'

import { ThemeProvider } from 'styled-components'
import * as Yup from 'yup'
import { useTranslation } from 'react-i18next'
import i18next from 'i18next'
import theme from '../../../theme'
import {
  AUTHORIZATION_TYPE_NONE,
  AUTHORIZATION_TYPE_WHITELIST,
} from '../../../vocabulary'
import FieldValidationMessage from './FieldValidationMessage'
import {
  StyledField,
  Fieldset,
  Label,
  FieldInfo,
  StyledLayerTypeOption,
  StyledLayerTypeOptionGroup,
  StyledExpectedDateGroup,
  StyledInput,
} from './index.styles'

const DEFAULT_INITIAL_VALUES = {
  name: '',
  endpointURL: '',
  documentationURL: '',
  apiSpecificationURL: '',
  internal: true,
  techSupportContact: '',
  publicSupportContact: '',
  authorizationMode: AUTHORIZATION_TYPE_WHITELIST,
}

const validationSchema = Yup.object().shape({
  name: Yup.string().required('Dit veld is verplicht.'), // TODO: use translations
  endpointURL: Yup.string().required('Ongeldig endpoint URL.'), // TODO: use translations

  documentationURL: Yup.string(),
  apiSpecificationURL: Yup.string(),
  internal: Yup.boolean(),
  techSupportContact: Yup.string(),
  publicSupportContact: Yup.string(),
  authorizationMode: Yup.mixed().oneOf([
    AUTHORIZATION_TYPE_WHITELIST,
    AUTHORIZATION_TYPE_NONE,
  ]),
})

const FieldWithValidation = ({ id, label, ...props }) => {
  const [field, meta] = useField(props)
  const { error, touched } = meta

  return (
    <>
      <Label htmlFor={field.name}>{label}</Label>
      <StyledInput
        id={id || undefined}
        {...field}
        className={error && touched ? 'invalid' : null}
      />
      {error && touched ? (
        <FieldValidationMessage data-testid={`${id}-error`}>
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
  return (
    <ThemeProvider theme={theme}>
      <Formik
        initialValues={{ ...DEFAULT_INITIAL_VALUES, ...initialValues }}
        validationSchema={validationSchema}
        onSubmit={(values) => onSubmitHandler(values)}
      >
        {({ handleSubmit, errors, touched, values, submitCount }) => (
          <form onSubmit={handleSubmit} data-testid="form" {...props}>
            <fieldset>
              <legend>{t('API details')}</legend>

              <FieldWithValidation
                label={t('API naam')}
                name="name"
                id="name"
              />

              <FieldWithValidation
                label={t('API endpoint URL')}
                name="endpointURL"
                id="endpointURL"
              />

              <FieldWithValidation
                label={t('API documentatie URL')}
                name="documentationURL"
                id="documentationURL"
              />

              <FieldWithValidation
                label={t('API specificatie URL')}
                name="apiSpecificationURL"
                id="apiSpecificationURL"
              />

              <FieldWithValidation
                label={t('Publiceren in de centrale directory')}
                name="internal"
                id="internal"
                type="checkbox"
              />
            </fieldset>

            <fieldset>
              <legend>{t('Contact')}</legend>

              <FieldWithValidation
                label={t('Tech support email')}
                name="techSupportContact"
                id="techSupportContact"
              />

              <FieldWithValidation
                label={t('Public support email')}
                name="publicSupportContact"
                id="publicSupportContact"
              />
            </fieldset>

            <fieldset>
              <legend>{t('Authorizatie')}</legend>

              <Label>{t('Type authorisatie')}</Label>
              <StyledInput
                id="authorizationModeWhitelist"
                name="authorizationMode"
                type="radio"
                value={AUTHORIZATION_TYPE_WHITELIST}
              />
              <Label htmlFor="authorizationModeWhitelist">
                {t('Whitelist voor geauthorizeerde organisaties')}
              </Label>

              <StyledInput
                id="authorizationModeNone"
                name="authorizationMode"
                type="radio"
                value={AUTHORIZATION_TYPE_NONE}
              />
              <Label htmlFor="authorizationModeNone">
                {t('Alle organisaties toestaan')}
              </Label>
            </fieldset>

            <Button type="submit">{submitButtonText}</Button>
          </form>
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
    authorizationMode: oneOf([
      AUTHORIZATION_TYPE_WHITELIST,
      AUTHORIZATION_TYPE_NONE,
    ]),
  }),
  submitButtonText: string,
}

AddServiceForm.defaultProps = {
  initialValues: DEFAULT_INITIAL_VALUES,
  submitButtonText: i18next.t('Service toevoegen'),
}

export default AddServiceForm
