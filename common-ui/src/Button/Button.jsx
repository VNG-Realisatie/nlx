import React from 'react'
import PropTypes from 'prop-types'
import styled from 'styled-components'

import BaseButton from 'src/BaseButton/BaseButton'
import { buttonStyle, iconStyle, iconRightStyle } from './buttonStyle.js'

const StyledButton = styled(BaseButton)`
    ${buttonStyle};
`

const StyledIcon = styled.div`
    ${iconStyle};
`

const StyledIconRight = styled.div`
    ${iconRightStyle};
`

export default class Button extends React.Component {
    render() {
        const {
            children,
            icon,
            iconRight,
        } = this.props

        return (
            <StyledButton
                {...this.props}
            >
                {icon &&
                    <StyledIcon {...this.props}>
                    {icon}
                    </StyledIcon>
                }
                <span>
                    {children}
                </span>
                {iconRight &&
                    <StyledIconRight {...this.props}>
                        {iconRight}
                    </StyledIconRight>
                }
            </StyledButton>
        )
    }
}

Button.propTypes = {
    size: PropTypes.oneOf(['small', 'normal', 'large']),
    variant: PropTypes.oneOf(['primary', 'secondary', 'tertiary']),
    icon: PropTypes.element,
    iconRight: PropTypes.element,
    disabled: PropTypes.bool,
}

Button.defaultProps = {
    size: 'normal',
    variant: 'primary',
}