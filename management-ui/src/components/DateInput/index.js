// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import propTypes from 'prop-types'
import { TextInput, FieldLabel } from '@commonground/design-system'
import { useTranslation } from 'react-i18next'

const isTypeDateSupported = (function () {
  const input = document.createElement('input')
  const value = 'a'
  input.setAttribute('type', 'date')
  input.setAttribute('value', value)
  return input.value !== value
})()

// Placeholder is ignored when browser supports type="date"
const DateInput = ({ name, label, ...props }) => {
  const { t } = useTranslation()
  const smallLabel = !isTypeDateSupported
    ? `${t('Date notation:')} ${t('yyyy-mm-dd')}`
    : ''

  return (
    <TextInput type="date" name={name} placeholder={t('yyyy-mm-dd')} {...props}>
      <FieldLabel label={label} small={smallLabel} />
    </TextInput>
  )
}

DateInput.propTypes = {
  name: propTypes.string.isRequired,
  label: propTypes.string,
}

export default DateInput
export { isoDateSchema, dateToIsoFormat } from './utils'
