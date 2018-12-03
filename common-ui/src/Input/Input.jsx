import React, { Fragment, PureComponent } from 'react'
import PropTypes from 'prop-types'
import styled from 'styled-components'
import { ErrorMessage } from 'formik'

import { inputStyle, labelStyle } from './inputStyle.js'
import Error from '../Error/Error'

const StyledWrapper = styled.div`
    position: relative;
    padding-top: 5px;
`

const StyledInput = styled.input`
    ${inputStyle};
`

const StyledLabel = styled.div`
    ${labelStyle};
`

class Input extends PureComponent {
    render() {
        const { field, id, label, required, disabled } = this.props
        const { name, value } = field
        const hasValue = !!value

        return (
            <Fragment>
                <StyledWrapper>
                    <StyledInput
                        type="text"
                        id={id || name}
                        placeholder={label || name}
                        disabled={disabled}
                        {...field}
                    />
                    <StyledLabel
                        htmlFor={id || name}
                        title={label || name}
                        small={hasValue}
                    >
                        <span>
                            {label || name}
                            {required && ' *'}
                        </span>
                    </StyledLabel>
                </StyledWrapper>
                <Error>
                    <ErrorMessage name={name} />
                </Error>
            </Fragment>
        )
    }
}

Input.propTypes = {
    field: PropTypes.object.isRequired, // provided by Formik
    id: PropTypes.string,
    label: PropTypes.string,
    required: PropTypes.bool,
    disabled: PropTypes.bool,
}

Input.defaultProps = {
    id: '',
    label: '',
    required: false,
    disabled: false,
}

export default Input
