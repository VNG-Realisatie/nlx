import React, { PureComponent } from 'react'
import { string, bool } from 'prop-types'
import styled from 'styled-components'

import { inputStyle, labelStyle } from './Switch.styles'

const StyledWrapper = styled.div`
    position: relative;
`

const StyledInput = styled.input`
    ${inputStyle};
`

const StyledLabel = styled.label`
    ${labelStyle};
`

class Switch extends PureComponent {
  render() {
    const { label, id, required, disabled, ...props } = this.props

    return (
      <StyledWrapper {...props}>
        <StyledInput
          type="checkbox"
          id={id}
          disabled={disabled}
          defaultChecked
        />
        <StyledLabel htmlFor={id} title={label}>
          {label}
          {required && ' *'}
        </StyledLabel>
      </StyledWrapper>
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
