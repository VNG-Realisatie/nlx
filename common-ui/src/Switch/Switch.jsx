import React, { Fragment, PureComponent } from 'react'
import PropTypes from 'prop-types'
import styled from 'styled-components'
import { ErrorMessage } from 'formik'

import { inputStyle, labelStyle } from './switchStyle.js'

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
        const { field, label, id, required, disabled } = this.props
        const { name } = field

        return (
            <Fragment>
                <StyledWrapper>
                    <StyledInput
                        type="checkbox"
                        id={id}
                        disabled={disabled}
                        {...field}
                    />
                    <StyledLabel htmlFor={id} title={label}>
                        {label}
                        {required && ' *'}
                    </StyledLabel>
                </StyledWrapper>
                <ErrorMessage name={name} />
            </Fragment>
        )
    }
}

Switch.propTypes = {
    field: PropTypes.object.isRequired, // provided by Formik
    id: PropTypes.string,
    label: PropTypes.string,
    required: PropTypes.bool,
    disabled: PropTypes.bool,
}

Switch.defaultProps = {
    id: '',
    label: '',
    required: false,
    disabled: false,
}

export default Switch
