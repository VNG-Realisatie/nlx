// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import React, { PureComponent } from 'react'
import { string, bool } from 'prop-types'
import { Wrapper, Input, Label } from './Switch.styles'

class Switch extends PureComponent {
  render() {
    const { label, id, required, disabled, ...props } = this.props

    return (
      <Wrapper {...props}>
        <Input type="checkbox" id={id} disabled={disabled} defaultChecked />
        <Label htmlFor={id} title={label}>
          {label}
          {required && ' *'}
        </Label>
      </Wrapper>
    )
  }
}

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
