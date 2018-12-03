import React, { Fragment, PureComponent } from 'react'
import PropTypes from 'prop-types'
import styled from 'styled-components'
import { ErrorMessage } from 'formik'

import { inputStyle, boxStyle, labelStyle } from './checkboxStyle.js'

const StyledCheckboxWrapper = styled.div`
    position: relative;
`

const StyledCheckboxInput = styled.input`
    ${inputStyle};
`

const StyledCheckboxBox = styled.div`
    ${boxStyle};
`

const StyledCheckboxLabel = styled.label`
    ${labelStyle};
`

class Checkbox extends PureComponent {
    render() {
        const { field, label, id, required, disabled } = this.props
        const { name } = field

        return (
            <Fragment>
                <StyledCheckboxWrapper>
                    <StyledCheckboxInput
                        type="checkbox"
                        id={id}
                        disabled={disabled}
                        {...field}
                    />
                    <StyledCheckboxLabel htmlFor={id} title={label}>
                        {label}
                        {required && ' *'}
                    </StyledCheckboxLabel>
                    <StyledCheckboxBox>
                        <svg width={10} height={8} viewBox="0 0 10 8">
                            <g
                                id="Feed"
                                stroke="none"
                                strokeWidth={1}
                                fill="none"
                                fillRule="evenodd"
                                transform="translate(-318 -2878)"
                            >
                                <path
                                    d="M327.921 2879.512l-4.963 6.339c-.14.179-.41.2-.578.046l-4.26-3.928a.368.368 0 0 1-.012-.533l1.072-1.089a.398.398 0 0 1 .55-.013l2.684 2.473 3.647-4.658a.398.398 0 0 1 .544-.073l1.24.91a.367.367 0 0 1 .076.526"
                                    id="Fill-1-Copy"
                                    fill="currentColor"
                                    fillRule="nonzero"
                                />
                            </g>
                        </svg>
                    </StyledCheckboxBox>
                </StyledCheckboxWrapper>
                <ErrorMessage name={name} />
            </Fragment>
        )
    }
}

Checkbox.propTypes = {
    field: PropTypes.object.isRequired, // provided by Formik
    id: PropTypes.string,
    label: PropTypes.string,
    required: PropTypes.bool,
    disabled: PropTypes.bool,
}

Checkbox.defaultProps = {
    id: '',
    label: '',
    required: false,
    disabled: false,
}

export default Checkbox
