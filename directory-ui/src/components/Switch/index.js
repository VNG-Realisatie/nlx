// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import React from 'react'
import { string, bool } from 'prop-types'
import { Wrapper, Input, Label } from './index.styles'

const Switch = ({ label, id, required, disabled, ...props }) => (
  <Wrapper {...props}>
    <Input type="checkbox" id={id} disabled={disabled} defaultChecked />
    <Label htmlFor={id} title={label}>
      {label}
      {required && ' *'}
    </Label>
  </Wrapper>
)

Switch.propTypes = {
  id: string,
  label: string,
  required: bool,
  disabled: bool,
}

Switch.defaultProps = {
  id: '',
  label: '',
  required: false,
  disabled: false,
}

export default Switch
