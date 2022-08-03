// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React, { useEffect } from 'react'
import { func, shape, string } from 'prop-types'
import { Formik } from 'formik'
import * as Yup from 'yup'
import { useTranslation } from 'react-i18next'
import {
  Button,
  Fieldset,
  Legend,
  Select,
  FieldLabel,
  TextInput,
} from '@commonground/design-system'
import { observer } from 'mobx-react'
import FormikFocusError from '../../../../components/FormikFocusError'
import { useConfirmationModal } from '../../../../components/ConfirmationModal'
import { useInwayStore } from '../../../../hooks/use-stores'
import {
  StyledForm,
  InwaysEmptyMessage,
  InwaysLoadingMessage,
} from './index.styles'

const DEFAULT_INITIAL_VALUES = {
  organizationInway: '',
  organizationEmailAddress: '',
}

const Form = ({ initialValues, onSubmitHandler, ...props }) => {
  const { t } = useTranslation()
  const inwayStore = useInwayStore()
  const [ConfirmationModal, confirmModal] = useConfirmationModal({
    okText: t('Save'),
    children: t(
      'By removing the organization Inway it is no longer possible to process or receive access requests',
    ),
  })

  useEffect(() => {
    inwayStore.fetchInways()
  }, [inwayStore])

  const validateAndSubmit = async (values) => {
    // Check if organization inway gets removed
    if (!values.organizationInway && initialValues.organizationInway) {
      if (await confirmModal()) {
        onSubmitHandler(values)
      }

      return
    }

    onSubmitHandler(values)
    return
  }

  const validationSchema = Yup.object().shape({
    organizationInway: Yup.string(),
    organizationEmailAddress: Yup.string().email(
      t('Please enter a valid email address'),
    ),
  })

  const selectInwayOptions = inwayStore.inways.map((inway) => ({
    value: inway.name,
    label: inway.name,
  }))

  const emptyOption = { value: '', label: t('None') }
  selectInwayOptions.unshift(emptyOption)

  return (
    <>
      <Formik
        initialValues={{
          ...DEFAULT_INITIAL_VALUES,
          ...initialValues,
        }}
        validationSchema={validationSchema}
        onSubmit={validateAndSubmit}
      >
        {({ handleSubmit }) => (
          <StyledForm onSubmit={handleSubmit} data-testid="form" {...props}>
            <Fieldset>
              <Legend>{t('General settings')}</Legend>
              {inwayStore.isFetching ? (
                <InwaysLoadingMessage />
              ) : inwayStore.inways.length === 0 ? (
                <InwaysEmptyMessage>
                  {t('There are no inways available')}
                </InwaysEmptyMessage>
              ) : (
                <Select
                  id="organizationInway"
                  name="organizationInway"
                  options={selectInwayOptions}
                >
                  <FieldLabel
                    label={t('Organization inway')}
                    small={t(
                      'This inway is used to be able to retrieve & confirm access requests from other organizations and synchronize orders with other organizations.',
                    )}
                  />
                </Select>
              )}
              <TextInput name="organizationEmailAddress" size="l">
                <FieldLabel
                  label={t('Organization email address')}
                  small={t(
                    'This email address will receive important updates about NLX.',
                  )}
                />
              </TextInput>
            </Fieldset>

            <Fieldset>
              <Button type="submit">{t('Save settings')}</Button>
            </Fieldset>

            <FormikFocusError />
          </StyledForm>
        )}
      </Formik>

      <ConfirmationModal />
    </>
  )
}

Form.propTypes = {
  onSubmitHandler: func,
  initialValues: shape({
    organizationInway: string,
    organizationEmailAddress: string,
  }),
}

Form.defaultProps = {
  // eslint-disable-next-line @typescript-eslint/no-empty-function
  onSubmitHandler: () => {},
  initialValues: DEFAULT_INITIAL_VALUES,
}

export default observer(Form)
