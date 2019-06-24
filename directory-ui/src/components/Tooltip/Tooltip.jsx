import React from 'react'
import Tippy from '../Tippy/Tippy'

const Tooltip = ({ children, ...rest }) => <Tippy {...rest}>{children}</Tippy>

Tooltip.defaultProps = {
    placement: 'bottom',
    animateFill: true,
    duration: [225, 200],
    hideOnClick: false,
    distance: 8,
}

export default Tooltip
